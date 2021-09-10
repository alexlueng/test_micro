package main

import (
	"fmt"
	"micro/cart/common"
	"micro/cart/domain/repository"
	service2 "micro/cart/domain/service"
	"micro/cart/handler"
	cart "micro/cart/proto"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	// "github.com/micro/micro/v3/service/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const QPS = 100

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

	t, io, err := common.NewTracer("go.micro.service.cart", "192.168.56.101:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8084"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	srv.Init()

	db, err := gorm.Open("mysql", "root:Turkey414!@#@tcp(192.168.56.101:3306)/micro_user?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	rp := repository.NewCartRepository(db)
	rp.InitTable()

	cartDataService := service2.NewCartDataService(rp)

	// Register handler

	// pb.RegisterUserHandler(srv.Server(), new(handler.User))
	err = cart.RegisterCartHandler(srv.Server(), &handler.Cart{CartDataService: cartDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
