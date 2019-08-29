package tests

import (
	"fmt"
	"geomapping/asira"
	"geomapping/migration"
	"net/http"
	"os"

	"github.com/gavv/httpexpect"
)

var (
	clientBasicToken string = "cmVhY3RrZXk6cmVhY3RwYXNz"
	adminBasicToken  string = "Z3JhZGlvczp1bHRpbXVz"
)

func init() {
	// restrict test to development environment only.
	if asira.App.ENV != "development" {
		fmt.Printf("test aren't allowed in %s environment.", asira.App.ENV)
		os.Exit(1)
	}
}

func RebuildData() {
	migration.Truncate([]string{"all"})
	migration.Seed()
}

func getAdminToken(e *httpexpect.Expect, auth *httpexpect.Expect) string {
	obj := auth.GET("/clientauth").
		Expect().
		Status(http.StatusOK).JSON().Object()

	admintoken := obj.Value("token").String().Raw()

	return admintoken
}
