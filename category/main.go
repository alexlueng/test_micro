package main

import (
	"fmt"
	"micro/category/common"
	"micro/category/domain/repository"
	service2 "micro/category/domain/service"
	"micro/category/handler"
	category "micro/category/proto"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	// "github.com/micro/micro/v3/service/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	consulConfig, err := common.GetConsulConfig("192.168.56.101", "/micro/config", 8500)
	if err != nil {
		log.Error(err)
	}

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.56.101:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
	)

	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	fmt.Println("get mysql config from configuration center: ", mysqlInfo)

	srv.Init()

	db, err := gorm.Open("mysql", "root:Turkey414!@#@tcp(192.168.56.101:3306)/micro_user?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	rp := repository.NewCategoryRepository(db)
	rp.InitTable()

	categoryDataService := service2.NewUserDataService(rp)

	// Register handler

	// pb.RegisterUserHandler(srv.Server(), new(handler.User))
	err = category.RegisterCategoryHandler(srv.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
