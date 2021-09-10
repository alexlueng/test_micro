package main

import (
	"fmt"
	"micro/user/domain/repository"
	service2 "micro/user/domain/service"
	"micro/user/handler"
	user "micro/user/proto"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	// "github.com/micro/micro/v3/service/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	srv.Init()

	db, err := gorm.Open("mysql", "root:Turkey414!@#@tcp(192.168.56.101:3306)/micro_user?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	rp := repository.NewUserRepository(db)
	rp.InitTable()

	userDataService := service2.NewUserDataService(rp)

	// Register handler

	// pb.RegisterUserHandler(srv.Server(), new(handler.User))
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
