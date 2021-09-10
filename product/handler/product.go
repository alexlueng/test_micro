package handler

import (
	"context"
	"micro/product/common"
	"micro/product/domain/model"
	"micro/product/domain/service"
	pb "micro/product/proto"

	"github.com/micro/go-micro/v2/logger"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, in *pb.ProductInfo, out *pb.ResponseProduct) error {
	product := &model.Product{}
	if err := common.SwapTo(in, product); err != nil {
		return err
	}
	productID, err := p.ProductDataService.AddProduct(product)
	if err != nil {
		return err
	}
	out.Id = productID
	return nil
}

func (p *Product) FindProductByID(ctx context.Context, in *pb.RequestID, out *pb.ProductInfo) error {
	product, err := p.ProductDataService.FindProductByID(in.Id)
	if err != nil {
		return err
	}
	if err := common.SwapTo(product, out); err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, in *pb.ProductInfo, out *pb.Response) error {
	product := &model.Product{}
	if err := common.SwapTo(in, product); err != nil {
		return err
	}
	err := p.ProductDataService.UpdateProduct(product)
	if err != nil {
		return err
	}
	out.Message = "Update Product Succeeded"
	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, in *pb.RequestID, out *pb.Response) error {
	id := in.Id
	if err := p.ProductDataService.DeleteProduct(id); err != nil {
		return err
	}
	out.Message = "Delete Product Succeeded"
	return nil
}

func (p *Product) FindAllProducts(ctx context.Context, in *pb.RequestAll, out *pb.AllProducts) error {
	products, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	productToResponse(products, out)
	return nil
}

func productToResponse(products []model.Product, resp *pb.AllProducts) {
	for _, p := range products {
		pr := &pb.ProductInfo{}
		err := common.SwapTo(p, pr)
		if err != nil {
			logger.Error(err)
			break
		}
		resp.ProductsInfo = append(resp.ProductsInfo, pr)
	}
}
