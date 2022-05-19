package downtime

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviderFactories map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"downtimes": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func testAccPreCheck(t *testing.T) {

	if value := os.Getenv("OBM_BA_USER"); value == "" {
		t.Fatal("OBM_BA_USER must be set for acceptance tests")
	}
	if value := os.Getenv("OBM_BA_PASSWORD"); value == "" {
		t.Fatal("OBM_BA_PASSWORD must be set for acceptance tests")
	}
}
