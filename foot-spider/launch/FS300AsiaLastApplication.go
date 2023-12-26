package launch

import (
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//执行抓取亚赔数据
	Before_spider_asiaLast()
	Spider_asiaLast()
}*/

func Before_spider_asiaLast() {
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_asia_last"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
//该页面已经被球探网废弃
func Spider_asiaLast() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindNotFinished()

	processer := proc.GetAsiaLastProcesser()
	processer.MatchLastList = matchLasts
	processer.Startup()
}

func Spider_asiaLastNew(spiderAll bool) {
	matchLastService := new(service2.MatchLastService)
	var matchLasts []*pojo.MatchLast
	if spiderAll {
		matchLasts = matchLastService.FindAll()
	} else {
		matchLasts = matchLastService.FindNotFinished()
	}

	processer := proc.GetAsiaLastNewProcesser()
	processer.MatchLastList = matchLasts
	processer.SingleThread = true
	processer.Startup()
}

func Spider_asia_not_start() {
	matchLastService := new(service2.MatchLastService)
	var matchLasts []*pojo.MatchLast
	matchLasts = matchLastService.FindNotStart()

	processer := proc.GetAsiaAllTrackProcesser()
	processer.MatchLastList = matchLasts
	processer.SingleThread = true
	processer.Startup()
}

func Spider_asiaLastNew_his(season string) {
	matchLastService := new(service2.MatchHisService)
	var matchLasts []*pojo.MatchLast
	matchLasts = matchLastService.FindBySeason(season)

	processer := proc.GetAsiaLastNewProcesser()
	processer.MatchLastList = matchLasts
	processer.Startup()
}

func Spider_asiaLastNew_near() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindNear()

	processer := proc.GetAsiaLastNewProcesser()
	processer.MatchLastList = matchLasts
	processer.SingleThread = true
	processer.Startup()
}
