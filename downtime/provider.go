package downtime

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

// this function returns a terraform .ResourceProvider interface
func Provider() *schema.Provider {
	return &schema.Provider{
		// setting up shared configuration objects, e.g. addresses, secrets, access keys
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Description: "The URL of the Action web service https://<server>:<port>/topaz/bsmservices/customers/[customerID]/downtimes, where <server> is the name of the OMi server.",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OBM_BA_ADDRESS", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Provider will use this username for BASIC auth to the Downtimes API. By default provider takes the username from the OBM_BA_USER environment variable.",
				DefaultFunc: schema.EnvDefaultFunc("OBM_BA_USER", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Provider will use this password for BASIC auth to the Downtimes API. By default provider takes the password from the OBM_BA_PASSWORD environment variable.",
				DefaultFunc: schema.EnvDefaultFunc("OBM_BA_PASSWORD", nil),
			},
		},

		// register downtime resource that provider exports to Terraform
		// downtime resource must implement the schema.Resource interface
		ResourcesMap: map[string]*schema.Resource{
			"downtime": resourceDowntime(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"downtime_item":      dataSourceDowntime(),
			"downtime_item_list": dataSourceDowntimeList(),
		},

		// initialize shared configuration objects - the SDK client which makes API requests to OBM Downtime Service
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	address := d.Get("address").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	LogMe("provider configuration", fmt.Sprintf("address:%s, username:%s, password:%s", address, username, password))
	var diags diag.Diagnostics
	conn, err := obmsdk.NewClient(&address, nil, &username, &password)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return conn, diags
}
