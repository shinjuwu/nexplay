package utils_test

import (
	"backend/pkg/utils"
	"log"
	"testing"
	"time"
)

func TestCreatreOrderIdByOrderTypeAndSalt(t *testing.T) {
	log.Println(utils.CreatreOrderIdByOrderTypeAndSalt(4, "123123", time.Now()))
}
