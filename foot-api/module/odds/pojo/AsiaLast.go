package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
亚赔当前表,仅初盘，即时盘
*/
type AsiaLast struct {
	Num int `xorm: comment('Num') index`
	//博彩公司id
	CompId int `xorm:"unique(CompId_MatchId)"`
	//CompNameEN string `xorm:"varchar(50) index"`

	CompName string `xorm:"varchar(50) index"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId) varchar(20)"`

	/**
	初主队盘口赔率
	*/
	Sp3 float64 `xorm:" comment('Sp3') index"`
	Sp0 float64 `xorm:" comment('Sp0') index"`
	//让球
	SPanKou float64 `xorm:" comment('s让球') index"`

	/**
	即时客队盘口赔率
	*/
	Ep3 float64 `xorm:" comment('Ep3') index"`
	Ep0 float64 `xorm:" comment('Ep0') index"`
	//让球
	EPanKou float64 `xorm:" comment('e让球') index"`

	//数据时间
	OddDate string `xorm:" comment('数据时间') index varchar(20)"`

	pojo.BasePojo `xorm:"extends"`
}
