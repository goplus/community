import (
	"context"
	"flag"
	"net/http"

	"github.com/goplus/community/internal/core"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qiniu/x/log"
)

const (
	limitConst      = 10
	labelConst      = "article"
	mediaLimitConst = 8
	firstConst      = "1"
)

var (
	community *core.Community
	domain    string
)

static "/static"
configFile := ".env"
flag.StringVar(&configFile, "config", ".env", "Path to the config file")
flag.Parse()
err := godotenv.Load(configFile)
if err != nil {
	log.Error(err)
}

conf := core.NewConfigFromEnv()
community, err = core.New(context.TODO(), conf)
if err != nil {
	log.Error(err)
}
domain = conf.QiNiuConfig.Domain
endpoint := conf.AppConfig.EndPoint

log.Info "Started in endpoint: ", endpoint

run(endpoint, func(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			if err := recover(); err != nil {
				log.Error(err)
				http.Redirect(w, r, "/failed", http.StatusFound)
			}
		}()

		h.ServeHTTP(w, r)
	})
})
