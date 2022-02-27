--- debugging
$env:TF_REATTACH_PROVIDERS='{"hashicorp/downtimes":{"Protocol":"grpc","ProtocolVersion":5,"Pid":8388,"Test":true,"Addr":{"Network":"tcp","String":"127.0.0.1:50367"}}}'

--- Development mode
-- suffix must be *.tfrc
[System.Environment]::SetEnvironmentVariable("TF_CLI_CONFIG_FILE","C:\\Users\\sa_malinod\\AppData\\terraform.tfrc",[System.EnvironmentVariableTarget]::Process)

-- terraform.trfc
provider_installation {
    dev_overrides {
        "hashicorp/downtimes" = "C:\\Users\\sa_malinod\\AppData\\terraform\\plugins\\hashicorp.com\\mf\\downtimes\\0.2\\amd64"
    }
}

--- run terraform cli in debug mode
-- start provide in debug mode (via launch.json)
-- set environment variable - reattachment in the coman line where terraform cli is executed
$env:TF_REATTACH_PROVIDERS='{"hashicorp/downtimes":{"Protocol":"grpc","ProtocolVersion":5,"Pid":3196,"Test":true,"Addr":{"Network":"tcp","String":"127.0.0.1:50828"}}}'


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

TypeMap with Elem *Resource not supported,use TypeList/TypeSet with Elem *Resource or TypeMap with Elem *Schema