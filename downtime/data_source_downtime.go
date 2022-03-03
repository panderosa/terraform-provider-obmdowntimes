package downtime

import (
	"context"

	//"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

func dataSourceDowntime() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDowntimeRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"approver": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"selected_cis": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"schedule": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timezone": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDowntimeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)

	var diags diag.Diagnostics

	downtimeID := d.Get("id").(string)
	dnt, err := conn.Downtimes.Read(downtimeID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", dnt.Name)
	d.Set("action", dnt.Action.Type)
	d.Set("description", dnt.Description)
	d.Set("approver", dnt.Approver)
	d.Set("category", reMapCategory(dnt.Category))
	if err := d.Set("selected_cis", flatten3CIs(dnt.SelectedCIs)); err != nil {
		return diag.FromErr(err)
	}

	item := make(map[string]interface{})
	item["type"] = dnt.Schedule.Type
	item["start_date"] = dnt.Schedule.StartDate
	item["end_date"] = dnt.Schedule.EndDate
	item["timezone"] = dnt.Schedule.TimeZone

	if err := d.Set("schedule", []interface{}{item}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(downtimeID)
	return diags
}

func flatten3CIs(data []obmsdk.Ci) []interface{} {
	array := make([]interface{}, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return array
}
