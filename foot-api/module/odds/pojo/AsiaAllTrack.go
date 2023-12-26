package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

/**
亚赔历史,变化过程表
*/
type AsiaAllTrack struct {
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

	//博彩公司id
	CompId   int    `xorm:"unique(CompId_MatchId_OddDate_Num)"`
	CompName string `xorm:"varchar(50) index"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId_OddDate_Num)  varchar(20)"`
	//数据时间
	OddDate string `xorm:"unique(CompId_MatchId_OddDate_Num)  varchar(20)"`
	OddsId  int    `xorm:"comment('OddsId')  varchar(20)"`
	Num     int    `xorm:"unique(CompId_MatchId_OddDate_Num)  index"`

	MainGuest string `xorm:"comment('Main_Guest_Name') varchar(20)"`

	pojo.BasePojo `xorm:"extends"`
}
