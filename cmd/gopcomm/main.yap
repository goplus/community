import (
	"context"
	"flag"
	"github.com/goplus/community/internal/core"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qiniu/x/xlog"
	"net/http"
)

var (
	community *core.Community
	domain    string
)

xLog := xlog.New("")
todo := context.TODO()
static "/static"
configFile := ".env"
flag.StringVar(&configFile, "config", ".env", "Path to the config file")
flag.Parse()
if err := godotenv.Load(configFile); err != nil {
	xLog.Error(err)
}

conf := core.NewConfigFromEnv()
community, _ = core.New(todo, conf)
domain = conf.QiNiuConfig.Domain
endpoint := conf.AppConfig.EndPoint

xLog.Info "Started in endpoint: ", endpoint

run(endpoint, func(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {

			if err := recover(); err != nil {
				xLog.Error(err)
				http.Redirect(w, r, "/failed", http.StatusFound)
			}
		}()

		h.ServeHTTP(w, r)
	})
})
