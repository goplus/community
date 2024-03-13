module github.com/goplus/community

go 1.19

require (
	github.com/go-sql-driver/mysql v1.8.0
	github.com/google/uuid v1.6.0
	github.com/goplus/yap v0.7.2 //gop:class
	github.com/qiniu/go-sdk/v7 v7.19.0
	github.com/yuin/goldmark v1.7.0
	gocloud.dev v0.37.0
	golang.org/x/text v0.14.0
)

replace github.com/goplus/yap v0.7.2 => github.com/LiusCraft/yap v0.8.6

require (
	github.com/casdoor/casdoor-go-sdk v0.35.1
	github.com/qiniu/go-cdk-driver v0.1.0
	github.com/qiniu/x v1.13.8
	golang.org/x/oauth2 v0.18.0
)

require (
	github.com/liuscraft/gop-casdoor-account-sdk v1.1.0
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/stretchr/testify v1.9.0
)

require github.com/gabriel-vasile/mimetype v1.4.3

replace github.com/qiniu/go-cdk-driver v0.1.0 => github.com/xhyqaq/go-cdk-driver v1.0.0

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/googleapis/gax-go/v2 v2.12.2 // indirect
	github.com/joho/godotenv v1.5.1
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.169.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240311173647-c811ad7063a7 // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
