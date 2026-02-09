package jwt

import (
	"monitorservice/pkg/cache"
	"monitorservice/pkg/config"
	"monitorservice/pkg/encrypt"
	"monitorservice/pkg/encrypt/base64url"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/*
每次產生的 token 都是由 config 內設定的 EncryptString 字串加上隨機的16碼字串，使用 MD5Hash (不可逆)
產生 JWT 加密所需要的 EncryptKey
TIP:
	* 程序每次重新啟動都會經由使用 MD5Hash 產生並更換 EncryptKey
	* 重新啟動之前所有發出的 token 都會失效
	* 在加密資料內加入目前時間資料，避免同一個帳號發出的 token 都相同
	* (?)將每次產生的 token 紀錄至 redis 內，避免重開後台服務，客戶端 token 失效問題

標準中註冊的聲明:(建議但不強制使用)
	* iss(Issuer)：頒發者，是區分大小寫的字串，可以是一個字串或是網址
	* sub(Subject)：主體內容，是區分大小寫的字串，可以是一個字串或是網址
	* aud(Audience)：觀眾，是區分大小寫的字串，可以是一個字串或是網址
	* exp(Expiration Time)：Expiration Time，過期時間，是數字日期
	* nbf(Not Before)：定義在什麼時間之前，不可用，是數字日期
	* iat(Issued At)：頒發時間，是數字日期
	* jti(JWT ID)：唯一識別碼，是區分大小寫的字串

QUEST:
* 僅使用jwt實現單點登入會遇到兩個問題

	1. 使用者無法主動登出，即服務端發出token後，無法主動銷燬token，使用者還可以用通過token訪問系統，
	要增加快取登出使用者token到黑名單的方式，變相實現登出。
	2. token續期問題，access_token攜帶有效期，有效期過了無法自動續期。需要提供續期介面（renewal），
	服務端在生成access_token同時還會生成refresh_token（有效期比access_token長），使用者可以通過有效的
	refresh_token和access_token訪問renewal介面重新獲取新的refresh_token和access_token。
		* refresh token 不傳給使用者，產生後直接丟棄
		* refresh token 與 access token 都由 JWT 產生

*/

type JwtManager struct {
	EncryptKey          []byte // jwt 加密字串
	Issuer              string // 指定發行人
	ExpireTimeMin       int    // 過期時間
	TokenCacheList      cache.ILocalDataCache
	LockList            cache.ILocalDataCache //[token, lock]
	UsernameList        cache.ILocalDataCache //[username, token]
	BlackTokenCacheList cache.ILocalDataCache
	// BlackTokenList
}

func NewJwtManager(config config.Config) *JwtManager {

	md5HashEncryptString := encrypt.GetMD5Hash(config.GetJwt().EncryptString + encrypt.CreateIV())

	return &JwtManager{
		EncryptKey:          []byte(md5HashEncryptString),
		Issuer:              config.GetJwt().Issuer,
		ExpireTimeMin:       config.GetJwt().ExpireTimeMin,
		TokenCacheList:      cache.NewLocalDataCache(),
		LockList:            cache.NewLocalDataCache(),
		UsernameList:        cache.NewLocalDataCache(),
		BlackTokenCacheList: cache.NewLocalDataCache(),
	}
}

// 根據用戶的用戶名和密碼參數產生憑證物件
func (p *JwtManager) CreateClaims(baseClaims BaseClaims) CustomClaims {
	nowTime := time.Now().UTC()

	// 產生憑證時，將用戶輸入的必要資訊作加密

	baseClaims.Username = base64url.Encode([]byte(baseClaims.Username))
	baseClaims.Password = base64url.Encode([]byte(baseClaims.Password))
	baseClaims.Nickname = base64url.Encode([]byte(baseClaims.Nickname))

	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: 0, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(nowTime.Add(time.Second * -1)),
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(time.Minute * time.Duration(p.ExpireTimeMin))), // 過期時間
			Issuer:    p.Issuer,
			IssuedAt:  jwt.NewNumericDate(nowTime), // 發行時間
			ID:        "",
		},
		// IsBlackList: false,
	}
	return claims
}

// 根據憑證物件產生 JWT token
func (p *JwtManager) GenerateToken(claims CustomClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(p.EncryptKey)
	p.AddToken(token, &claims)
	username, _ := base64url.Decode(claims.Username)
	p.AddUsername(string(username), token)
	p.AddLock(token, new(sync.Mutex))
	return token, err
}

// 將產生之 token 存儲至本地端記憶體內
func (p *JwtManager) AddToken(token string, claims *CustomClaims) {
	p.TokenCacheList.Add(token, claims)
}

// 將產生之 token 存儲至黑名單內
func (p *JwtManager) AddBlackToken(token string, claims *CustomClaims) {
	p.BlackTokenCacheList.Add(token, claims)
}

// 藉由傳入的用戶帳號檢查將該帳號目前有效的token存儲至黑名單內
func (p *JwtManager) AddBlackTokenByUsername(username string) bool {
	if token, ok := p.UsernameList.Get(username); ok {
		if claims, ok := p.TokenCacheList.Get(token); ok {
			p.BlackTokenCacheList.Add(token, claims)
			return true
		}
	}

	return false
}

// 確認傳入的 token 是否已在黑名單中
func (p *JwtManager) GetClaims(token string) *CustomClaims {
	if claims, ok := p.TokenCacheList.Get(token); ok {
		return claims.(*CustomClaims)
	} else {
		return nil
	}
}

// 確認傳入的 token 是否已在黑名單中
func (p *JwtManager) CheckBlackTokenExist(token string) bool {
	_, isHad := p.BlackTokenCacheList.Get(token)
	return isHad
}

// 根據傳入的token值獲取到Claims對象信息(進而獲取其中的用戶名和密碼)
func (p *JwtManager) ParseToken(token string) (*CustomClaims, error) {
	// 用于解析鑒權的聲明，方法內部主要是具體的解碼和校驗的過程，最終返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return p.EncryptKey, nil
	})
	if tokenClaims != nil {
		// 從tokenClaims中獲取到Claims對象，並使用斷言，將該對象轉換為我們自己定義的Claims
		// 要傳入指針，項目結構體都是用指針傳遞，節省空間
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid { // Valid()驗證基于時間的聲明
			return claims, nil
		}
	}
	return nil, err
}

// 將成功產生token 之 username 存儲至本地端記憶體內
func (p *JwtManager) AddUsername(username string, token string) {
	p.UsernameList.Add(username, token)
}

// 確認傳入的 username 是否已產生 token
func (p *JwtManager) CheckUsernameExist(username string) bool {

	_, isHad := p.UsernameList.Get(username)
	return isHad
}

func (p *JwtManager) AddLock(token string, mu *sync.Mutex) {
	p.LockList.Add(token, mu)
}

func (p *JwtManager) GetLock(token string) *sync.Mutex {
	if mu, ok := p.LockList.Get(token); ok {
		return mu.(*sync.Mutex)
	} else {
		return nil
	}
}
