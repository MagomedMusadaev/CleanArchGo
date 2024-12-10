package logg

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	infoLog    *log.Logger // Логгер для информационных сообщений
	warningLog *log.Logger // Логгер для предупреждающих сообщений
	errorLog   *log.Logger // Логгер для сообщений об ошибках
)

func Logging() {
	logFile, err := os.OpenFile("Loggers.logg", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Не удалось открыть файл для логов: ", err)
		return
	}

	multiWrite := io.MultiWriter(os.Stdout, logFile)

	infoLog = log.New(multiWrite, "INFO: ", log.Ldate|log.Ltime)
	warningLog = log.New(multiWrite, "WARNING: ", log.Ldate|log.Ltime)
	errorLog = log.New(multiWrite, "ERROR: ", log.Ldate|log.Ltime)
}

func Info(msg string) {
	infoLog.Println(msg)
}

func Warning(msg string) {
	warningLog.Println(msg)
}

func Error(msg string) {
	errorLog.Println(msg)
}
