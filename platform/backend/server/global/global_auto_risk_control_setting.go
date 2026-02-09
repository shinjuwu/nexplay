package global

import (
	"backend/pkg/redis"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"encoding/json"
	"time"
)

type GlobalAutoRiskControlSettingCache struct {
	rdb_idx      int
	rdb_key_name string
	rdb          redis.IRedisCliect
}

func NewGlobalAutoRiskControlSettingCache(rdb redis.IRedisCliect, idx int, keyName string) *GlobalAutoRiskControlSettingCache {
	return &GlobalAutoRiskControlSettingCache{
		rdb:          rdb,
		rdb_idx:      idx,
		rdb_key_name: keyName,
	}
}

func (p *GlobalAutoRiskControlSettingCache) Get() *table_model.AutoRiskControlSetting {
	ret := new(table_model.AutoRiskControlSetting)
	jsonStr, err := p.rdb.LoadValue(p.rdb_idx, p.rdb_key_name)
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), ret)
	}

	if ret != nil && err == nil {
		return ret
	}

	return nil
}

func (p *GlobalAutoRiskControlSettingCache) Add(data *table_model.AutoRiskControlSetting) error {
	return p.rdb.StoreValue(p.rdb_idx, p.rdb_key_name, utils.ToJSON(data), time.Duration(0))
}
