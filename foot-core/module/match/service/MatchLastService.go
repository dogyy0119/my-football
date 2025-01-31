package service

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//足球比赛信息
type MatchLastService struct {
	mysql.BaseService
}

/**
通过比赛时间,主队id,客队id,判断比赛信息是否已经存在
*/
func (this *MatchLastService) Exist(v *pojo.MatchLast) bool {
	has, err := mysql.GetEngine().Table("`t_match_last`").Where(" `Id` = ?  ", v.Id).Exist()

	temp := &pojo.MatchLast{MainTeamId: v.MainTeamId, GuestTeamId: v.GuestTeamId, MatchDate: v.MatchDate}
	//var id string
	exist, err := mysql.GetEngine().Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist || has {
		return true
	}
	return exist
}

func (this *MatchLastService) FindByMatchId(Id string) []*pojo.MatchLast {
	dataList := make([]*pojo.MatchLast, 0)                         // 初始化和分配内存空间
	err := mysql.GetEngine().Where(" Id = ? ", Id).Find(&dataList) // 将 dataList 的地址传递给 Find 方法
	if err != nil {
		base.Log.Error("FindById:", err)
	}
	return dataList
}

func (this *MatchLastService) FindAll() []*pojo.MatchLast {
	dataList := make([]*pojo.MatchLast, 0)
	mysql.GetEngine().OrderBy("MatchDate").Find(&dataList)
	return dataList
}

func (this *MatchLastService) FindReady() []*pojo.MatchLast {
	sql_build_1 := `
SELECT 
  la.* 
FROM
  t_match_last la 
WHERE DATE_ADD(la.MatchDate, INTERVAL 5 MINUTE) >= NOW() 
  AND la.MatchDate <= DATE_ADD(NOW(), INTERVAL 30 MINUTE)
	`
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build_1, &dataList)

	return dataList
}

/**
获取临场比赛
*/
func (this *MatchLastService) FindNear() []*pojo.MatchLast {
	sql_build_1 := `
SELECT 
  la.* 
FROM
  t_match_last la 
WHERE DATE_ADD(la.MatchDate, INTERVAL 5 MINUTE) >= NOW() 
  AND la.MatchDate <= DATE_ADD(NOW(), INTERVAL 30 MINUTE)
	`
	sql_build_2 := `
SELECT DISTINCT 
  la.* 
FROM
  t_match_last la,
  t_analy_result ar 
WHERE la.Id = ar.MatchId 
  AND DATE_ADD(la.MatchDate, INTERVAL 6 MINUTE) >= NOW() 
  AND la.MatchDate <= DATE_ADD(NOW(), INTERVAL 30 MINUTE)
  AND ar.AlFlag != 'C1'
	`
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build_1, &dataList)

	//如果数据量过多,则配置分析表重新获取...只默认只处理临场12场
	if len(dataList) <= 10 {
		return dataList
	}

	//结果值
	dataList = make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build_2, &dataList)

	return dataList
}

/**
查找未结束的比赛
*/
func (this *MatchLastService) FindNotFinished() []*pojo.MatchLast {
	sql_build := `SELECT 
    la.* 
FROM
    t_match_his la,
    t_league l 
WHERE la.LeagueId = l.Id 
    AND la.MatchDate > DATE_SUB(NOW(), INTERVAL 24 HOUR)
ORDER BY la.MatchDate ASC`
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build, &dataList)
	return dataList
}

/**
查找未开始的比赛
*/
func (this *MatchLastService) FindNotStart() []*pojo.MatchLast {
	sql_build := `
SELECT 
  la.* 
FROM
  t_match_his la,
  t_league l 
WHERE la.LeagueId = l.Id 
  AND la.MatchDate > NOW()
  AND la.MatchDate <= DATE_ADD(NOW(), INTERVAL 1 DAY)
ORDER BY la.MatchDate ASC
	`
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build, &dataList)
	return dataList
}

/**
查找欧赔不完整的比赛
*/
func (this *MatchLastService) FindEuroIncomplete(count int) []*pojo.MatchLast {
	sql_build := `
SELECT 
  la.* 
FROM
  t_euro_last l,
  t_match_last la 
WHERE l.MatchId = la.Id 
GROUP BY l.MatchId
	`
	sql_build += " HAVING COUNT(1) < " + strconv.Itoa(count)
	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build, &dataList)
	return dataList
}
