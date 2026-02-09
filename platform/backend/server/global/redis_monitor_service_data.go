package global

import (
	"backend/internal/notification"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"encoding/json"
	"fmt"
)

func NewGlobalMonitorServiceDataCache(rdb redis.IRedisCliect, idx int, keyPrefix string) *GlobalMonitorServiceDataCache {
	return &GlobalMonitorServiceDataCache{
		rdb:     rdb,
		rdb_idx: idx,
		rdb_key_send_collector_abnormal_win_and_lose_data: fmt.Sprintf("%sSendCollectorAbnormalWinAndLoseData", keyPrefix),
		rdb_key_send_collector_rtp_stat_data:              fmt.Sprintf("%sSendCollectorRTPStatData", keyPrefix),
		rdb_key_send_collector_coin_in_out_data:           fmt.Sprintf("%sSendCollectorCoinInOutData", keyPrefix),
	}
}

type GlobalMonitorServiceDataCache struct {
	rdb_idx                                           int
	rdb                                               redis.IRedisCliect
	rdb_key_send_collector_abnormal_win_and_lose_data string
	rdb_key_send_collector_rtp_stat_data              string
	rdb_key_send_collector_coin_in_out_data           string
}

func (c *GlobalMonitorServiceDataCache) AddSendCollectorAbnormalWinAndLoseData(data *notification.TmpCollectorAbnormalWinAndLoseRequest) (int64, error) {
	vals := make([]interface{}, 0)
	vals = append(vals, utils.ToJSON(data))
	return c.rdb.StoreLValue(c.rdb_idx, c.rdb_key_send_collector_abnormal_win_and_lose_data, vals)
}

func (c *GlobalMonitorServiceDataCache) GetSendCollectorAbnormalWinAndLoseData(count int64) ([]*notification.TmpCollectorAbnormalWinAndLoseRequest, error) {
	results, err := c.rdb.LoadLValue(c.rdb_idx, c.rdb_key_send_collector_abnormal_win_and_lose_data, 0, count-1)
	if err != nil {
		return nil, err
	}

	resp := make([]*notification.TmpCollectorAbnormalWinAndLoseRequest, 0)
	for _, result := range results {
		var tmp notification.TmpCollectorAbnormalWinAndLoseRequest
		json.Unmarshal([]byte(result), &tmp)
		resp = append(resp, &tmp)
	}
	return resp, nil
}

func (c *GlobalMonitorServiceDataCache) DeleteSendCollectorAbnormalWinAndLoseData(count int64) (string, error) {
	return c.rdb.DeleteLValue(c.rdb_idx, c.rdb_key_send_collector_abnormal_win_and_lose_data, count, -1)
}

func (c *GlobalMonitorServiceDataCache) AddSendCollectorRTPStatData(data *notification.TmpCollectorRTPStatRequest) (int64, error) {
	vals := make([]interface{}, 0)
	vals = append(vals, utils.ToJSON(data))
	return c.rdb.StoreLValue(c.rdb_idx, c.rdb_key_send_collector_rtp_stat_data, vals)
}

func (c *GlobalMonitorServiceDataCache) GetSendCollectorRTPStatData(count int64) ([]*notification.TmpCollectorRTPStatRequest, error) {
	results, err := c.rdb.LoadLValue(c.rdb_idx, c.rdb_key_send_collector_rtp_stat_data, 0, count-1)
	if err != nil {
		return nil, err
	}

	resp := make([]*notification.TmpCollectorRTPStatRequest, 0)
	for _, result := range results {
		var tmp notification.TmpCollectorRTPStatRequest
		json.Unmarshal([]byte(result), &tmp)
		resp = append(resp, &tmp)
	}
	return resp, nil
}

func (c *GlobalMonitorServiceDataCache) DeleteSendCollectorRTPStatData(count int64) (string, error) {
	return c.rdb.DeleteLValue(c.rdb_idx, c.rdb_key_send_collector_rtp_stat_data, count, -1)
}

func (c *GlobalMonitorServiceDataCache) AddSendCollectorCoinInOutData(data *notification.TempCollectorCoinInOutRequest) (int64, error) {
	vals := make([]interface{}, 0)
	vals = append(vals, utils.ToJSON(data))
	return c.rdb.StoreLValue(c.rdb_idx, c.rdb_key_send_collector_coin_in_out_data, vals)
}

func (c *GlobalMonitorServiceDataCache) GetSendCollectorCoinInOutData(count int64) ([]*notification.TempCollectorCoinInOutRequest, error) {
	results, err := c.rdb.LoadLValue(c.rdb_idx, c.rdb_key_send_collector_coin_in_out_data, 0, count-1)
	if err != nil {
		return nil, err
	}

	resp := make([]*notification.TempCollectorCoinInOutRequest, 0)
	for _, result := range results {
		var tmp notification.TempCollectorCoinInOutRequest
		json.Unmarshal([]byte(result), &tmp)
		resp = append(resp, &tmp)
	}
	return resp, nil
}

func (c *GlobalMonitorServiceDataCache) DeleteSendCollectorCoinInOutData(count int64) (string, error) {
	return c.rdb.DeleteLValue(c.rdb_idx, c.rdb_key_send_collector_coin_in_out_data, count, -1)
}
