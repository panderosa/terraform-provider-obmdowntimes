-- suffix must be *.tfrc
[System.Environment]::SetEnvironmentVariable("TF_CLI_CONFIG_FILE","C:\\Users\\sa_malinod\\AppData\\terraform.tfrc",[System.EnvironmentVariableTarget]::Process)

go build -a -o terraform-provider-downtimes.exe
go build -a -o C:\Users\sa_malinod\AppData\terraform\plugins\hashicorp.com\mf\downtimes\0.2\amd64\terraform-provider-downtimes.exe
move-item .\terraform-provider-downtimes.exe -Destination C:\Users\sa_malinod\AppData\terraform\plugins\hashicorp.com\mf\downtimes\0.2\amd64\ -Force


Field name may only contain lowercase alphanumeric characters & underscores.
forceNew (if update not exists)
resource downtime: name: Default cannot be set with Required

-- terraform.trfc
provider_installation {
    dev_overrides {
        "hashicorp/downtimes" = "C:\\Users\\sa_malinod\\AppData\\terraform\\plugins\\hashicorp.com\\mf\\downtimes\\0.2\\amd64"
    }

}


---

OBM_