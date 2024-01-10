package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

/**
历史对战
*/
type BFBattle struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID') index"`

	/**
	 * 比赛时间
	 */
	BattleMatchDate time.Time `xorm:"unique(BattleMatchDate_MainTeamId_GuestTeamId) index"`
	/**
	 * 联赛Id
	 */
	BattleLeagueId string `xorm:" comment('联赛Id') index"`
	/**
	 * 主队id,目前为主队名称
	 */
	BattleMainTeamId   string `xorm:"unique(BattleMatchDate_MainTeamId_GuestTeamId) index"`
	BattleMainTeamName string `xorm:" comment('主队队名') index"`
	/**
	 * 主队进球数
	 */
	BattleMainTeamHalfGoals int `xorm:" comment('主队半场进球数') index"`
	BattleMainTeamGoals     int `xorm:" comment('主队进球数') index"`
	/**
	 * 客队id,目前为客队名称
	 */
	BattleGuestTeamId   string `xorm:"unique(BattleMatchDate_MainTeamId_GuestTeamId) index"`
	BattleGuestTeamName string `xorm:" comment('主队队名') index"`

	/**
	 * 客队进球数
	 */
	BattleGuestTeamHalfGoals int `xorm:" comment('客队半场进球数') index"`
	BattleGuestTeamGoals     int `xorm:" comment('客队进球数') index"`

	/**
	 * 主队红卡
	 */
	BattleMainTeamRedCard int `xorm:" comment('主队红卡数') index"`

	/**
	 * 客队红卡
	 */
	BattleGuestTeamRedCard int `xorm:" comment('客队红卡数') index"`

	/**
	 * 主队黄卡
	 */
	BattleMainTeamYellowCard int `xorm:" comment('主队黄卡数') index"`

	/**
	 * 客队黄卡
	 */
	BattleGuestTeamYellowCard int `xorm:" comment('客队黄卡数') index"`

	/**
	 * 主队角球
	 */
	BattleMainTeamCorner int `xorm:" comment('主队角球数') index"`

	/**
	 * 客队角球
	 */
	BattleGuestTeamCorner int `xorm:" comment('客队角球数') index"`

	/**
	 * 半场让球
	 */
	BattleHalfLetgoal float64 `xorm:" comment('半场让球') index"`

	/**
	 * 半场让球结果
	 */
	BattleHalfLetgoalResult string `xorm:" comment('半场让球结果') index"`

	/**
	 * 全场让球
	 */
	BattleLetgoal float64 `xorm:" comment('全场让球') index"`

	/**
	 * 全场让球结果
	 */
	BattleLetgoalResult string `xorm:" comment('全场让球结果') index"`

	/**
	 * 半场大球
	 */
	BattleHalfOu float64 `xorm:" comment('半场大球') index"`

	/**
	 * 半场大球结果
	 */
	BattleHalfOuResult string `xorm:" comment('半场大球结果') index"`

	/**
	 * 全场大球
	 */
	BattleOu float64 `xorm:" comment('全场大球') index"`

	/**
	 * 全场大球结果
	 */
	BattleOuResult string `xorm:" comment('全场大球结果') index"`

	pojo.BasePojo `xorm:"extends"`
}
