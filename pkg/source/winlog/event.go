package winlog

import (
	"fmt"
	"log"
	"time"

	"github.com/elastic/beats/v7/winlogbeat/checkpoint"
	"github.com/elastic/beats/v7/winlogbeat/eventlog"

	"github.com/elastic/beats/v7/libbeat/common"
)

func main() {
	// 配置事件日志读取参数
	config := common.MustNewConfigFrom(map[string]interface{}{
		"name":            "Application", // 读取的事件日志名称，例如 "Application", "System" 等
		"batch_read_size": 10,            // 每次读取的事件数量
	})

	// 创建事件日志对象
	eventLog, err := eventlog.New(config)
	if err != nil {
		log.Fatalf("Failed to create event log: %v", err)
	}

	// 打开事件日志
	err = eventLog.Open(checkpoint.EventLogState{})
	if err != nil {
		log.Fatalf("Failed to open event log: %v", err)
	}
	// 关闭事件日志
	defer func(eventLog eventlog.EventLog) {
		err := eventLog.Close()
		if err != nil {
		}
	}(eventLog)

	for {
		// 读取事件日志
		records, err := eventLog.Read()
		if err != nil {
			log.Fatalf("Failed to read event log: %v", err)
		}

		// 打印事件日志
		for _, record := range records {
			fmt.Printf("Event ID: %d\n", record.EventIdentifier.ID)
			fmt.Printf("Provider Name: %s\n", record.Provider.Name)
			fmt.Printf("Event Message: %s\n", record.Message)
			fmt.Printf("Time Created: %s\n", record.TimeCreated.SystemTime.Format(time.RFC3339))
			fmt.Println("===================================")
		}

		// 暂停一段时间以避免频繁读取
		time.Sleep(1 * time.Second)
	}
}
