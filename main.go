package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/panderosa/obmprovider/downtime"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run provider with support for debugger")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return downtime.Provider()
		},
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "hashicorp/downtimes", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	plugin.Serve(opts)
}
