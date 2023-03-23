package db

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"os"
	"strings"
)

func ExecuteSQLFile(file files.File) error {
	err := file.Open(os.O_RDWR)
	if err != nil {
		logger.Error(err)
	}
	content := file.ReadString()
	for _, cmd := range strings.Split(content, ";") {
		_, err = Instance.Exec(context.Background(), cmd)
		if err != nil {
			logger.Error(err)
			return err
		}
	}
	return nil
}
