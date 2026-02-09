package model_test

import (
	"backend/server/table/model"
	"log"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	tmp := model.NewEmptyMarquee()

	tmp.StartTime = time.Now().UTC()
	tmp.EndTime = tmp.StartTime.AddDate(0, 0, 1)

	log.Printf("StartTime: %v", tmp.StartTime)
	log.Printf("EndTime: %v", tmp.EndTime)
	log.Printf("Compare: %v", tmp.StartTime.After(tmp.EndTime))
}
