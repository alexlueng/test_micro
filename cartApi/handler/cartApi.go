package handler

import (
	cart "micro/cart/domain/service"
	pb "micro/cartApi/proto"
)

type CartApi struct {
	CartService cart.CartDataService
}

func (c *CartApi) FindAll(context.Context, in *pb.Request, out *pb.Response) error {

}