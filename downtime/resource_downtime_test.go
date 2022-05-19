package downtime

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/panderosa/obmprovider/obmsdk"
	"github.com/panderosa/obmprovider/utils"
)

func TestAccDowntime_basic(t *testing.T) {
	resourceName := "downtime.new"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDowntimeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCheckDowntimeConfig(resourceName),
				Check: resource.ComposeTestCheckFunc(
					testCheckDowntimeExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "approver", "Dariusz Malinowski"),
				),
			},
		},
	})
}

func testAccCheckDowntimeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*obmsdk.Client)
	utils.LogMe("testAccCheckDowntimeDestroy", "", "")
	for _, rs := range s.RootModule().Resources {
		utils.LogMe(rs.Primary.ID, rs.Type, "")
		if rs.Type != "downtime" {
			continue
		}

		downtimeId := rs.Primary.ID
		utils.LogMe(rs.Primary.ID, "deletion", "start delete")
		err := c.Downtimes.Delete(downtimeId)
		if err != nil {
			utils.LogMe(rs.Primary.ID, "deletion error", err)
			return err
		}
	}
	return nil
}

func testCheckDowntimeConfig(name string) string {

	arrays := strings.Split(name, ".")

	return fmt.Sprintf(`
	terraform {
		required_providers {
		  downtimes = {
			source = "hashicorp/downtimes"
		  }
		}
	  }
	  
	  provider "downtimes" {
		# Configuration options
		address  = "https://a-obm.ecb.de/topaz/bsmservices/customers/1"
		path     = "/downtimes"
		username = "obm_downtime"
		password = "I8OrWXQT4mRq6Ek6jQvyA"
	  }
	
	resource "%s" "%s" {
			provider     = downtimes
			name         = "OBM Downtime test"
			action       = "STOP_MONITORING"
			approver     = "Dariusz Malinowski"
			description  = "Cerated and Updated By Acceptance Tests"
			category     = "APP_MAINTENANCE"
			selected_cis = ["b390f03f01f31a28b354503faa970c60"]
			schedule {
			type       = "ONCE"
			start_date = "2022-03-17T13:00:00+01:00"
			end_date   = "2022-03-17T23:40:00+01:00"
			timezone   = "Europe/Berlin"
			}
		}
	`, arrays[0], arrays[1])
}

func testCheckDowntimeExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("not found %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no downtime ID set")
		}

		return nil
	}
}
