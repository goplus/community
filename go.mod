module github.com/goplus/community

go 1.19

require (
	github.com/go-sql-driver/mysql v1.8.0
	github.com/google/uuid v1.6.0
	github.com/goplus/yap v0.7.2 //gop:class
	github.com/qiniu/go-sdk/v7 v7.19.0
	github.com/yuin/goldmark v1.7.0
	gocloud.dev v0.36.0
	golang.org/x/text v0.14.0
)

replace github.com/goplus/yap v0.7.2 => github.com/LiusCraft/yap v0.8.6

require (
	github.com/casdoor/casdoor-go-sdk v0.35.1
	github.com/qiniu/go-cdk-driver v0.1.0
	github.com/qiniu/x v1.13.8
	golang.org/x/oauth2 v0.17.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.3
	github.com/liuscraft/gop-casdoor-account-sdk v1.1.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

replace github.com/qiniu/go-cdk-driver v0.1.0 => github.com/xhyqaq/go-cdk-driver v1.0.0

require (
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/joho/godotenv v1.5.1
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.151.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
