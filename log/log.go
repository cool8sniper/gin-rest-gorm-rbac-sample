package log

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("ai-backend")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} : %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLog() {
	logPath := fmt.Sprintf("/tmp/%s", "gin-rest-gorm-rbac-sample.log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Create(logPath)
	}

	logFile, err := os.OpenFile(logPath, os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Fail to get log file. Exit.")
		os.Exit(5)
	}

	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)
}
