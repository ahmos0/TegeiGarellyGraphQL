package graphql

import (
	"github.com/ahmos0/DyanamodbConnectMobile/pkg/database"
	"github.com/ahmos0/DyanamodbConnectMobile/pkg/models"
)

func getAllItems() ([]models.Item, error) {
	return database.GetAllItems()
}

func saveItem(uuid string, works string, author string, url string, other string) (*models.Item, error) {
	return database.SaveItem(uuid, works, author, url, other)
}
