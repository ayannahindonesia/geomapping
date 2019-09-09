package tests

import (
	"asira_geomapping/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestClientCity(t *testing.T) {
	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	obj := auth.GET("/clientauth").
		Expect().
		Status(http.StatusOK).JSON().Object()

	admintoken := obj.Value("token").String().Raw()

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+admintoken)
	})
	// valid response of province
	obj = auth.GET("/client/provinsi/1/kota").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj = auth.GET("/client/kota/1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj = auth.GET("/client/kota/99999").
		Expect().
		Status(http.StatusNotFound).JSON().Object()
}
