package global

import (
	"backend/pkg/redis"
	"backend/pkg/utils"
	"encoding/json"
	"strconv"
)

type ExchangeData struct {
	Id       int     `json:"id"`
	Currency string  `json:"currency"` // pkey
	ToCoin   float64 `json:"to_coin"`
}

func NewGlobalExchangeDataCache(rdb redis.IRedisCliect, idx int, hashName string) *GlobalExchangeDataCache {
	return &GlobalExchangeDataCache{
		rdb:           rdb,
		rdb_idx:       idx,
		rdb_hash_name: hashName,
	}
}

type GlobalExchangeDataCache struct {
	rdb_idx       int
	rdb_hash_name string
	rdb           redis.IRedisCliect
}

// param game_user_id
func (p *GlobalExchangeDataCache) Get(currency string) *ExchangeData {
	ret := new(ExchangeData)
	jsonStr, err := p.rdb.LoadHValue(p.rdb_idx, p.rdb_hash_name, currency)
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), ret)
	}

	if ret != nil && ret.Id > 0 && ret.Currency != "" && ret.ToCoin > 0 && err == nil {
		return ret
	}

	return nil
}

func (p *GlobalExchangeDataCache) GetAll() []*ExchangeData {
	exchangeData := make([]*ExchangeData, 0)
	kvPairs, err := p.rdb.LoadHAllValue(p.rdb_idx, p.rdb_hash_name)
	if err == nil {
		for _, jsonStr := range kvPairs {
			ret := new(ExchangeData)
			err = json.Unmarshal([]byte(jsonStr), ret)
			if err != nil {
				continue
			}
			exchangeData = append(exchangeData, ret)
		}
		if len(exchangeData) > 0 {
			return exchangeData
		}
	}
	return nil
}

// func (p *GlobalExchangeDataCache) GetChildAgents(agentId int) []*ExchangeData {
// 	return nil
// }

func (p *GlobalExchangeDataCache) Add(e *ExchangeData) {
	p.rdb.StoreHValue(p.rdb_idx, p.rdb_hash_name, e.Currency, utils.ToJSON(e))
}

func (p *GlobalExchangeDataCache) Adds(es []*ExchangeData) {

	vals := make([]string, 0)
	for _, v := range es {
		vals = append(vals, v.Currency)
		vals = append(vals, utils.ToJSON(v))
	}
	p.rdb.StoreHValue(p.rdb_idx, p.rdb_hash_name, vals...)
}

func (p *GlobalExchangeDataCache) Remove(id int) {
	p.rdb.DeleteHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(id))
}

func (p *GlobalExchangeDataCache) RemoveAll() {
	p.rdb.Del(p.rdb_idx, p.rdb_hash_name)
}
