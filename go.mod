module github.com/go-mesh/registrator

require (
	github.com/DataDog/zstd v1.3.5 // indirect
	github.com/Shopify/sarama v1.20.1 // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/go-chassis/go-chassis v1.2.3-0.20190128100058-43e245df146c
	github.com/go-mesh/openlogging v0.0.0-20181122085847-3daf3ad8ed35
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/huaweicse/auth v0.0.0-20181213084859-bfbe63726167
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/urfave/cli v1.20.1-0.20181029213200-b67dcf995b6a
	google.golang.org/genproto v0.0.0-20181221175505-bd9b4fb69e2f // indirect
	gopkg.in/yaml.v2 v2.2.1

)

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/net v0.0.0-20180824152047-4bcd98cce591 => github.com/golang/net v0.0.0-20180824152047-4bcd98cce591
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys v0.0.0-20180824143301-4910a1d54f87 => github.com/golang/sys v0.0.0-20180824143301-4910a1d54f87
	golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/sys v0.0.0-20181116152217-5ac8a444bdc5 => github.com/golang/sys v0.0.0-20181116152217-5ac8a444bdc5
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	google.golang.org/genproto v0.0.0-20181114220301-adae6a3d119a => github.com/google/go-genproto v0.0.0-20181114220301-adae6a3d119a
	google.golang.org/genproto v0.0.0-20181221175505-bd9b4fb69e2f => github.com/google/go-genproto v0.0.0-20181221175505-bd9b4fb69e2f
	google.golang.org/grpc v1.16.0 => github.com/grpc/grpc-go v1.16.0
)
