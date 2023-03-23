package api

import (
	"context"
	"github.com/Galagoshin/GoLogger/logger"
	"orgkombot/db"
)

type Inventory struct {
	Owner *User
	items []*Item
}

func (user *User) GetInventory() *Inventory {
	items := []*Item{}
	result := &Inventory{Owner: user}
	rows, err := db.Instance.Query(context.Background(), "SELECT item_id FROM items WHERE owner_id = $1", user.GetId())
	if err != nil {
		logger.Error(err)
		return &Inventory{}
	}
	for rows.Next() {
		var item_id uint8
		err := rows.Scan(&item_id)
		if err != nil {
			logger.Error(err)
			return nil
		}
		items = append(items, &Item{Inventory: result, Id: item_id})
	}
	result.items = items
	return result
}

func (inventory *Inventory) AddItem(item Item) {
	err := db.Instance.QueryRow(context.Background(), "INSERT INTO items (owner_id, item_id) VALUES ($1, $2);", inventory.Owner.GetId(), item.Id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	inventory.items = append(inventory.items, &item)
}

func (inventory *Inventory) RemoveItem(item Item) {
	err := db.Instance.QueryRow(context.Background(), "DELETE FROM items WHERE owner_id = $1 AND item_id = $2;", inventory.Owner.GetId(), item.Id).Scan()
	if err != nil {
		if err.Error() != "no rows in result set" {
			logger.Error(err)
			return
		}
	}
	new_items := []*Item{}
	deleted := false
	for _, old_item := range inventory.items {
		if !deleted && old_item.Id == item.Id {
			deleted = true
		} else {
			new_items = append(new_items, old_item)
		}
	}
	inventory.items = new_items
}

func (inventoty *Inventory) GetItems() []*Item {
	return inventoty.items
}
