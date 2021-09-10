package handler

import (
	"context"
	"micro/category/common"
	"micro/category/domain/model"
	"micro/category/domain/service"
	pb "micro/category/proto"

	"github.com/micro/go-micro/v2/util/log"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, in *pb.CategoryRequest, out *pb.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(in, category)
	if err != nil {
		return err
	}
	categoryID, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	out.CategoryId = categoryID
	out.Message = "add category succeeded"
	return nil
}

func (c *Category) UpdateCategory(ctx context.Context, in *pb.CategoryRequest, out *pb.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(in, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	out.Message = "Update category succeeded"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, in *pb.DeleteCategoryRequest, out *pb.DeleteCategoryResponse) error {
	categoryId := in.CategoryId
	err := c.CategoryDataService.DeleteCategory(categoryId)
	if err != nil {
		return err
	}
	out.Message = "Delete category succeeded"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, in *pb.FindByNameRequest, out *pb.CategoryResponse) error {
	name := in.CategoryName
	category, err := c.CategoryDataService.FindCategoryByName(name)
	if err != nil {
		return err
	}
	return common.SwapTo(category, out)
}

func (c *Category) FindCategoryByID(ctx context.Context, in *pb.FindByIDRequest, out *pb.CategoryResponse) error {
	id := in.CategoryId
	category, err := c.CategoryDataService.FindCategoryByID(id)
	if err != nil {
		return err
	}
	return common.SwapTo(category, out)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, in *pb.FindByLevelRequest, out *pb.FindAllResponse) error {
	level := in.CategoryLevel
	categories, err := c.CategoryDataService.FindCategoryByLevel(level)
	if err != nil {
		return err
	}
	categoryToResponse(categories, out)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, in *pb.FindByParentRequest, out *pb.FindAllResponse) error {
	parent := in.CategoryParent
	categories, err := c.CategoryDataService.FindCategoryByParent(parent)
	if err != nil {
		return err
	}
	categoryToResponse(categories, out)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, in *pb.FindAllRequest, out *pb.FindAllResponse) error {
	categories, err := c.CategoryDataService.FindAllCategories()
	if err != nil {
		return err
	}
	categoryToResponse(categories, out)
	return nil
}

func categoryToResponse(categories []model.Category, resp *pb.FindAllResponse) {
	for _, cat := range categories {
		cr := &pb.CategoryResponse{}
		err := common.SwapTo(cat, cr)
		if err != nil {
			log.Error(err)
			break
		}
		resp.Category = append(resp.Category, cr)
	}
}
