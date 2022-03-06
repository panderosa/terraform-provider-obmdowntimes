# downtime_item_list (Data Source)

Get the list of the Downtime objects based on the Downtime Filters.

## Example Usage

```terrform
data "downtime_item_list" "list_my" {
  provider     = downtimes.dev
  filter {
    name = "name"
    value = "OBM Downtime"
  }

  filter {
    name = "active"
    value = "false"
  }
}

output "my_downtimes" {
  value = data.downtime_item_list.list_my
}
```


## Argument Reference

### Required Arguments

- `filter` (Block) At least one filter block must be provided with a pair of "name" and "value" attributes.
- `name` (String) The name of the Downtime field. Possible fields are: `name`, `ciId`, `ciLabel`, `ciGlobalId`, `expired`, `active`.
- `value` (String) The value of the field to search for. For details check out Micro Focus [documentation](https://docs.microfocus.com/doc/Operations_Bridge_Manager/2020.10/DowntimeFilters) 





