package service

import (
	"micro/category/domain/model"
	"micro/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategories() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

func NewUserDataService(rp repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{
		CategoryRepository: rp,
	}
}

func (c *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return c.CategoryRepository.CreateCategory(category)
}

func (c *CategoryDataService) DeleteCategory(id int64) error {
	return c.CategoryRepository.DeleteCategoryByID(id)
}

func (c *CategoryDataService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c *CategoryDataService) FindCategoryByID(id int64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByID(id)
}

func (c *CategoryDataService) FindAllCategories() ([]model.Category, error) {
	return c.CategoryRepository.FindAll()
}

func (c *CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByName(name)
}

func (c *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByLevel(level)
}

func (c *CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByParent(parent)
}
