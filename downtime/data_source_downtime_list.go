package downtime

import (
	"context"

	//"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

func dataSourceDowntimeList() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDowntimeReadList,

		Schema: map[string]*schema.Schema{
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

	downtimeList := dnts.Downtimes
	dsis := make([]interface{}, 0, len(downtimeList))
	//dsis := new(schema.Set)

	for _, dnt := range downtimeList {
		dsi := make(map[string]interface{})

		dsi["id"] = dnt.ID
		dsi["name"] = dnt.Name
		dsi["action"] = dnt.Action.Type
		dsi["description"] = dnt.Description
		dsi["approver"] = dnt.Approver
		dsi["category"] = reMapCategory(dnt.Category)
		dsi["selected_cis"] = flatten4CIs(dnt.SelectedCIs)
		/*schedule := make(map[string]interface{})
		schedule["type"] = dnt.Schedule.Type
		schedule["start_date"] = dnt.Schedule.StartDate
		schedule["end_date"] = dnt.Schedule.EndDate
		schedule["timezone"] = dnt.Schedule.TimeZone
		dsi["schedule"] = schedule*/
		dsis = append(dsis, dsi)
	}
	/*buf, err := json.Marshal(dsis[0])
	if err != nil {
		return diag.FromErr(err)
	}
	LogMe("dsis[0] >>>>", string(buf))

	tt := reflect.ValueOf(dsis[0]).Kind().String()
	LogMe("dsis[0] type is ...", tt)*/

	/*dsi := make(map[string]interface{})
	dsi["name"] = downtimeList[0].Name
	dsi["id"] = downtimeList[0].ID

	dsi1 := make(map[string]interface{})
	dsi1["name"] = downtimeList[1].Name
	dsi1["id"] = downtimeList[2].ID
	ios := []interface{}{dsi, dsi1}
	LogMe("ios type is ...", fmt.Sprintf("%T and value %v", ios, ios))
	dsis := make([]interface{}, 0, 2)
	dsis = append(dsis, dsi)
	dsis = append(dsis, dsi1)
	LogMe("dsis type is ...", fmt.Sprintf("%T and value %v", dsis, dsis))*/
	//if err := d.Set("item", ios); err != nil { // - this works
	if err := d.Set("item", dsis); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("123456789")
	return diags
}

func flatten4CIs(data []obmsdk.Ci) []interface{} {
	array := make([]interface{}, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return array
}
