package controller_test

import (
	"backend/api/game/model"
	"backend/pkg/encrypt/aescbc"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	table_model "backend/server/table/model"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// 參數加解密測式
func TestEncryptChannelhandleParam(t *testing.T) {
	/*
	   s=0&account=111111&money=100&orderid=1000120170306143036949111111&ip=127.0.0.1&lineCode=text11&KindID=0

	   {"s":100,"m":"/channelHandle","d":{"code":0,"url":"https://h5.ky34.com/index.ht
	   ml?account=10001_111111&token=FBE54A7273EE4F15B363C3F98F32B19F&lang=zh-CN&KindI
	   D=0"}}
	*/
	aesKey := "ddxbst648uf7hdbc" //16bit
	md5Key := "f7c637cd39679e99"
	agentId := 3
	timestamp := time.Now().UnixMilli()
	log.Printf("時間戳(Unix 時間戳帶上毫秒) is %v", timestamp)

	aesKey = "45b52080ccf56cf6"
	md5Key = "b06c136821bb29ee"

	paramData := "s=0&account=kinco&money=0&orderid=1000120170306143076949222244&ip=172.30.0.152&linecode=1&kind=1001"
	paramData = "s=0&account=brvnd_brtest0168&money=0&orderid=100220230630133432604brvnd_brtest0168&kind=0"
	// paramData = "s=1&account=kenny&money=100&orderid=1000120170306143036949111111"
	// paramData = "s=2&account=kenny&money=100&orderid=1000121171306143036949111111"
	// paramData = "s=3&account=kinco&money=100&orderid=100012017030614303694911555"
	// paramData = "s=1&account=dcc_00000001"
	// paramData = "s=5&account=kinco"

	paramDataAesEncoding, err := aescbc.AesEncrypt([]byte(paramData), []byte(aesKey))
	if err != nil {
		t.Fatalf("EncryptAES err : %v", err)
	}

	b64Encoding := base64.StdEncoding.EncodeToString([]byte(paramDataAesEncoding))
	log.Printf("參數加密字符串 is %s", b64Encoding)

	ss := strconv.Itoa(agentId) + utils.Int64ToString(timestamp) + md5Key
	s32 := md5.Hash32bit(ss)
	log.Printf("Md5 校驗字符串 is %s", s32)
}

// DB 連線測試
func TestConnectDB(t *testing.T) {
	// database/sql
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}

	gameName := make([]string, 0)
	gameName = append(gameName, "user_play_log_baccarat")
	gameName = append(gameName, "user_play_log_fantan")
	gameName = append(gameName, "user_play_log_colordisc")
	gameName = append(gameName, "user_play_log_prawncrab")
	gameName = append(gameName, "user_play_log_hundredsicbo")
	gameName = append(gameName, "user_play_log_blackjack")
	gameName = append(gameName, "user_play_log_sangong")
	gameName = append(gameName, "user_play_log_bullbull")
	gameName = append(gameName, "user_play_log_texas")
	gameName = append(gameName, "user_play_log_fruitslot")
	gameName = append(gameName, "user_play_log_rcfishing")
	gameName = append(gameName, "user_play_log_ockfight")
	gameName = append(gameName, "user_play_log_dogracing")
	gameName = append(gameName, "user_play_log_rummy")
	gameName = append(gameName, "user_play_log_goldenflower")
	gameName = append(gameName, "user_play_log_pokdeng")
	gameName = append(gameName, "user_play_log_rocket")
	gameName = append(gameName, "user_play_log_catte")
	gameName = append(gameName, "user_play_log_fruit777slot")
	gameName = append(gameName, "user_play_log_chinesepoker")
	gameName = append(gameName, "user_play_log_plinko")
	gameName = append(gameName, "user_play_log_andarbahar")
	gameName = append(gameName, "user_play_log_okey")
	gameName = append(gameName, "user_play_log_teenpatti")
	gameName = append(gameName, "user_play_log_megsharkslot")
	gameName = append(gameName, "user_play_log_midasslot")
	gameName = append(gameName, "user_play_log_roulette")
	gameName = append(gameName, "user_play_log_friendstexas")

	preQuery :=
		`SELECT date_trunc('hour'::text, tmp.bet_time) AS log_time,
			tmp.level_code,
			tmp.game_id,
			count(DISTINCT tmp.user_id) AS bet_user,
			count(tmp.lognumber) AS bet_count,
			sum(tmp.ya_score) AS sum_ya,
			sum(tmp.valid_score) AS sum_valid_ya,
			sum(tmp.de_score) AS sum_de,
			0 AS sum_bonus,
			sum(tmp.tax) AS sum_tax
		FROM ( %s ) tmp
		GROUP BY (date(tmp.bet_time)), tmp.agent_id;`

	tableQuery := `SELECT lognumber,
						level_code,
						user_id,
						game_id,
						ya_score,
						valid_score,
						de_score,
						tax,
						bonus,
						bet_time
					FROM %s
					WHERE is_robot = 0 AND bet_time >= $%d AND bet_time < $%d`
	unionQuery := " UNION "
	fromQuery := ""
	dateTimeCount := 1
	// arges := make([]any, 0)
	for k, v := range gameName {

		fromQuery = fromQuery + fmt.Sprintf(tableQuery, v, dateTimeCount, dateTimeCount+1)
		dateTimeCount += 2
		// arges = append(arges, startTime, endTime)

		if len(gameName)-1 > k {
			fromQuery = fromQuery + unionQuery
		}

	}

	query := fmt.Sprintf(preQuery, fromQuery)

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("err is: %v", err)
	}

	defer rows.Close()

	tmps := make([]table_model.RpRecordData, 0)

	for rows.Next() {

		var tmp table_model.RpRecordData
		// newline every scan 5 column
		if err := rows.Scan(&tmp.LogTime, &tmp.AgentId, &tmp.BetUser, &tmp.BetCount, &tmp.SumYa,
			&tmp.SumValidYa, &tmp.SumDe, &tmp.SumTax); err != nil {
			rows.Close()
			return
		}

		tmps = append(tmps, tmp)

	}

	log.Println(tmps)
}

func TestCallSP(t *testing.T) {
	// database/sql
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}

	userID := 1008
	lock := new(sync.Mutex)

	/*
	  "public"."usp_game_users_stat"(
	    "_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_game_id" int4, "_de" float8,
	    "_ya" float8, "_vaild_ya" float8, "_tax" float8, "_bonus" float8, "_play_count" int4,
	    "_big_win_count" int4, "_win_count" int4, "_lose_count" int4, "_last_bet_time" timestamptz
	  )
	*/
	query := `CALL "public"."usp_game_users_stat"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	for i := 0; i < 100; i++ {
		go func(idx int, lock *sync.Mutex) {

			lock.Lock()
			defer lock.Unlock()

			_, err := db.Exec(query,
				1001, "0001", userID+idx, 1001, 20,
				10, 10, 1, 0, 1,
				0, 1, 0, time.Now())
			if err != nil {
				t.Logf("db err: %v", err)
			} else {
				t.Logf("close db connect: userID: %d", userID+i)
			}

			t.Logf("finish func")
		}(i, lock)
	}
	time.Sleep(15 * time.Second)
}

func checkFloatPlaces(d string, place int) (bool, error) {
	rebate, err := decimal.NewFromString(d)
	if err != nil {
		return false, err
	}
	if rebate.Round(int32(place)).String() != rebate.String() {
		return false, fmt.Errorf("格式错误，非小数点后2位小数")
	}

	return true, nil
}

func TestDecimal2(t *testing.T) {

	s := "100.820"
	p := 2
	t.Log(checkFloatPlaces(s, p))

	s = "100.82"
	t.Log(checkFloatPlaces(s, p))

	s = "100.8"
	t.Log(checkFloatPlaces(s, p))

	s = "100.823"
	t.Log(checkFloatPlaces(s, p))

}

func TestCommand(t *testing.T) {
	command := 0
	switch command {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		strMoney := "1000.55"
		t.Log(checkFloatPlaces(strMoney, 2))
	default:
		t.Log("nothing")
	}

}

func TestConnectRedis(t *testing.T) {
	rdb := redis.NewRedisCliectV9Test()

	for i := 0; i < 5; i++ {
		err := rdb.CreateRedisObject(i, "127.0.0.1:6379", "")
		if err != nil {
			t.Logf("Redis database ping error: %v", err)
		}
	}

	jsonStr, err := rdb.LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, "132")
	if err != nil {
		t.Logf("Redis database LoadHValue error: %v", err)
	}

	t.Logf("result: %v", jsonStr)
}

func TestReturnParams(t *testing.T) {

	type CoinOutChannelHandleResponseData struct {
		Account string  `json:"account"` // 會員帳號
		Code    int     `json:"code"`    // 錯誤碼
		Balance float64 `json:"balance"` // 最後餘額
	}

	thirdAccount := "123465789"

	returnData := CoinOutChannelHandleResponseData{Account: thirdAccount}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CoinOut, &returnData)

	returnData.Code = 999

	t.Log(returnTmp)

}

func TestCombinationSQL(t *testing.T) {
	gamenames := []string{
		"sangong",
		"rcfishing",
		"fantan",
		"hundredsicbo",
		"blackjack",
		"fruitslot",
		"cockfight",
		"dogracing",
		"rummy",
		"goldenflower",
		"pokdeng",
		"prawncrab",
		"colordisc",
		"baccarat",
		"texas",
		"bullbull"}

	sqlStr1 := `DELETE FROM user_play_log_%s WHERE bet_time < '2023-05-01';`
	// sqlStr2 := `CREATE INDEX "idx_user_play_log_%s_lognumber_level_code" ON "public"."user_play_log_%s" USING btree (
	// "lognumber",
	// "level_code"
	// );`

	res := make([]string, 0)

	for _, v := range gamenames {
		sss1 := fmt.Sprintf(sqlStr1, v)
		// sss2 := fmt.Sprintf(sqlStr2, v, v)
		res = append(res, sss1)
		// res = append(res, sss2)
	}
	log.Println(res)
}
