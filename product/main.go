package main

import (
	"fmt"
	"micro/product/common"
	"micro/product/domain/repository"
	service2 "micro/product/domain/service"
	"micro/product/handler"
	product "micro/product/proto"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	// "github.com/micro/micro/v3/service/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	consulConfig, err := common.GetConsulConfig("192.168.56.101", "/micro/config", 8500)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(consulConfig)

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.56.101:8500",
		}
	})

	t, io, err := common.NewTracer("go.micro.service.product", "192.168.56.101:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()

	db, err := gorm.Open("mysql", "root:Turkey414!@#@tcp(192.168.56.101:3306)/micro_user?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	rp := repository.NewProductRepository(db)
	// rp.InitTable()

	productDataService := service2.NewProductDataService(rp)

	// Register handler

	// pb.RegisterUserHandler(srv.Server(), new(handler.User))
	err = product.RegisterProductHandler(srv.Server(), &handler.Product{ProductDataService: productDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
