package tests

import (
	"geomapping/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestClientVillage(t *testing.T) {
	RebuildData()
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
	// valid response of kelurahan
	obj = auth.GET("/client/provinsi/kota/kecamatan/1/kelurahan").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj = auth.GET("/client/kelurahan/1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj = auth.GET("/client/kelurahan/99999").
		Expect().
		Status(http.StatusNotFound).JSON().Object()
}
