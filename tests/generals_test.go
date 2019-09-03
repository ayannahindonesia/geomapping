package tests

import (
	"asira_geomapping/asira"
	"asira_geomapping/migration"
	"fmt"
	"net/http"
	"os"

	"github.com/gavv/httpexpect"
)

var (
	clientBasicToken string = "Y2xpZW50OmNsaWVudGdlbw=="
	adminBasicToken  string = "YWRtaW46YWRtaW5rZXk="
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
