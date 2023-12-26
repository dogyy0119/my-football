package service

import (
	"fmt"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaAllTrackService struct {
	mysql.BaseService
}

func (this *AsiaAllTrackService) Exist(v *pojo.AsiaAllTrack) (string, bool) {
	temp := &pojo.AsiaAllTrack{MatchId: v.MatchId, CompId: v.CompId, OddDate: v.OddDate, Num: v.Num}
	var id string
	exist, err := mysql.GetEngine().Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist {
		id = temp.Id
	}
	return id, exist
}

//根据比赛ID查找亚赔
func (this *AsiaAllTrackService) FindByMatchId(matchId string) []*pojo.AsiaAllTrack {
	dataList := make([]*pojo.AsiaAllTrack, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

func (this *AsiaAllTrackService) AnalyOneMatch(matchId string) ([]*pojo.AsiaAllTrack, int, int) {
	allList := this.FindByMatchId(matchId)
	normal, unnormal := this.AnalyData(allList)
	return allList, normal, unnormal
}

func (this *AsiaAllTrackService) AnalyData(allList []*pojo.AsiaAllTrack) (int, int) {
	if len(allList) < 1 {
		return -1, -1
	}
	_letball := allList[0].SPanKou

	if _letball > -2 && _letball < 2 {

	} else {
		fmt.Println("_letball 不在 -2 和 2 之间 matchId:" + allList[0].MatchId)
		return -1, -1
	}
	_normal := 0
	_unnormal := 0
	// 统计赔率正常、非正常个数
	for _, one := range allList {
		if one.EPanKou != _letball {
			return -1, -1
		}
		if one.EPanKou > 0 {
			if one.Ep0 > 0.88 || one.Ep0 < 1.05 {
				_normal = _normal + 1
			} else {
				_unnormal = _unnormal + 1
			}
		} else {
			if one.Ep3 > 0.88 || one.Ep3 < 1.05 {
				_normal = _normal + 1
			} else {
				_unnormal = _unnormal + 1
			}
		}
	}

	return _normal, _unnormal
}

//根据比赛ID和波菜公司ID查找亚赔
func (this *AsiaAllTrackService) FindByMatchIdCompId(matchId string, compIds ...string) []*pojo.AsiaAllTrack {
	dataList := make([]*pojo.AsiaAllTrack, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' AND CompId in ( '0' ")
	for _, v := range compIds {
		sql_build.WriteString(" ,'")
		sql_build.WriteString(v)
		sql_build.WriteString("'")
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchIdCompId:", err)
	}
	return dataList
}
