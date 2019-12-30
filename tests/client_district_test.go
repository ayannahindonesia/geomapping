package tests

import (
	"geomapping/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestClientDistrict(t *testing.T) {
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
	// valid response of province
	obj = auth.GET("/client/provinsi/kota/1/kecamatan").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj = auth.GET("/client/kecamatan/1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj = auth.GET("/client/kecamatan/99999").
		Expect().
		Status(http.StatusNotFound).JSON().Object()
}
