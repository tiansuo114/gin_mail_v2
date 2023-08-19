package dao

import "context"

var formList []string = []string{
	"address",
	"admin",
	"cart",
	"category",
	"favorite",
	"notice",
	"order",
	"product",
	"user",
}

func migration(dbService *db_service) {
	for _, name := range formList {
		dbService.Mongo_Client.CreateCollection(context.Background(), name)
	}
}
