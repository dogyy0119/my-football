package launch

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

var (
	maxLetBall   = 1.0
	showSql      = false
	printOddData = false
)

func Analy_new(all bool) {
	mysql.ShowSQL(showSql)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------MyE3模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e3 := new(service.E4Service)
	e3.MaxLetBall = 2
	e3.PrintOddData = printOddData
	e3.Analy(all)

	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
	mysql.ShowSQL(all)
}

func Analy(all bool) {
	//关闭SQL输出
	mysql.ShowSQL(showSql)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Q1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	q1 := new(service.Q1Service)
	q1.MaxLetBall = maxLetBall
	q1.PrintOddData = printOddData
	q1.Analy(all)
	mysql.ShowSQL(showSql)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E3模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e3 := new(service.E3Service)
	e3.MaxLetBall = maxLetBall
	e3.PrintOddData = printOddData
	e3.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E2模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e2 := new(service.E2Service)
	e2.MaxLetBall = maxLetBall
	e2.PrintOddData = printOddData
	e2.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e1 := new(service.E1Service)
	e1.MaxLetBall = maxLetBall
	e1.PrintOddData = printOddData
	e1.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
	mysql.ShowSQL(all)
}

func Analy_Near() {
	//关闭SQL输出
	mysql.ShowSQL(showSql)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E2模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e2 := new(service.E2Service)
	e2.MaxLetBall = maxLetBall
	e2.PrintOddData = printOddData
	e2.Analy_Near()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e1 := new(service.E1Service)
	e1.MaxLetBall = maxLetBall
	e1.PrintOddData = printOddData
	e1.Analy_Near()
	mysql.ShowSQL(showSql)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
}
