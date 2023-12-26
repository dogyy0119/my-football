package proc

import (
	"encoding/json"
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"math/rand"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"

	entity2 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
	"time"
)

type AsiaAllTrackProcesser struct {
	service2.MatchLastService
	service.AsiaAllTrackService
	//是否是单线程
	SingleThread       bool
	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

func GetAsiaAllTrackProcesser() *AsiaAllTrackProcesser {
	processer := &AsiaAllTrackProcesser{}
	processer.Init()
	return processer
}

func (this *AsiaAllTrackProcesser) Init() {
	//初始化参数值
	this.Win007idMatchidMap = map[string]string{}
}

func (this *AsiaAllTrackProcesser) Setup(temp *AsiaAllTrackProcesser) {
	//设置参数值
}

func (this *AsiaAllTrackProcesser) Startup() {

	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "AsiaAllTrackProcesser")
	for i, v := range this.MatchLastList {

		if !this.SingleThread && i%1000 == 0 { //10000个比赛一个spider,一个赛季大概有30万场比赛,最多30spider
			//先将前面的spider启动
			newSpider.SetDownloader(down.NewMAsiaLastApiDownloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
			newSpider.SetThreadnum(10).Run()

			processer = GetAsiaAllTrackProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "AsiaAllTrackProcesser"+strconv.Itoa(i))
		}

		temp_flag := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(temp_flag)
		matchExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchExt)
		win007_id := matchExt.Sid

		processer.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_ASIAODD_NEW_URL_PATTERN, "${matchId}", win007_id, 1)
		url = strings.Replace(url, "${flesh}", strconv.FormatFloat(rand.Float64(), 'f', -1, 64), 1)
		newSpider = newSpider.AddUrl(url, "json")
	}

	newSpider.SetDownloader(down.NewMAsiaLastApiDownloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
	newSpider.SetThreadnum(1).Run()

}

func (this *AsiaAllTrackProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Error("URL:", request.Url, p.Errormsg())
		return
	}

	track_slice := make([]interface{}, 0)
	//track_update_slice := make([]interface{}, 0)
	hdata_str := p.GetBodyStr()
	if hdata_str == "" {
		base.Log.Error("hdata_str:为空,URL:", request.Url)
		return
	}

	asiaData := &vo.AsiaData{}
	json.Unmarshal([]byte(hdata_str), asiaData)

	matchId := this.Win007idMatchidMap[strconv.Itoa(asiaData.ScheduleID)]
	//没有数据,则返回
	if nil == asiaData.Companies || len(asiaData.Companies) <= 0 {
		return
	}
	for _, e := range asiaData.Companies {
		//fmt.Println("liuhang NameEn", e.NameEn)
		if e.NameEn != "Bet365" && e.NameEn != "Crown" {
			continue
		}

		//for _, odd := range e.Details {
		odd := e.Details[0]

		track := new(entity2.AsiaAllTrack)
		track.CompId = e.CompanyID
		track.CompName = e.NameCn
		track.MatchId = matchId

		track.Sp3 = odd.FirstHomeOdds
		track.SPanKou = odd.FirstDrawOdds
		track.Sp0 = odd.FirstAwayOdds
		track.Ep3 = odd.HomeOdds
		track.EPanKou = odd.DrawOdds
		track.Ep0 = odd.AwayOdds
		track.OddsId = odd.OddsID

		if len(odd.ModifyTime) > 0 {
			tempMt, err := strconv.ParseInt(odd.ModifyTime, 0, 64)
			if nil != err {
				base.Log.Error(err.Error())
			}
			track.OddDate = time.Unix(tempMt, 0).Format("2006-01-02 15:04:05")
		}

		track_temp_id, track_exists := this.AsiaAllTrackService.Exist(track)

		fmt.Println("liuhang track_exists :", track_exists)
		fmt.Println("liuhang track.OddDate :", track.OddDate)
		fmt.Println("liuhang odd NameEn", e.NameEn)
		fmt.Println("liuhang odd.HomeOdds", odd.HomeOdds)
		fmt.Println("liuhang odd.DrawOdds", odd.DrawOdds)
		fmt.Println("liuhang odd.AwayOdds", odd.AwayOdds)

		if !track_exists {
			//fmt.Println("liuhang !track_exists OddsId:", track.OddsId)
			//fmt.Println("liuhang matchLast matchId :", matchId)
			matchLast := this.MatchLastService.FindByMatchId(matchId)
			//fmt.Println("liuhang matchLast len :", len(matchLast))
			if len(matchLast) > 0 {
				track.MainGuest = matchLast[0].MainTeamId + "vs" + matchLast[0].GuestTeamId
			}

			track_slice = append(track_slice, track)
		} else {
			// 发现已经存在则不处理。
			track.Id = track_temp_id
			//track_update_slice = append(track_update_slice, track)
		}
		//}

	}

	this.AsiaAllTrackService.SaveList(track_slice)
	//this.AsiaAllTrackService.ModifyList(track_update_slice)
}

func (this *AsiaAllTrackProcesser) Finish() {
	base.Log.Info("亚赔抓取解析完成 \r\n")
}
