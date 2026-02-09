package squirrel_test

import (
	"fmt"
	"log"
	"testing"

	sq "github.com/Masterminds/squirrel"
)

type testUser struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"username" db:"un"`
	Nickname string `json:"nickname"`
	Coin     int    `json:"coin" db:"coin"`
}

func TestSquirrelModel(t *testing.T) {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	username := fmt.Sprintf("%%%s%%", "kinco")
	gameUsers := psql.Select("count(*)").From("game_users").
		Where(
			sq.And{
				sq.Eq{"agent_id": 1},
				sq.Like{"username": username},
				sq.GtOrEq{"start_time": 1},
				sq.LtOrEq{"end_time": 10},
			})

	query, args, _ := gameUsers.ToSql()
	log.Println(query) // SELECT count(*) FROM game_users WHERE (agent_id = $1 AND username = $2)
	log.Println(args)

}
