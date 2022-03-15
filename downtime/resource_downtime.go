package downtime

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/panderosa/obmprovider/checking"
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
				Type:         schema.TypeString,
				Description:  "Name of the downtime.",
				ValidateFunc: validation.StringIsNotEmpty,
				Required:     true,
			},
			"action": {
				Type:         schema.TypeString,
				Description:  "A value which represents the downtime action name. Valid values: REMINDER, SUPPRESS_NOTIFICATIONS, ENFORCE_ON_KPI_CALCULATION, ENFORCE_ON_REPORTS, STOP_MONITORING, ENFORCE_ON_REPORTS.",
				Required:     true,
				ValidateFunc: validation.StringInSlice(checking.DowntimeActions(), false),
			},
			"approver": {
				Type:        schema.TypeString,
				Description: "Name of the approver of the downtime",
				Required:    true,
			},
			"category": {
				Type:         schema.TypeString,
				Description:  "A value which represents the downtime category name. Valid values: NW_MAINTENANCE, HW_MAINTENANCE, HW_INSTALLATION, SECURITY, OTHER, OS_CONFIGURATION, APP_MAINTENANCE, APP_INSTALLATION.",
				Required:     true,
				ValidateFunc: checking.ValidateCategory,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Long description for the downtime.",
				Optional:    true,
			},
			"selected_cis": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "Array of RTSM configuration items id's impacted by the downtime: [\"id1\",\"id2\",...].",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"schedule": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "For the current version of the provider must be set to ONCE. The donwtime occurs only once on a specified start date and lasts till a specified end date.",
							ValidateFunc: checking.IsOnce,
						},
						"start_date": {
							Type:         schema.TypeString,
							Description:  "The downtime start date specified as the RFC 3339 date-time format: yyyy-MM-dd'T'HH:mm:ss('+'/'-')HH:mm.",
							ValidateFunc: validation.IsRFC3339Time,
							Required:     true,
						},
						"end_date": {
							Type:         schema.TypeString,
							Description:  "The downtime end date specified as the RFC 3339 date-time format: yyyy-MM-dd'T'HH:mm:ss('+'/'-')HH:mm.",
							ValidateFunc: validation.IsRFC3339Time,
							Required:     true,
						},
						"timezone": {
							Type:         schema.TypeString,
							Description:  "Timezone used in start_date and end_date.",
							Required:     true,
							ValidateFunc: checking.ValidateTimezone,
						},
					},
				},
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func mapCategory(name string) string {
	for k, v := range checking.Categories() {
		if k == name {
			return v
		}
	}
	return name
}

func reMapCategory(cid string) string {
	for k, v := range checking.Categories() {
		if v == cid {
			return k
		}
	}
	return cid
}

func resourceDowntimeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*obmsdk.Client)

	cis := d.Get("selected_cis").(*schema.Set)
	cis_list := cis.List()
	selected_cis := make([]obmsdk.Ci, len(cis_list))

	for i := range cis_list {
		selected_cis[i].ID = cis_list[i].(string)
	}

	schedule := d.Get("schedule").(*schema.Set)
	schedule_list := schedule.List()
	schedule_map := schedule_list[0].(map[string]interface{})

	options := obmsdk.DowntimeCreateOptions{
		UserId:      "1",
		Planned:     "true",
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Approver:    d.Get("approver").(string),
		Category:    mapCategory(d.Get("category").(string)),
		Action: obmsdk.Action{
			Type: d.Get("action").(string),
		},
		SelectedCIs: selected_cis, //mapCIs(d.Get("selected_cis").(string)),
		Schedule: obmsdk.Schedule{
			Type:      schedule_map["type"].(string),
			StartDate: schedule_map["start_date"].(string),
			EndDate:   schedule_map["end_date"].(string),
			TimeZone:  schedule_map["timezone"].(string),
		},
	}

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
		cis := d.Get("selected_cis").(*schema.Set)
		cis_list := cis.List()
		selected_cis := make([]obmsdk.Ci, len(cis_list))

		for i := range cis_list {
			selected_cis[i].ID = cis_list[i].(string)
		}

		schedule := d.Get("schedule").(*schema.Set)
		schedule_list := schedule.List()
		schedule_map := schedule_list[0].(map[string]interface{})

		options := obmsdk.Downtime{
			Name:        d.Get("name").(string),
			ID:          downtimeID,
			Description: d.Get("description").(string),
			Approver:    d.Get("approver").(string),
			Category:    mapCategory(d.Get("category").(string)),
			Action: obmsdk.Action{
				Type: d.Get("action").(string),
			},
			SelectedCIs: selected_cis, //mapCIs(d.Get("selected_cis").(string)),
			Schedule: obmsdk.Schedule{
				Type:      schedule_map["type"].(string),
				StartDate: schedule_map["start_date"].(string),
				EndDate:   schedule_map["end_date"].(string),
				TimeZone:  schedule_map["timezone"].(string),
			},
		}

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
	d.Set("selected_cis", Flatten3CIs(dnt.SelectedCIs))

	item := make(map[string]interface{})
	item["type"] = dnt.Schedule.Type
	item["start_date"] = dnt.Schedule.StartDate
	item["end_date"] = dnt.Schedule.EndDate
	item["timezone"] = dnt.Schedule.TimeZone

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
