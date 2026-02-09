package global

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// REFRESH MATERIALIZED VIEW
func ReFreshMVGameStatRecord(ctx context.Context, db *sql.DB, dataType string, startTime, endTime time.Time) {

	// combition table name
	mvTablename := "mv_cal_game_stat_" + dataType

	query := `REFRESH MATERIALIZED VIEW "public"."%s" WITH DATA;`

	// create table name
	finalQuery := fmt.Sprintf(query, mvTablename)
	// combination sql query end

	_, err := db.Exec(finalQuery)
	if err != nil {
		log.Printf("ReFreshMVGameStatRecord(), err is: %v", err)
	}
}
