package downtime

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

func dataSourceDowntimeList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDowntimeReadList,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"filter": {
				Type:     schema.TypeSet,
				Required: true,
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
			"item": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
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

	queryMap := make(map[string]string)
	for i := range filter_list {
		item := filter_list[i].(map[string]interface{})
		name := item["name"].(string)
		value := item["value"].(string)
		queryMap[name] = value
	}

	dnts, err := conn.Downtimes.Search(queryMap)
	if err != nil {
		return diag.FromErr(err)
	}
	// will be used to create id for data source
	var ids []string
	downtimeList := dnts.Downtimes
	dsis := make([]interface{}, 0, len(downtimeList))

	for _, dnt := range downtimeList {
		dsi := make(map[string]interface{})
		ids = append(ids, dnt.ID)
		dsi["id"] = dnt.ID
		dsi["name"] = dnt.Name
		dsi["action"] = dnt.Action.Type
		dsi["description"] = dnt.Description
		dsi["approver"] = dnt.Approver
		dsi["category"] = reMapCategory(dnt.Category)
		dsi["selected_cis"] = Flatten4CIs(dnt.SelectedCIs)
		schedule := make(map[string]interface{})
		schedule["type"] = dnt.Schedule.Type
		schedule["start_date"] = dnt.Schedule.StartDate
		schedule["end_date"] = dnt.Schedule.EndDate
		schedule["timezone"] = dnt.Schedule.TimeZone
		dsi["schedule"] = []interface{}{schedule}
		dsis = append(dsis, dsi)
	}

	if err := d.Set("item", dsis); err != nil {
		return diag.FromErr(err)
	}
	id := GenerateIdByHash(ids)
	d.SetId(id)

	return diags
}
