
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
      version = "1.0.3"
    }
  }
}

# Configure Provider options
provider "downtimes" {
  address = "https://<server>::<port>/topaz/bsmservices/customers/[customerID]"
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
  selected_cis = ["<ucmdb_id1>","<ucmdb_id2>","<ucmdb_id3>"]
  schedule {
    type       = "ONCE"
    start_date = "<start_date>"
    end_date   = "<end_date>"
    timezone   = "<timezone>"
  }
}
````

## Authentication

Provider uses Basic Authentication to authenticate in the Operations Bridge Manager Downtimes service.
Configure `username` and `password` in the Provider block or set the values in the environment variables `OBM_BA_USER` and `OBM_BA_PASSWORD`.

## Argument Reference

### Required Arguments

- `address` (String) The URL of the Operations Bridge Manager Downtimes service

### Optional Arguments

- `username` (String) Provider will use this username for Basic authentication to the Downtimes Service. By default provider takes the username from the `OBM_BA_USER` environment variable.
- password' (String, Sensitive) Provider will use this password for Basic auth to the Downtimes Service. By default provider takes the password from the `OBM_BA_PASSWORD` environment variable.
- `path` (String, Optional) Possible value is `/downtimes`.

