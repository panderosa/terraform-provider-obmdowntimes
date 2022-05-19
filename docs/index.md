
# OBM Downtimes Provider

The OBM Downtimes Provider can be used to create, update or delete downtimes in [Micro Focus Operations Bridge Manager](https://docs.microfocus.com/doc/Operations_Bridge_Manager/2020.10) using the REST API. Documentation regarding the Resource supported by the OBM Downtimes Provider can be found in the navigation to the left.

## Getting Started

If you're new to the Operations Bridge Manager Downtimes, check out Micro Focus [documentation](https://docs.microfocus.com/doc/Operations_Bridge_Manager/2020.10/DowntimeRESTService)


## Example Usage

```hcl
# Configure Terraform
terraform {
  required_providers {
    downtimes = {
      source = "hashicorp/downtimes"
      version = "1.0.5"
    }
  }
}

# Configure Provider options
provider "downtimes" {
  address = "https://<server>::<port>/topaz/bsmservices/customers/[customerID]/downtimes"
  path    = "/downtimes"
  username = "<username>"
  password = "<password>"
  alias   = "<alias>"
}

# Create Downtime
resource "downtime" "<resource_name>" {
  provider     = downtimes.<alias>
  name         = "<downtime name>"
  action       = "<downtime action>"
  approver     = "<user name>"
  description  = "<long_description>"
  category     = "APP_MAINTENANCE"
  selected_cis = ["<ucmdbId>",..."]
  schedule {
    type       = "ONCE"
    start_date = "<start_date>"
    end_date   = "<end_date>"
    timezone   = "<timezone>"
  }
}
````

## Argument Reference

- `address` (String) The URL address of the Operations Bridge Manager Downtimes service. Add as explicit configuration in the provider block or set as in the `OBM_BA_ADDRESS` Environmental Variable
- `username` (String) Provider will use this username for Basic authentication. Add as explicit configuration in the provider block or set as in the `OBM_BA_USER` Environment Variable.
- `password` (String, Sensitive) Provider will use this password for Basic authentication. Add as explicit configuration in the provider block or set as in the `OBM_BA_PASSWORD` Environment Variable.

