package service

import (
	"micro/cart/domain/model"
	"micro/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func NewCartDataService(rp repository.ICartRepository) ICartDataService {
	return &CartDataService{
		CartRepository: rp,
	}
}

func (c *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return c.CartRepository.CreateCart(cart)
}

func (c *CartDataService) DeleteCart(id int64) error {
	return c.CartRepository.DeleteCartByID(id)
}

func (c *CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartRepository.UpdateCart(cart)
}

func (c *CartDataService) FindCartByID(id int64) (*model.Cart, error) {
	return c.CartRepository.FindCartByID(id)
}

func (c *CartDataService) FindAllCart(user_id int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(user_id)
}

func (c *CartDataService) CleanCart(user_id int64) error {
	return c.CartRepository.CleanCart(user_id)
}

func (c *CartDataService) DecrNum(cart_id int64, num int64) error {
	return c.CartRepository.DecrNum(cart_id, num)
}

func (c *CartDataService) IncrNum(cart_id int64, num int64) error {
	return c.CartRepository.IncrNum(cart_id, num)
}
