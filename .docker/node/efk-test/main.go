package main

import (
	"efk-test/logger"
)

func main() {
	logger.InitZap()

	logger.LogDebug("Hello World")
}
