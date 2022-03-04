package downtime

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogMe(statement string, message string) {
	filename, cond1 := os.LookupEnv("OBM_PROVIDER_LOG_FILE")
	if cond1 {
		currentTime := time.Now()
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		data := fmt.Sprintf("--- %v --- %v ---\n%v\n", currentTime.Format("2006-01-02 15:04:05"), statement, message)

		_, err1 := file.WriteString(data)
		if err1 != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}
