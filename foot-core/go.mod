module tesou.io/platform/foot-parent/foot-core

require (
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/astaxie/beego v1.12.0
	github.com/chanxuehong/wechat v0.0.0-20190521093015-fafb751f9916 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.3
	github.com/go-xorm/xorm v0.7.9
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.51.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	tesou.io/platform/foot-parent/foot-api v1.0.0

)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	tesou.io/platform/foot-parent/foot-api => ../foot-api
)

go 1.13
