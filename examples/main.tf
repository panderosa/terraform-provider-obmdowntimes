

terraform {
  required_providers {
    downtimes = {
      source = "hashicorp/downtimes"
    }
  }
}

provider "downtimes" {
  # Configuration options
  address = "https://epmwd211.t-mgmt.tadnet.net/topaz/bsmservices/customers/1"
  path    = "/downtimes"
  alias   = "dev"
}


resource "downtime" "downtime_01" {
  provider     = downtimes.dev
  name         = "OBM Dowtime"
  description  = "Created from terraform. change 2."
  action       = "STOP_MONITORING"
  approver     = "Dariusz Malinowski"
  category     = "OS_CONFIGURATION"
  selected_cis = "4830ca40d21593b7bf85c3d070b8b8c2"
  start_date    = "2022-02-25T14:40:00+01:00"
  end_date      = "2022-02-28T14:40:00+01:00"
}

resource "downtime" "downtime_02" {
  provider     = downtimes.dev
  name         = "OBM Dowtime second"
  description  = "Created from terraform. change 1."
  action       = "STOP_MONITORING"
  approver     = "Dariusz Malinowski"
  category     = "SECURITY"
  selected_cis = "4830ca40d21593b7bf85c3d070b8b8c2"
  start_date    = "2022-03-01T14:40:00+01:00"
  end_date      = "2022-03-03T14:40:00+01:00"
}

