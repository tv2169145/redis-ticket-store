package utils

import (
	"log"
	"os"
	"strings"
)

func WriteLog(msg string, logPath string) {
	fd, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}
