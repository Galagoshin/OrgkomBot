package db

import (
	"context"
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/files"
	"github.com/Galagoshin/VKGoBot/bot/framework"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

func Init() {
	host := os.Getenv("POSTGRES_HOST")
	if framework.IsUnderDocker() {
		host = "host.docker.internal:5432"
	}
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		host,
		os.Getenv("POSTGRES_DB")))
	if err != nil {
		logger.Panic(err)
	}
	Instance = conn
	attempts := 0
	for true {
		err = ExecuteSQLFile(files.File{Path: "sql/init_tables.sql"})
		if attempts == 5 && err != nil {
			logger.Panic(err)
		} else if err != nil {
			attempts++
			logger.Warning("Retrying connect to database in 20 seconds...")
			time.Sleep(time.Second * 20)
		} else {
			break
		}
	}
	logger.Print("Database has been loaded.")
}
