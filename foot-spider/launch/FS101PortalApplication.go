package launch

import (
	"fmt"
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/spider/constants"
	"time"
)

func Clean() {
	//清空数据表
	//Before_spider_match()
	//Before_spider_baseFace()
	Before_spider_asiaLast()
	Before_spider_euroLast()
}

func Spider() {

	mysql.ShowSQL(false)
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
	constants.FullSpiderDateStr = constants.SpiderDateStr
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	matchLevelStr := utils.GetVal("spider", "match_level")
	//if len(matchLevelStr) <= 0 {
	//	matchLevelStr = "4"
	//}
	fmt.Println(matchLevelStr)
	matchLevel, _ := strconv.Atoi(matchLevelStr)
	Spider_match(matchLevel)
	Spider_baseFace(false)

	/**
	根据比赛id,抓取亚赔数据，类型为 text，解析text获取赔率信息。
	更新亚赔last his 表数据，由于exist 进行find 时候这两个表进加入了 比赛id 和 菠菜id,故仅仅保留最后一条数据。
	亚赔数据表track 由于exist 进行find 时候这两个表进加入了odddata ，在插入数据的时候可以保留不同 odddata 数据。
	*/
	Spider_asiaLastNew(false)

	/**
	根据比赛id,抓取欧赔数据，类型为网页数据，不存在 odddata, 只包含即时赔率和初始赔率，
	保存在 last，his 表中， odddata 需要参考modifytime。
	*/
	Spider_euroLast()
	/**
	根据比赛id, 菠菜id 抓取欧赔数据，类型为网页数据，包含菠菜公司的历史数据，将数据存入track 表中。
	将菠菜公司初始赔率和即时赔率同时更新到 last,his 表中。
	*/
	Spider_euroHis()

	//再对欧赔数据不完整的比赛进行两次抓取
	Spider_euroHis_Incomplete()
}

func My_Spider() {

	mysql.ShowSQL(false)
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
	constants.FullSpiderDateStr = constants.SpiderDateStr
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	matchLevelStr := utils.GetVal("spider", "match_level")
	if len(matchLevelStr) <= 0 {
		matchLevelStr = "4"
	}
	matchLevel, _ := strconv.Atoi(matchLevelStr)
	Spider_match(matchLevel)
	Spider_asia_not_start()
}

func Spider_Near() {
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")

	matchLevelStr := utils.GetVal("spider", "match_level")
	if len(matchLevelStr) <= 0 {
		matchLevelStr = "4"
	}
	matchLevel, _ := strconv.Atoi(matchLevelStr)
	Spider_match(matchLevel)
	//基本面不会改变
	//Spider_baseFace_near()
	Spider_asiaLastNew_near()
	Spider_euroLast_near()
	Spider_euroHis_near()
}
