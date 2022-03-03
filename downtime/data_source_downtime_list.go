package downtime

import (
	"context"
	"fmt"
	"net/url"

	//"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

func dataSourceDowntimeList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDowntimeReadList,

		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceDowntimeReadList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)

	var diags diag.Diagnostics

	filter := d.Get("filter").(*schema.Set)
	filter_list := filter.List()

	queryString := ""
	for i := range filter_list {
		item := filter_list[i].(map[string]string)
		name := item["name"]
		value := item["value"]
		if queryString == "" {
			queryString = fmt.Sprintf("filter=%v==%v", name, url.QueryEscape(value))
		} else {
			queryString = fmt.Sprintf("%v&filter=%v==%v", queryString, name, url.QueryEscape(value))
		}
	}

	dnts, err := conn.Downtimes.Search(queryString)
	if err != nil {
		return diag.FromErr(err)
	}
	downtimeList := dnts.Downtimes
	dsis := make([]interface{}, len(downtimeList))

	for i, dnt := range downtimeList {
		dsi := make(map[string]interface{})
		dsi["id"] = dnt.ID
		dsi["name"] = dnt.Name
		dsi["action"] = dnt.Action.Type
		dsi["description"] = dnt.Description
		dsi["approver"] = dnt.Approver
		dsi["category"] = reMapCategory(dnt.Category)
		dsi["selected_cis"] = flatten3CIs(dnt.SelectedCIs)
		schedule := make(map[string]interface{})
		schedule["type"] = dnt.Schedule.Type
		schedule["start_date"] = dnt.Schedule.StartDate
		schedule["end_date"] = dnt.Schedule.EndDate
		schedule["timezone"] = dnt.Schedule.TimeZone
		dsi["schedule"] = schedule
		dsis[i] = dsi
		d.SetId(dnt.ID)
	}
	if err := d.Set("items", dsis); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func flatten4CIs(data []obmsdk.Ci) []interface{} {
	array := make([]interface{}, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return array
}
