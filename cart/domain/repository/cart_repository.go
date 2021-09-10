package repository

import (
	"errors"
	"micro/cart/domain/model"

	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{
		mysqlDB: db,
	}
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDB.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) FindCartByID(id int64) (*model.Cart, error) {
	cart := &model.Cart{}
	return cart, c.mysqlDB.First(cart, id).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDB.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("Create cart error")
	}
	return cart.ID, nil
}

func (c *CartRepository) DeleteCartByID(id int64) error {
	return c.mysqlDB.Where("id = ?").Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDB.Model(cart).Update(cart).Error
}

func (c *CartRepository) FindAll(user_id int64) (cartAll []model.Cart, err error) {
	return cartAll, c.mysqlDB.Where("user_id = ?", user_id).Find(&cartAll).Error
}

func (c *CartRepository) CleanCart(user_id int64) error {
	return c.mysqlDB.Where("user_id = ?", user_id).Delete(&model.Cart{}).Error
}

func (c *CartRepository) IncrNum(cart_id int64, num int64) error {
	cart := &model.Cart{ID: cart_id}
	return c.mysqlDB.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (c *CartRepository) DecrNum(cart_id int64, num int64) error {
	cart := &model.Cart{ID: cart_id}
	db := c.mysqlDB.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("desc product item failed")
	}
	return nil
}
