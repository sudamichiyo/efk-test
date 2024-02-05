package main

import (
	// "efk-test/logger"
	"fmt"

	"github.com/fluent/fluent-logger-golang/fluent"
)

func main() {
	// logger.InitZap()

	// logger.LogDebug("Hello World")

	logger, err := fluent.New(fluent.Config{FluentHost: "fluentd", FluentPort: 24224})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()

	tag := "service.log"
	var data = map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// ログを転送する
	err = logger.Post(tag, data)
	if err != nil {
		fmt.Println(err)
	}
}
