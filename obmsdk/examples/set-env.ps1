$obmenv = @{
    OBM_BASE_URL = "https://epmwd211.t-mgmt.tadnet.net/topaz/bsmservices/customers/1";
    OBM_DOWNTIME_PATH = "/downtimes";
    OBM_BA_USER = "obm_api_user";
    OBM_BA_PASSWORD = "Qwer1234";
    OBM_LOG_FILE = "C:\\Users\\sa_malinod\\log.txt"; # if set store HTTP request/response payload to the file
}

foreach ($key in $obmenv.Keys) {
    [System.Environment]::SetEnvironmentVariable($key,$($obmenv.$key),[System.EnvironmentVariableTarget]::Process)    
}          

