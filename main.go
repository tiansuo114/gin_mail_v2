package main

import (
	"fmt"
	"tiansuo_gin_mail_v2.1/model/dao"
	"tiansuo_gin_mail_v2.1/model/form"
)

func main() {
	old_user := form.User{
		UserName: "2",
	}
	new_user := form.User{
		UserName:       "2",
		PasswordDigest: "3",
	}

	res, err := dao.UpdateOne(&old_user, &new_user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
