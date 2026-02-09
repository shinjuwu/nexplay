package global

import (
	"backend/internal/notification"
	"backend/pkg/logger"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"time"
)

var (
	resendMonitorServiceDataTimeInterval       = 1 * time.Minute
	resendMonitorServiceDataCount        int64 = 100
)

func T_ResendMonitorServiceData(logger *logger.RuntimeGoLogger) {
	for {
		_execResendMonitorServiceData(logger)
		time.Sleep(resendMonitorServiceDataTimeInterval)
	}
}

func _execResendMonitorServiceData(logger *logger.RuntimeGoLogger) {
	// 檢測時間
	now := time.Now().UnixNano()
	ftTimeNow := now / int64(time.Millisecond)

	logger.Printf("_execResendMonitorServiceData run, ftTimeNow: %v", ftTimeNow)

	if MonitorServiceCache == nil {
		return
	}

	addr, ok := ServerInfoCache.Load("monitor")
	if !ok {
		return
	}
	addres, ok := addr.(table_model.ServerInfo)
	if !ok {
		return
	}
	conninfo := utils.ToMap(addres.AddressesBytes)

	resendCollectorAbnormalWinAndLoseData(logger, conninfo)
	resendCollectorRTPStatData(logger, conninfo)
	resendCollectorCoinInOutData(logger, conninfo)
}

func resendCollectorAbnormalWinAndLoseData(logger *logger.RuntimeGoLogger, conninfo map[string]interface{}) {
	awls, err := MonitorServiceCache.GetSendCollectorAbnormalWinAndLoseData(resendMonitorServiceDataCount)
	if err != nil {
		return
	}

	if len(awls) > 0 {
		resendCount := int64(0)
		for _, awl := range awls {
			ok, _ := notification.SendCollectorAbnormalWinAndLoseToMonitorService(conninfo, awl)
			if !ok {
				break
			}
			resendCount++
		}

		_, err := MonitorServiceCache.DeleteSendCollectorAbnormalWinAndLoseData(resendCount)
		if err != nil {
			logger.Error("ResendCollectorAbnormalWinAndLoseToMonitorService success but delete cache data fail, data:%v, resend count:%d", awls, resendCount)
		}
	}
}

func resendCollectorRTPStatData(logger *logger.RuntimeGoLogger, conninfo map[string]interface{}) {
	rtpss, err := MonitorServiceCache.GetSendCollectorRTPStatData(resendMonitorServiceDataCount)
	if err != nil {
		return
	}

	if len(rtpss) > 0 {
		resendCount := int64(0)
		for _, rtps := range rtpss {
			ok, _ := notification.SendCollectorRTPStatToMonitorService(conninfo, rtps)
			if !ok {
				break
			}
			resendCount++
		}

		_, err := MonitorServiceCache.DeleteSendCollectorRTPStatData(resendCount)
		if err != nil {
			logger.Error("SendCollectorRTPStatToMonitorService success but delete cache data fail, data:%v, resend count:%d", rtpss, resendCount)
		}
	}
}

func resendCollectorCoinInOutData(logger *logger.RuntimeGoLogger, conninfo map[string]interface{}) {
	cios, err := MonitorServiceCache.GetSendCollectorCoinInOutData(resendMonitorServiceDataCount)
	if err != nil {
		return
	}

	if len(cios) > 0 {
		resendCount := int64(0)
		for _, cio := range cios {
			ok, _ := notification.SendCollectorCoinInOutToMonitorService(conninfo, cio)
			if !ok {
				break
			}
			resendCount++
		}

		_, err := MonitorServiceCache.DeleteSendCollectorCoinInOutData(resendCount)
		if err != nil {
			logger.Error("SendCollectorCoinInOutToMonitorService success but delete cache data fail, data:%v, resend count:%d", cios, resendCount)
		}
	}
}
