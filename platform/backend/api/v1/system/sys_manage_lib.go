package system

import "fmt"

/*
SELECT date_trunc('hour'::text, tmp.bet_time) AS log_time,
    tmp.level_code,
    tmp.game_id,
    count(DISTINCT tmp.user_id) AS bet_user,
    count(tmp.lognumber) AS bet_count,
    sum(tmp.ya_score) AS sum_ya,
    sum(tmp.valid_score) AS sum_valid_ya,
    sum(tmp.de_score) AS sum_de,
    0 AS sum_bonus,
    sum(tmp.tax) AS sum_tax
   FROM ( SELECT user_play_log_baccarat.lognumber,
            user_play_log_baccarat.level_code,
            user_play_log_baccarat.user_id,
            user_play_log_baccarat.game_id,
            user_play_log_baccarat.ya_score,
            user_play_log_baccarat.valid_score,
            user_play_log_baccarat.de_score,
            user_play_log_baccarat.tax,
            user_play_log_baccarat.bonus,
            user_play_log_baccarat.bet_time
           FROM user_play_log_baccarat
          WHERE ((user_play_log_baccarat.is_robot = 0) AND (user_play_log_baccarat.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_colordisc.lognumber,
            user_play_log_colordisc.level_code,
            user_play_log_colordisc.user_id,
            user_play_log_colordisc.game_id,
            user_play_log_colordisc.ya_score,
            user_play_log_colordisc.valid_score,
            user_play_log_colordisc.de_score,
            user_play_log_colordisc.tax,
            user_play_log_colordisc.bonus,
            user_play_log_colordisc.bet_time
           FROM user_play_log_colordisc
          WHERE ((user_play_log_colordisc.is_robot = 0) AND (user_play_log_colordisc.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_fantan.lognumber,
            user_play_log_fantan.level_code,
            user_play_log_fantan.user_id,
            user_play_log_fantan.game_id,
            user_play_log_fantan.ya_score,
            user_play_log_fantan.valid_score,
            user_play_log_fantan.de_score,
            user_play_log_fantan.tax,
            user_play_log_fantan.bonus,
            user_play_log_fantan.bet_time
           FROM user_play_log_fantan
          WHERE ((user_play_log_fantan.is_robot = 0) AND (user_play_log_fantan.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_prawncrab.lognumber,
            user_play_log_prawncrab.level_code,
            user_play_log_prawncrab.user_id,
            user_play_log_prawncrab.game_id,
            user_play_log_prawncrab.ya_score,
            user_play_log_prawncrab.valid_score,
            user_play_log_prawncrab.de_score,
            user_play_log_prawncrab.tax,
            user_play_log_prawncrab.bonus,
            user_play_log_prawncrab.bet_time
           FROM user_play_log_prawncrab
          WHERE ((user_play_log_prawncrab.is_robot = 0) AND (user_play_log_prawncrab.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_blackjack.lognumber,
            user_play_log_blackjack.level_code,
            user_play_log_blackjack.user_id,
            user_play_log_blackjack.game_id,
            user_play_log_blackjack.ya_score,
            user_play_log_blackjack.valid_score,
            user_play_log_blackjack.de_score,
            user_play_log_blackjack.tax,
            user_play_log_blackjack.bonus,
            user_play_log_blackjack.bet_time
           FROM user_play_log_blackjack
          WHERE ((user_play_log_blackjack.is_robot = 0) AND (user_play_log_blackjack.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_sangong.lognumber,
            user_play_log_sangong.level_code,
            user_play_log_sangong.user_id,
            user_play_log_sangong.game_id,
            user_play_log_sangong.ya_score,
            user_play_log_sangong.valid_score,
            user_play_log_sangong.de_score,
            user_play_log_sangong.tax,
            user_play_log_sangong.bonus,
            user_play_log_sangong.bet_time
           FROM user_play_log_sangong
          WHERE ((user_play_log_sangong.is_robot = 0) AND (user_play_log_sangong.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_hundredsicbo.lognumber,
            user_play_log_hundredsicbo.level_code,
            user_play_log_hundredsicbo.user_id,
            user_play_log_hundredsicbo.game_id,
            user_play_log_hundredsicbo.ya_score,
            user_play_log_hundredsicbo.valid_score,
            user_play_log_hundredsicbo.de_score,
            user_play_log_hundredsicbo.tax,
            user_play_log_hundredsicbo.bonus,
            user_play_log_hundredsicbo.bet_time
           FROM user_play_log_hundredsicbo
          WHERE ((user_play_log_hundredsicbo.is_robot = 0) AND (user_play_log_hundredsicbo.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_bullbull.lognumber,
            user_play_log_bullbull.level_code,
            user_play_log_bullbull.user_id,
            user_play_log_bullbull.game_id,
            user_play_log_bullbull.ya_score,
            user_play_log_bullbull.valid_score,
            user_play_log_bullbull.de_score,
            user_play_log_bullbull.tax,
            user_play_log_bullbull.bonus,
            user_play_log_bullbull.bet_time
           FROM user_play_log_bullbull
          WHERE ((user_play_log_bullbull.is_robot = 0) AND (user_play_log_bullbull.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_texas.lognumber,
            user_play_log_texas.level_code,
            user_play_log_texas.user_id,
            user_play_log_texas.game_id,
            user_play_log_texas.ya_score,
            user_play_log_texas.valid_score,
            user_play_log_texas.de_score,
            user_play_log_texas.tax,
            user_play_log_texas.bonus,
            user_play_log_texas.bet_time
           FROM user_play_log_texas
          WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_fruitslot.lognumber,
            user_play_log_fruitslot.level_code,
            user_play_log_fruitslot.user_id,
            user_play_log_fruitslot.game_id,
            user_play_log_fruitslot.ya_score,
            user_play_log_fruitslot.valid_score,
            user_play_log_fruitslot.de_score,
            user_play_log_fruitslot.tax,
            user_play_log_fruitslot.bonus,
            user_play_log_fruitslot.bet_time
           FROM user_play_log_fruitslot
          WHERE ((user_play_log_fruitslot.is_robot = 0) AND (user_play_log_fruitslot.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_rcfishing.lognumber,
            user_play_log_rcfishing.level_code,
            user_play_log_rcfishing.user_id,
            user_play_log_rcfishing.game_id,
            user_play_log_rcfishing.ya_score,
            user_play_log_rcfishing.valid_score,
            user_play_log_rcfishing.de_score,
            user_play_log_rcfishing.tax,
            user_play_log_rcfishing.bonus,
            user_play_log_rcfishing.bet_time
           FROM user_play_log_rcfishing
          WHERE ((user_play_log_rcfishing.is_robot = 0) AND (user_play_log_rcfishing.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_cockfight.lognumber,
            user_play_log_cockfight.level_code,
            user_play_log_cockfight.user_id,
            user_play_log_cockfight.game_id,
            user_play_log_cockfight.ya_score,
            user_play_log_cockfight.valid_score,
            user_play_log_cockfight.de_score,
            user_play_log_cockfight.tax,
            user_play_log_cockfight.bonus,
            user_play_log_cockfight.bet_time
           FROM user_play_log_cockfight
          WHERE ((user_play_log_cockfight.is_robot = 0) AND (user_play_log_cockfight.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_dogracing.lognumber,
            user_play_log_dogracing.level_code,
            user_play_log_dogracing.user_id,
            user_play_log_dogracing.game_id,
            user_play_log_dogracing.ya_score,
            user_play_log_dogracing.valid_score,
            user_play_log_dogracing.de_score,
            user_play_log_dogracing.tax,
            user_play_log_dogracing.bonus,
            user_play_log_dogracing.bet_time
           FROM user_play_log_dogracing
          WHERE ((user_play_log_dogracing.is_robot = 0) AND (user_play_log_dogracing.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_rummy.lognumber,
            user_play_log_rummy.level_code,
            user_play_log_rummy.user_id,
            user_play_log_rummy.game_id,
            user_play_log_rummy.ya_score,
            user_play_log_rummy.valid_score,
            user_play_log_rummy.de_score,
            user_play_log_rummy.tax,
            user_play_log_rummy.bonus,
            user_play_log_rummy.bet_time
           FROM user_play_log_rummy
          WHERE ((user_play_log_rummy.is_robot = 0) AND (user_play_log_rummy.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_goldenflower.lognumber,
            user_play_log_goldenflower.level_code,
            user_play_log_goldenflower.user_id,
            user_play_log_goldenflower.game_id,
            user_play_log_goldenflower.ya_score,
            user_play_log_goldenflower.valid_score,
            user_play_log_goldenflower.de_score,
            user_play_log_goldenflower.tax,
            user_play_log_goldenflower.bonus,
            user_play_log_goldenflower.bet_time
           FROM user_play_log_goldenflower
          WHERE ((user_play_log_goldenflower.is_robot = 0) AND (user_play_log_goldenflower.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_pokdeng.lognumber,
            user_play_log_pokdeng.level_code,
            user_play_log_pokdeng.user_id,
            user_play_log_pokdeng.game_id,
            user_play_log_pokdeng.ya_score,
            user_play_log_pokdeng.valid_score,
            user_play_log_pokdeng.de_score,
            user_play_log_pokdeng.tax,
            user_play_log_pokdeng.bonus,
            user_play_log_pokdeng.bet_time
           FROM user_play_log_pokdeng
          WHERE ((user_play_log_pokdeng.is_robot = 0) AND (user_play_log_pokdeng.bet_time > (now() - '3 days'::interval)))) tmp
  GROUP BY (date_trunc('hour'::text, tmp.bet_time)), tmp.level_code, tmp.game_id
*/
func generateMVCalGameStatHourSQL(conditionQueryPart string, gameName []string) string {
	preQuery :=
		`SELECT date_trunc('hour'::text, tmp.bet_time) AS log_time,
		tmp.level_code,
		tmp.game_id,
		count(DISTINCT tmp.user_id) AS bet_user,
		count(tmp.lognumber) AS bet_count,
		sum(tmp.ya_score) AS sum_ya,
		sum(tmp.valid_score) AS sum_valid_ya,
		sum(tmp.de_score) AS sum_de,
		sum(tmp.bonus) AS sum_bonus,
		sum(tmp.tax) AS sum_tax
	FROM ( %s ) tmp
	GROUP BY (date_trunc('hour'::text, tmp.bet_time)), tmp.level_code, tmp.game_id`

	tableQuery := `SELECT lognumber,
					level_code,
					user_id,
					game_id,
					ya_score,
					valid_score,
					de_score,
					tax,
					bonus,
					bet_time
				FROM %s
				WHERE %s`
	unionQuery := " UNION ALL "

	fromQuery := ""
	// arges := make([]any, 0)
	for k, v := range gameName {
		tablename := "user_play_log_" + v
		fromQuery = fromQuery + fmt.Sprintf(tableQuery, tablename, conditionQueryPart)

		if len(gameName)-1 > k {
			fromQuery = fromQuery + unionQuery
		}
	}

	query := fmt.Sprintf(preQuery, fromQuery)

	return query
}
