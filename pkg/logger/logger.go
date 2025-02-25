package logger

import (
	"log"
	"os"
	"strings"
)

func InitLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, "["+strings.ToUpper(prefix)+"] ", log.LstdFlags)
}
