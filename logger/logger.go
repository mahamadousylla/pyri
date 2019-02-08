package logger

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func GetLogger(ctx context.Context, fileName, funcName string) *log.Entry {
	logger := log.WithFields(log.Fields{
		"fileName": fileName,
		"funcName": funcName,
	})

	return logger
}
