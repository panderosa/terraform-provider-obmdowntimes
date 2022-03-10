# create OBM Downtime

### 1. add resource block to the terraform configuration, amend `start_date` and `end_date` 

```terraform
resource "downtime" "downtime_01" {
  provider     = downtimes.dev
  name         = "OBM Downtime"
  action       = "STOP_MONITORING"
  approver     = "Dariusz Malinowski"
  description  = ""
  category     = "APP_MAINTENANCE"
  selected_cis = ["4830ca40d21593b7bf85c3d070b8b8c2","45a87733d08497b5963608756d47eae1","471b3e8f086f6aeb90dae487c512aacc"]
  schedule {
    type       = "ONCE"
    start_date = "2022-03-10T10:40:00+01:00"
    end_date   = "2022-03-12T14:40:00+01:00"
    timezone   = "Europe/Berlin"
  }
}
````

### 2. execute plan

> *terraform plan*

### 3. execute apply

> *terraform apply -auto-approve=true*

### 4. list resources in the terraform state

> *terraform state list*

### 5. show created downtime from the state

> *terraform state show `downtime.downtime_01`*



# destroy the OBM Downtime

### 1. comment or remove the resource block in from terraform configuration

### 2. execute plan

> *terraform plan*

### 3. execute apply

> *terraform apply -auto-approve=true.*

# create data source providing list of OBM Downtimes based on a given filter

## 1. add the data block to the terraform configuration  

```terraform
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
```
## 2. execute apply to add data source to the terraform state

> *terraform apply*

## 3. read data from the data source

> *terraform state show `data.downtime_item_list.list_my`*

## 4. change filter condition, apply and read data source again


# add output to the terraform module

## 1. add the output block to terraform configuration

```terraform
output "my_downtimes" {
  value = data.downtime_item_list.list_my
}
```
## 2. execute plan

> *terraform plan*

## 3. execute apply 

> *terraform apply -auto-approve=true*

## 4. show output

> *terraform output*

## 5. comment or remove the date and output block from the terraform configuration and execute apply

> *terraform apply -auto-approve=true*

# import existing downtime to manage it in the terraform state

## 1. add to terraform configuration

```terraform
resource "downtime" "sample" {
  provider = downtimes.dev
}
```
## 2. run from command line

terraform import `downtime.sample` `01d9a07112e248d68976411305e07442`

### where

- `downtime.sample` - resource name in configuration
- `01d9a07112e248d68976411305e07442` - OBM downtime ID

## 3. list the terraform state

terraform state list

## 4. show the imported resource details from the terraform state

terraform state show `downtime.sample`

## 4. replace the resource block with the output from the previous step

```terraform
resource "downtime" "sample" {
    action       = "STOP_MONITORING"
    approver     = "Dariusz Malinowski"
    category     = "APP_MAINTENANCE"
    description  = "Updated today from map"
    id           = "01d9a07112e248d68976411305e07442"
    name         = "OBM Dowtime"
    selected_cis = [
        "4830ca40d21593b7bf85c3d070b8b8c2",
    ]

    schedule {
        end_date   = "2022-02-27T14:40:00+01:00"
        start_date = "2022-02-25T14:40:00+01:00"
        timezone   = "Europe/Berlin"
        type       = "ONCE"
    }
}
```
## 5. amend the resource block by adding `provider = "downtimes.dev"` meta-argument and delete the `id` attribute as follow

```terraform
resource "downtime" "sample" {
    provider     = "downtimes.dev"
    action       = "STOP_MONITORING"
    approver     = "Dariusz Malinowski"
    category     = "APP_MAINTENANCE"
    description  = "Updated today from map"
    name         = "OBM Dowtime"
    selected_cis = [
        "4830ca40d21593b7bf85c3d070b8b8c2",
    ]

    schedule {
        end_date   = "2022-02-27T14:40:00+01:00"
        start_date = "2022-02-25T14:40:00+01:00"
        timezone   = "Europe/Berlin"
        type       = "ONCE"
    }
}
```

## 6. execute terraform plan to confirm if the infrastructure matches now the configuration

terraform plan


