package downtime

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

func resourceDowntime() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDowntimeCreate,
		ReadContext:   resourceDowntimeRead,
		UpdateContext: resourceDowntimeUpdate,
		DeleteContext: resourceDowntimeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"approver": {
				Type:     schema.TypeString,
				Required: true,
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"selected_cis": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							InputDefault: "ONCE",
						},
						"start_date": {
							Type:     schema.TypeString,
							Required: true,
						},
						"end_date": {
							Type:     schema.TypeString,
							Required: true,
						},
						"timezone": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			/*"schedule": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "ONCE",
						},
						"start_date": {
							Type:     schema.TypeString,
							Required: true,
						},
						"end_date": {
							Type:     schema.TypeString,
							Required: true,
						},
						"timezone": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},*/
			/*"start_date": {
				Type:     schema.TypeString,
				Required: true,
			},
			"end_date": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ONCE",
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Europe/Berlin",
			},*/
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

var (
	categories = map[string]string{
		"OTHER":            "1",
		"OS_CONFIGURATION": "2",
		"APP_MAINTENANCE":  "3",
		"APP_INSTALLATION": "4",
		"NW_MAINTENANCE":   "5",
		"HW_MAINTENANCE":   "6",
		"HW_INSTALLATION":  "7",
		"SECURITY":         "8",
	}
)

func mapCategory(name string) string {
	for k, v := range categories {
		if k == name {
			return v
		}
	}
	return name
}

func reMapCategory(cid string) string {
	for k, v := range categories {
		if v == cid {
			return k
		}
	}
	return cid
}

func mapCIs(data string) []obmsdk.Ci {
	array := strings.Split(data, ",")
	cis := make([]obmsdk.Ci, 0, len(array))
	for i := range array {
		cis = append(cis, obmsdk.Ci{
			ID: array[i],
		})
	}
	return cis
}

func flattenCIs(data []obmsdk.Ci) string {
	array := make([]string, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return strings.Join(array, ",")
}

func resourceDowntimeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)
	options := obmsdk.DowntimeCreateOptions{
		UserId:      "1",
		Planned:     "true",
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Approver:    d.Get("approver").(string),
		Category:    mapCategory(d.Get("category").(string)),
		SelectedCIs: mapCIs(d.Get("selected_cis").(string)),
	}
	schedule := d.Get("schedule").(*schema.Set)
	schedule_list := schedule.List()
	schedule_map := schedule_list[0].(map[string]interface{})
	options.Schedule.Type = schedule_map["type"].(string)
	options.Schedule.TimeZone = schedule_map["timezone"].(string)
	options.Schedule.StartDate = schedule_map["start_date"].(string)
	options.Schedule.EndDate = schedule_map["end_date"].(string)
	options.Action.Type = d.Get("action").(string)

	downtime, err := conn.Downtimes.Create(options)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(downtime.ID)
	return resourceDowntimeRead(ctx, d, meta)
}

func anyUpdate(d *schema.ResourceData) bool {
	status := false
	keys := []string{"name", "description", "approver", "category", "schedule", "action", "selected_cis"}
	for i := range keys {
		status = status || d.HasChange(keys[i])
	}
	return status
}

func resourceDowntimeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)
	downtimeID := d.Id()
	if anyUpdate(d) {
		options := obmsdk.Downtime{
			Name:        d.Get("name").(string),
			ID:          downtimeID,
			Description: d.Get("description").(string),
			Approver:    d.Get("approver").(string),
			Category:    mapCategory(d.Get("category").(string)),
			SelectedCIs: mapCIs(d.Get("selected_cis").(string)),
		}
		schedule := d.Get("schedule").(*schema.Set)
		schedule_list := schedule.List()
		schedule_map := schedule_list[0].(map[string]interface{})
		options.Schedule.Type = schedule_map["type"].(string)
		options.Schedule.TimeZone = schedule_map["timezone"].(string)
		options.Schedule.StartDate = schedule_map["start_date"].(string)
		options.Schedule.EndDate = schedule_map["end_date"].(string)
		options.Action.Type = d.Get("action").(string)

		err := conn.Downtimes.Update(downtimeID, options)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}
	return resourceDowntimeRead(ctx, d, meta)
}

func resourceDowntimeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)

	var diags diag.Diagnostics

	downtimeID := d.Id()
	dnt, err := conn.Downtimes.Read(downtimeID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", dnt.Name)
	d.Set("action", dnt.Action.Type)
	d.Set("description", dnt.Description)
	d.Set("approver", dnt.Approver)
	d.Set("category", reMapCategory(dnt.Category))
	d.Set("selected_cis", flattenCIs(dnt.SelectedCIs))

	/*schedule_item := &map[string]string{
		"type":       dnt.Schedule.Type,
		"start_date": dnt.Schedule.StartDate,
		"end_date":   dnt.Schedule.EndDate,
		"timezone":   dnt.Schedule.TimeZone,
	}*/

	//schedule := &schema.Set{}
	//schedule.Add(schedule_item)
	item := make(map[string]interface{})
	item["type"] = dnt.Schedule.Type
	item["start_date"] = dnt.Schedule.StartDate
	item["end_date"] = dnt.Schedule.EndDate
	item["timezone"] = dnt.Schedule.TimeZone

	//var set *schema.Set = new(schema.Set)
	//set.Add(item)

	if err := d.Set("schedule", []interface{}{item}); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceDowntimeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)

	var diags diag.Diagnostics

	downtimeID := d.Id()
	err := conn.Downtimes.Delete(downtimeID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
