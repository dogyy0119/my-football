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
	analyService := new(service.MyAnalyService)
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

	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------A1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	a1 := new(service.A1Service)
	a1.MaxLetBall = maxLetBall
	a1.PrintOddData = printOddData
	a1.Analy(all)

	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------A2模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	a2 := new(service.A2Service)
	a2.MaxLetBall = maxLetBall
	a2.PrintOddData = printOddData
	a2.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------A3模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	a3 := new(service.A3Service)
	a3.MaxLetBall = maxLetBall
	a3.PrintOddData = printOddData
	a3.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------C1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	c1 := new(service.C1Service)
	c1.MaxLetBall = maxLetBall
	c1.PrintOddData = printOddData
	c1.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------C2模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	c2 := new(service.C2Service)
	c2.MaxLetBall = maxLetBall
	c2.PrintOddData = printOddData
	c2.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------C3模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	c3 := new(service.C3Service)
	c3.MaxLetBall = maxLetBall
	c3.PrintOddData = printOddData
	c3.Analy(all)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------C4模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	c4 := new(service.C4Service)
	c4.MaxLetBall = maxLetBall
	c4.PrintOddData = printOddData
	c4.Analy(all)

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
