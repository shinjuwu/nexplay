package global

import (
	"backend/pkg/redis"
	"backend/pkg/utils"
	"fmt"
	"strconv"
	"time"
)

type GlobalAutoRiskControlStatCache struct {
	rdb_idx        int
	rdb_key_prefix string
	rdb            redis.IRedisCliect
}

func NewGlobalAutoRiskControlStatCache(rdb redis.IRedisCliect, idx int, keyPrefix string) *GlobalAutoRiskControlStatCache {
	return &GlobalAutoRiskControlStatCache{
		rdb:            rdb,
		rdb_idx:        idx,
		rdb_key_prefix: keyPrefix,
	}
}

// 統計玩家每秒API的Request次數
func (p *GlobalAutoRiskControlStatCache) IncrGameUserApiRequestCount(userId int) (int64, error) {
	key := fmt.Sprintf("%sGameUserApiRequestCount_%d", p.rdb_key_prefix, userId)
	expiration := time.Second
	return p.rdb.IncrWithExpire(p.rdb_idx, key, expiration)
}

// 統計玩家每分鐘上下分的Request次數
func (p *GlobalAutoRiskControlStatCache) IncrGameUserCoinInAndOutRequestCount(userId int) (int64, error) {
	key := fmt.Sprintf("%sGameUserCoinInAndOutRequestCount_%d", p.rdb_key_prefix, userId)
	expiration := time.Minute
	return p.rdb.IncrWithExpire(p.rdb_idx, key, expiration)
}

func (p *GlobalAutoRiskControlStatCache) getGameUserStatHash() string {
	return fmt.Sprintf("%sGameUserStat", p.rdb_key_prefix)
}

func (p *GlobalAutoRiskControlStatCache) getGameUserStatTotalCoinInFiled(userId int) string {
	return fmt.Sprintf("%d_TotalCoinIn", userId)
}

func (p *GlobalAutoRiskControlStatCache) getGameUserStatTotalCoinOutFiled(userId int) string {
	return fmt.Sprintf("%d_TotalCoinOut", userId)
}

func (p *GlobalAutoRiskControlStatCache) getGameUserStatTotalWinFiled(userId int) string {
	return fmt.Sprintf("%d_TotalWin", userId)
}

func (p *GlobalAutoRiskControlStatCache) getGameUserStatTotalGameFiled(userId int) string {
	return fmt.Sprintf("%d_TotalGame", userId)
}

// 統計玩家上分
func (p *GlobalAutoRiskControlStatCache) IncrGameUserTotalCoinIn(userId int, coinIn float64) (float64, error) {
	return p.rdb.HIncrByFloat(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalCoinInFiled(userId), coinIn)
}

// 統計玩家下分
func (p *GlobalAutoRiskControlStatCache) IncrGameUserTotalCoinOut(userId int, coinOut float64) (float64, error) {
	return p.rdb.HIncrByFloat(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalCoinOutFiled(userId), coinOut)
}

// 取得玩家上下分統計
func (p *GlobalAutoRiskControlStatCache) GetGameUserTotalCoinInAndTotalCoinOut(userId int) (totalCoinIn, totalCoinOut float64, err error) {
	hash := p.getGameUserStatHash()
	fileds := []string{p.getGameUserStatTotalCoinInFiled(userId), p.getGameUserStatTotalCoinOutFiled(userId)}

	results, err := p.rdb.LoadHValues(p.rdb_idx, hash, fileds)
	if err == nil {
		if results[0] != nil {
			totalCoinIn, _ = strconv.ParseFloat(results[0].(string), 64)
		}
		if results[1] != nil {
			totalCoinOut, _ = strconv.ParseFloat(results[1].(string), 64)
		}
	}

	return
}

// 清除玩家上分統計
func (p *GlobalAutoRiskControlStatCache) ClearGameUserTotalCoinIn(userId int) error {
	return p.rdb.DeleteHValue(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalCoinInFiled(userId))
}

// 清除玩家下分統計
func (p *GlobalAutoRiskControlStatCache) ClearGameUserTotalCoinOut(userId int) error {
	return p.rdb.DeleteHValue(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalCoinOutFiled(userId))
}

// 統計玩家勝場
func (p *GlobalAutoRiskControlStatCache) IncrGameUserTotalWin(userId int) (int64, error) {
	return p.rdb.HIncrBy(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalWinFiled(userId), 1)
}

// 統計玩家總場數
func (p *GlobalAutoRiskControlStatCache) IncrGameUserTotalGame(userId int) (int64, error) {
	return p.rdb.HIncrBy(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalGameFiled(userId), 1)
}

// 取得玩家勝場及總場數
func (p *GlobalAutoRiskControlStatCache) GetGameUserTotalWinAndTotalGame(userId int) (totalWin, totalGame int64, err error) {
	hash := p.getGameUserStatHash()
	fileds := []string{p.getGameUserStatTotalWinFiled(userId), p.getGameUserStatTotalGameFiled(userId)}

	results, err := p.rdb.LoadHValues(p.rdb_idx, hash, fileds)
	if err == nil {
		if results[0] != nil {
			totalWin = utils.ToInt64(results[0], 0)
		}
		if results[1] != nil {
			totalGame = utils.ToInt64(results[1], 0)
		}
	}

	return
}

// 清除玩家勝場統計
func (p *GlobalAutoRiskControlStatCache) ClearGameUserTotalWin(userId int) error {
	return p.rdb.DeleteHValue(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalWinFiled(userId))
}

// 清除玩家總場數統計
func (p *GlobalAutoRiskControlStatCache) ClearGameUserTotalGame(userId int) error {
	return p.rdb.DeleteHValue(p.rdb_idx, p.getGameUserStatHash(), p.getGameUserStatTotalGameFiled(userId))
}

// 清除所有玩家統計
func (p *GlobalAutoRiskControlStatCache) ClearAllStat() error {
	return p.rdb.DeleteValue(p.rdb_idx, p.getGameUserStatHash())
}
