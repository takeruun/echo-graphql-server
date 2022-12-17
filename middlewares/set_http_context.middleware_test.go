package middlewares_test

import (
	"app/helpers"
	"app/middlewares"
	"app/test_utils"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

// HTTP の context が取得されているかのテスト
// 公式参考 : https://github.com/go-chi/chi/blob/master/mux_test.go#L1697

func TestHttpContext(t *testing.T) {
	e := echo.New()
	e.Use(echo.WrapMiddleware(middlewares.SetHttpContextMiddleware()))

	e.GET("/", echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Context().Value(http.ServerContextKey).(*http.Server); !ok {
			panic("missing server context")
		}

		if _, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); !ok {
			panic("missing local addr context")
		}

		if _, ok := r.Context().Value(helpers.HTTPKey("http")).(helpers.HTTP); !ok {
			panic("missing http context")
		}
	})))

	ts := httptest.NewUnstartedServer(e)
	ts.Start()

	defer ts.Close()

	test_utils.TestRequest(t, ts, "GET", "/", nil)
}
