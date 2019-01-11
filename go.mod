module github.com/go-mesh/registrator

require (
	github.com/Shopify/sarama v1.20.1 // indirect
	github.com/go-chassis/go-chassis v1.2.2-0.20190111015835-0b8d1a199be8
	github.com/go-chassis/go-sc-client v0.0.0-20190110124355-7e78d2170dca
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/go-mesh/openlogging v0.0.0-20181122085847-3daf3ad8ed35
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/huaweicse/auth v0.0.0-20181213084859-bfbe63726167
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/urfave/cli v1.20.1-0.20181029213200-b67dcf995b6a
	gopkg.in/yaml.v2 v2.2.1

)

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180824152047-4bcd98cce591 => github.com/golang/net v0.0.0-20180824152047-4bcd98cce591
	golang.org/x/sys v0.0.0-20180824143301-4910a1d54f87 => github.com/golang/sys v0.0.0-20180824143301-4910a1d54f87
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2 => github.com/golang/time v0.0.0-20180412165947-fbb02b2291d2
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.14.0 => github.com/grpc/grpc-go v1.14.0
)
