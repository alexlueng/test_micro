package handler

import (
	"context"
	"micro/cart/common"
	"micro/cart/domain/model"
	"micro/cart/domain/service"
	pb "micro/cart/proto"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c *Cart) AddCart(ctx context.Context, in *pb.CartInfo, out *pb.ResponseAdd) error {
	cart := &model.Cart{}
	common.SwapTo(in, cart)
	cartID, err := c.CartDataService.AddCart(cart)
	out.CartId = cartID
	return err
}
func (c *Cart) CleanCart(ctx context.Context, in *pb.Clean, out *pb.Response) error {
	if err := c.CartDataService.CleanCart(in.UserId); err != nil {
		return err
	}
	out.Msg = "Clean cart succeeded"
	return nil
}
func (c *Cart) Incr(ctx context.Context, in *pb.Item, out *pb.Response) error {
	if err := c.CartDataService.IncrNum(in.Id, in.ChangeNum); err != nil {
		return err
	}
	out.Msg = "Incr cart item succeeded"
	return nil
}
func (c *Cart) Decr(ctx context.Context, in *pb.Item, out *pb.Response) error {
	if err := c.CartDataService.DecrNum(in.Id, in.ChangeNum); err != nil {
		return err
	}
	out.Msg = "Decr cart item succeeded"
	return nil
}
func (c *Cart) DeleteItemByID(ctx context.Context, in *pb.CartID, out *pb.Response) error {
	if err := c.CartDataService.DeleteCart(in.CartId); err != nil {
		return nil
	}
	out.Msg = "Delete cart succeeded"
	return nil
}
func (c *Cart) GetAll(ctx context.Context, in *pb.CartFindAll, out *pb.CartAll) error {
	carts, err := c.CartDataService.FindAllCart(in.UserId)
	if err != nil {
		return err
	}
	for _, v := range carts {
		cart := &pb.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		out.CartInfo = append(out.CartInfo, cart)
	}
	return nil
}
