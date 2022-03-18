package downtime

import (
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
var testAccProvider *schema.Provider
var testAccProviders map[string]*schema.Provider

func init() {
	LogMe("init", time.RFC3339)
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"downtimes": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	LogMe("TestProvider", time.RFC3339)
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	LogMe("TestProvider_impl", time.RFC3339)
	var _ *schema.Provider = Provider()
}
*/

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
	LogMe("testAccPreCheck", time.RFC3339)
	if value := os.Getenv("OBM_BA_USER"); value == "" {
		t.Fatal("OBM_BA_USER must be set for acceptance tests")
	}
	if value := os.Getenv("OBM_BA_PASSWORD"); value == "" {
		t.Fatal("OBM_BA_PASSWORD must be set for acceptance tests")
	}
}
