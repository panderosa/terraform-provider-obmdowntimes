# downtime_item (Data Source)

Get the details of the Downtime object based on the Downtime ID.

## Example Usage

```terrform
data "downtime_item" "one" {
  provider     = downtimes.dev
  id = "5943b57615034e948866c95595a97022"
}

output "single_downtime" {
  value = data.downtime_item.one
}
```


## Argument Reference

### Required Arguments

- `id` (String) The Downtime ID which is which is the same as the id of the Downtime object in the Terraform state.






