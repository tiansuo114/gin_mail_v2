package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"tiansuo_gin_mail_v2.1/conf"
	"tiansuo_gin_mail_v2.1/middleware/logger"
	"tiansuo_gin_mail_v2.1/model/form"
	"time"
)

type db_service struct {
	Mongo_Client *mongo.Database
}

var Db_service_instance *db_service

var timeOut = 10

func connection_to_db() *mongo.Database {
	db_url := "mongodb://" + conf.Config.Database.DbHost + ":" + conf.Config.Database.DbPort + "/?connect=direct"

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut))
	defer cancel()

	clientOpts := options.Client().ApplyURI(db_url)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		logger.WriteLog("[Emergency] emergency level log" + err.Error())
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		logger.WriteLog("[Emergency] emergency level log" + err.Error())
	}

	return client.Database(conf.Config.Database.DbName)
}

func init() {
	new_db_service := db_service{
		Mongo_Client: connection_to_db(),
	}
	if new_db_service.Mongo_Client == nil {
		logger.WriteLog("[Emergency] emergency level log " + "db_service_instance is nil")
	}

	str, err := new_db_service.Mongo_Client.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		logger.WriteLog("[Emergency] emergency level log " + err.Error())
	}
	if StringSliceEqualBCE(str, formList) {
		migration(&new_db_service)
	}
	Db_service_instance = &new_db_service
}

func InsertOne(data interface{}) (result *mongo.InsertOneResult, err error) {
	err = nil
	form_name, err := getFormType(data)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}

	insertOneResult, err := Db_service_instance.Mongo_Client.Collection(form_name).InsertOne(context.TODO(), data)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}

	return insertOneResult, err
}

func FindOne(data interface{}) (result interface{}, err error) {
	err = error(nil)
	form_name, err := getFormType(data)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}

	sig := Db_service_instance.Mongo_Client.Collection(form_name).FindOne(context.TODO(), data)
	err = sig.Decode(&result)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}

	return result, nil
}

func UpdateOne(old, new interface{}) (updateOneResult *mongo.UpdateResult, err error) {
	err = nil
	form, err := getFormType(old)
	form2, err := getFormType(new)
	if strings.EqualFold(form, form2) == false {
		err = errors.New("Update type is not equal")
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	} else if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}
	form_name, err := getFormType(new)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}
	res, err := FindOne(old)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}
	var id primitive.ObjectID
	id = res.(bson.D)[0].Value.(primitive.ObjectID)
	filter := bson.D{{"_id", id}}

	Db_service_instance.Mongo_Client.Collection(form_name).FindOneAndUpdate(context.Background(), filter, bson.D{{"$set", new}})

	return updateOneResult, err
}

func DeleteOne(data interface{}) (deleteOneResult *mongo.DeleteResult, err error) {
	err = nil
	form_name, err := getFormType(data)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}

	deleteOneResult, err = Db_service_instance.Mongo_Client.Collection(form_name).DeleteOne(context.TODO(), data)
	if err != nil {
		logger.WriteLog("[Error] error level log)" + err.Error())
		return nil, err
	}

	return deleteOneResult, err
}

func StringSliceEqualBCE(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	mapA := make(map[string]bool)
	for _, v := range a {
		mapA[v] = true
	}
	for _, v := range b {
		if _, ok := mapA[v]; ok {
			delete(mapA, v)
		}
	}

	return len(mapA) == 0
}

func getFormType(data interface{}) (formType string, err error) {
	err = error(nil)
	switch data.(type) {
	case *form.Address:
		formType = "address"
	case *form.Admin:
		formType = "admin"
	case *form.Cart:
		formType = "cart"
	case *form.Category:
		formType = "category"
	case *form.Favorite:
		formType = "favorite"
	case *form.Notice:
		formType = "notice"
	case *form.Order:
		formType = "order"
	case *form.Product:
		formType = "product"
	case *form.User:
		formType = "user"
	default:
		// 处理无法匹配的情况
		logger.WriteLog("[Error] error level log)" + err.Error())
		formType = ""
	}
	return
}
