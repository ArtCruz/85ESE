package data

import (
	"log"
	"os"
	"sync"
)

var (
	logger     *log.Logger
	loggerOnce sync.Once
)

// GetLogger retorna a instância singleton do logger
func GetLogger() *log.Logger {
	loggerOnce.Do(func() {
		logger = log.New(os.Stdout, "products-api ", log.LstdFlags)
	})
	return logger
}
