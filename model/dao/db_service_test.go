package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"tiansuo_gin_mail_v2.1/model/form"
)

func TestInsert(t *testing.T) {
	//在test模块下暂时没有找到可以使其正常运行的方法,因为本函数需要config.go读取yaml配置文件
	//而因为test下的根目录与main函数执行时的根目录不同.导致配置文件无法正常读取
	user := form.User{
		UserName:       "114514",
		Email:          "114514",
		PasswordDigest: "114514",
		NickName:       "114514",
		Status:         "114514",
		Avatar:         "114514",
		Money:          "114514",
	}

	insertRes, err := InsertOne(&user)
	if err != nil {
		fmt.Println(err)
	} else {
		// 根据返回的插入结果,获取_id进行删除
		id := insertRes.InsertedID

		// 构造删除过滤条件
		filter := bson.D{{"_id", id}}

		// 删除文档
		_, err = Db_service_instance.Mongo_Client.Collection("user").DeleteOne(context.TODO(), filter)
		if err != nil {
			fmt.Println("delete error:", err)
		}
	}
}

func TestFindOne(t *testing.T) {
	user := form.User{
		UserName: "114514",
	}
	_, err := FindOne(&user)
	if err != nil {
		fmt.Println(err)
	}
}
