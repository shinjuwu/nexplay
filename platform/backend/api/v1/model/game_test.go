package model_test

import (
	"backend/pkg/utils"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestLoadGameUser(t *testing.T) {

	// database/sql
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}

	selectLognumber := "100101-1563043497206878208"
	sql := fmt.Sprintf(`SELECT lognumber, playlog, game_id, room_type, bet_time FROM play_log_common WHERE lognumber='%s'`, selectLognumber)
	var jsonStr, lognumber string
	var gameId, roomType int
	var betTime time.Time

	db.QueryRow(sql).Scan(&lognumber, &jsonStr, &gameId, &roomType, &betTime)

	log.Println(jsonStr)

	if jsonStr == "" {
		log.Println("No data, change lognumber")
		return
	}

	tempMap := utils.ToMap([]byte(jsonStr))

	log.Println(tempMap)

	sqlStr := `INSERT INTO
		"public"."user_play_log_baccarat" ("lognumber", "agent_id", "user_id", "game_id", "room_type", 
		"desk_id", "seat_id", "exchange", "de_score","ya_score", 
		"valid_score", "start_score", "end_score", "create_time", "is_robot",
		 "is_big_win", "is_issue", "bet_time") VALUES `

	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),` //18

	playerDataTemp, ok := tempMap["playerlog"].([]interface{})
	if !ok {
		log.Println("playerData is not array")
	}
	dataCount := len(playerDataTemp)

	log.Printf("dataCount is %d", dataCount)

	vals := []interface{}{}

	idx := 1
	for _, val := range playerDataTemp {

		// "lognumber", "agent_id", "user_id", "game_id",
		// "room_type", "desk_id", "seat_id", "exchange", "de_score",
		// "ya_score", "valid_score", "start_score", "end_score", "create_time",
		// "is_robot", "is_big_win", "is_issue", "bet_time"
		valMap, ok := val.(map[string]interface{})
		if ok {
			sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4,
				idx+5, idx+6, idx+7, idx+8, idx+9,
				idx+10, idx+11, idx+12, idx+13, idx+14,
				idx+15, idx+16, idx+17)
			idx += 18
			vals = append(vals,
				lognumber, 1, valMap["user_id"], gameId,
				roomType, 100101, -1, 1, valMap["de_score"],
				valMap["ya_score"], valMap["valid_score"], valMap["start_score"], valMap["end_score"], time.Now().UTC(),
				valMap["is_robot"], false, false, betTime)
		}
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	log.Printf("finished sqlStr is %s", sqlStr)

	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		res, _ := stmt.Exec(vals...)
		log.Printf("%v", res)
		// t.Errorf("%v", res)
	} else {
		//
		log.Println("success")
	}

	// sqlStr := "INSERT INTO test(v1, v2, v3) VALUES "

	// vals := []interface{}{}

	// data := []map[string]string{
	// 	{"v1": "1", "v2": "1", "v3": "1"},
	// 	{"v1": "2", "v2": "2", "v3": "2"},
	// 	{"v1": "3", "v2": "3", "v3": "3"},
	// }

	// idx := 1
	// for _, row := range data {
	// 	sqlStr += fmt.Sprintf("($%d, $%d, $%d),", idx, idx+1, idx+2)
	// 	idx += 3
	// 	vals = append(vals, row["v1"], row["v2"], row["v3"])
	// }
	// //trim the last ,
	// sqlStr = sqlStr[0 : len(sqlStr)-1]
	// //prepare the statement
	// stmt, err := db.Prepare(sqlStr)
	// if err != nil {
	// 	t.Errorf("%v", err)
	// }
	// defer stmt.Close()

	// if stmt != nil {
	// 	//format all vals at once
	// 	res, _ := stmt.Exec(vals...)
	// 	t.Errorf("%v", res)
	// } else {
	// 	//
	// }

}
