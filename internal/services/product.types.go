package services

import (
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
)

type IProduct interface {
	CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error)
}


////////////////////////////////////
/////// 	Product			////////
//////////////////////////////////

type product struct {
	productRepo repositories.IProductRepository
}

func NewProduct(productRepo repositories.IProductRepository) IProduct {
	return &product{
		productRepo: productRepo,
	}
}

func (p *product) CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error) {
	result, err := p.productRepo.CreateProduct(
		payload.ProductName, payload.ProductThumb, &payload.ProductDescription, payload.ProductPrice,
		payload.ProductQuantity, payload.ProductType, productShop, payload.ProductAttributes,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}




////////////////////////////////////
/////// 	Clothing		///////
//////////////////////////////////

type clothing struct {
	product IProduct
	clothingRepo repositories.IClothingRepository
}


func NewClothing(product IProduct, clothingRepo repositories.IClothingRepository) IProduct {
	return &clothing{
		product: product,
		clothingRepo: clothingRepo,
	}
}

func (c *clothing) CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error) {
	newProduct, err := c.product.CreateProduct(payload, productShop)
	if err != nil {
		return nil, err
	}

	brand := payload.ProductAttributes["brand"].(string)
	size := payload.ProductAttributes["size"].(string)
	material := payload.ProductAttributes["material"].(string)
	_, err = c.clothingRepo.CreateClothing(newProduct.ID, brand, size, material, productShop)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
} 


////////////////////////////////////
/////// 	Electronics		////////
//////////////////////////////////
type electronic struct {
	product IProduct
	electronicRepo repositories.IElectronicsRepository
}

func NewElectronic(
	product IProduct,
	electronicRepo repositories.IElectronicsRepository,
) IProduct {
	return &electronic{
		product: product,
		electronicRepo: electronicRepo,
	}
}

func (e *electronic) CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error) {
	newProduct, err := e.product.CreateProduct(payload, productShop)
	if err != nil {
		return nil, err
	}

	manufacturer := payload.ProductAttributes["manufacturer"].(string)
	model := payload.ProductAttributes["model"].(string)
	color := payload.ProductAttributes["color"].(string)
	_, err = e.electronicRepo.CreateElectronic(newProduct.ID, manufacturer, model, color, productShop)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}



////////////////////////////////////
/////// 	Furniture		////////
//////////////////////////////////

type furniture struct {
	product IProduct
	furnitureRepo repositories.IFurnitureRepository
}

func NewFurniture(product IProduct, furnitureRepo repositories.IFurnitureRepository) IProduct {
	return &furniture{
		product: product,
		furnitureRepo: furnitureRepo,
	}
}

func (f *furniture) CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error) {
	newProduct, err := f.product.CreateProduct(payload, productShop)
	if err != nil {
		return nil, err
	}

	brand := payload.ProductAttributes["brand"].(string)
	size := payload.ProductAttributes["size"].(string)
	material := payload.ProductAttributes["material"].(string)
	_, err = f.furnitureRepo.CreateFurniture(newProduct.ID, brand, size, material, productShop)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}