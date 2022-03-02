# downtime (Resource)




## Argument Reference

### Required Arguments

- `name` (String) Name of the downtime.
- `action` (String) A value which represents the downtime action name. Possible values are: `REMINDER`, `SUPPRESS_NOTIFICATIONS`, `ENFORCE_ON_KPI_CALCULATION`, `ENFORCE_ON_REPORTS`, `STOP_MONITORING`, `ENFORCE_ON_REPORTS`.
- `approver` (String) Name of the approver of the downtime
- `category` (String) A value which represents the downtime category name. Possible values are: `NW_MAINTENANCE`, `HW_MAINTENANCE`, `HW_INSTALLATION`, `SECURITY`, `OTHER`, `OS_CONFIGURATION`, `APP_MAINTENANCE`, `APP_INSTALLATION`.
- `schedule` (Block Set, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--schedule))
- `selected_cis` (Set of String) List configuration item ucmdbIds which are impacted by the downtime. Take it from RTSM.

### Optional Arguments

- `description` (String) Long description for the downtime.
  

<a id="nestedblock--schedule"></a>
### Nested Schema for `schedule`

##### Required:

- `type` (String) For the current version of the provider must be set to ONCE. The donwtime occurs only once on a specified start date and lasts till a specified end date.
- `end_date` (String) The downtime end date specified as the RFC 3339 date-time format: `yyyy-MM-dd'T'HH:mm:ss('+'/'-')HH:mm`.
- `start_date` (String) The downtime start date specified as the RFC 3339 date-time format: `yyyy-MM-dd'T'HH:mm:ss('+'/'-')HH:mm`.
- `timezone` (String) Timezone.



