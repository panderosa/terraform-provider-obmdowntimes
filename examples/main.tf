terraform {
  required_providers {
    // Used for development together with terrafrom_dev.tfrc
    /*downtimes = {
      source = "hashicorp/downtimes"
    }*/
    obmdowntimes = {
      source  = "panderosa/obmdowntimes"
      version = "1.0.7"
    } 
    ecbdowntimes = {
      source  = "ismxa421.ecb01.ecb.de/ecb/downtimes"
      version = "1.0.7"
    }
  }
}

resource "downtime" "downtime_01" {
  provider    = ecbdowntimes
  action      = "STOP_MONITORING"
  approver    = "Dariusz Malinowski"
  category    = "APP_MAINTENANCE"
  description = "Created and Updated By Terraform obmdowntimes provider compiled locally"
  name        = "OBM Downtime - compiled locally"
  selected_cis = [
    "b390f03f01f31a28b354503faa970c60",
  ]

  schedule {
    end_date   = "2022-03-20T21:30:00+01:00"
    start_date = "2022-03-21T22:00:00+01:00"
    timezone   = "Europe/Berlin"
    type       = "ONCE"
  }
}


/*
resource "downtime" "downtime_02" {
  provider    = obmdowntimes
  action      = "STOP_MONITORING"
  approver    = "Dariusz Malinowski"
  category    = "APP_MAINTENANCE"
  description = "Created and Updated By Terraform obmdowntimes provider installed from terraform registry"
  name        = "OBM Downtime - from regirstry"
  selected_cis = [
    "b390f03f01f31a28b354503faa970c60",
  ]

  schedule {
    end_date   = "2022-03-20T21:30:00+01:00"
    start_date = "2022-03-21T22:00:00+01:00"
    timezone   = "Europe/Berlin"
    type       = "ONCE"
  }
}
*/