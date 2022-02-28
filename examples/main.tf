

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
  category     = "APP_MAINTENANCE"
  selected_cis = ["4830ca40d21593b7bf85c3d070b8b8c2","45a87733d08497b5963608756d47eae1","471b3e8f086f6aeb90dae487c512aacc"]
  schedule {
    type       = "ONCE"
    start_date = "2022-02-27T14:40:00+01:00"
    end_date   = "2022-02-28T14:40:00+01:00"
    timezone   = "Europe/Berlin"
  }

}


