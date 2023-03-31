package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"orgkombot/db"
)

func CreateTransaction(from *User, to *User, amount uint) {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO transactions (admin_id, user_id, amount) VALUES ($1, $2, $3);", from.GetId(), to.GetId(), amount).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
}
