module tesou.io/platform/foot-parent/foot-spider

require (
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/andybalholm/brotli v1.0.6
	github.com/bitly/go-simplejson v0.5.1
	github.com/hu17889/go_spider v0.0.0-20150809033053-85ede20bf88b
	//github.com/hu17889/go_spider v0.0.0-20150113074338-ffedfe4e0f3c
	golang.org/x/net v0.7.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.5 // indirect
	tesou.io/platform/foot-parent/foot-api v1.0.0
	//opensource.io/go_spider v1.0.0
	tesou.io/platform/foot-parent/foot-core v1.0.0
)

replace (
	github.com/go-xorm/core v0.6.3 => github.com/go-xorm/core v0.6.2
	//opensource.io/go_spider => ../../../../opensource.io/go_spider
	tesou.io/platform/foot-parent/foot-api => ../foot-api
	tesou.io/platform/foot-parent/foot-core => ../foot-core
)

go 1.13
