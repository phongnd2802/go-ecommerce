package services

import (
	"github.com/gosimple/slug"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/internal/dtos"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
)

type IProduct interface {
	CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error)
	UpdateProduct(bodyUpdate dtos.ProductUpdateRequest, productID string) (*database.Product, error)
}

////////////////////////////////////
/////// 	Product			////////
//////////////////////////////////

type product struct {
	productRepo repositories.IProductRepository
}

// UpdateProduct implements IProduct.
func (p *product) UpdateProduct(bodyUpdate dtos.ProductUpdateRequest, productID string) (*database.Product, error) {
	productSlug := slug.Make(bodyUpdate.ProductName)
	result, err := p.productRepo.UpdateProductByID(
		productID, bodyUpdate.ProductName, bodyUpdate.ProductThumb,
		&bodyUpdate.ProductDescription, bodyUpdate.ProductPrice, bodyUpdate.ProductQuantity,
		bodyUpdate.ProductType, productSlug, bodyUpdate.ProductAttributes,
	)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (p *product) CreateProduct(payload dtos.ProductCreateRequest, productShop string) (*database.Product, error) {
	productSlug := slug.Make(payload.ProductName)
	result, err := p.productRepo.CreateProduct(
		payload.ProductName, payload.ProductThumb, &payload.ProductDescription, payload.ProductPrice,
		payload.ProductQuantity, payload.ProductType, productShop, productSlug, payload.ProductAttributes,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewProduct(productRepo repositories.IProductRepository) IProduct {
	return &product{
		productRepo: productRepo,
	}
}

////////////////////////////////////
/////// 	Clothing		///////
//////////////////////////////////

type clothing struct {
	product      IProduct
	clothingRepo repositories.IClothingRepository
}

// UpdateProduct implements IProduct.
func (c *clothing) UpdateProduct(bodyUpdate dtos.ProductUpdateRequest, productID string) (*database.Product, error) {
	if bodyUpdate.ProductAttributes != nil {
		// Update Clothing
		foundProduct, err := c.clothingRepo.GetClothingByID(productID)
		if err != nil {
			return nil, err
		}

		var brand, size, material string

		if v, ok := bodyUpdate.ProductAttributes["brand"].(string); ok && v != "" {
			brand = v
		} else {
			brand = foundProduct.Brand
		}


		if v, ok := bodyUpdate.ProductAttributes["size"].(string); ok && v != "" {
			size = v
		} else {
			size = foundProduct.Size
		}


		if v, ok := bodyUpdate.ProductAttributes["material"].(string); ok && v != "" {
			material = v
		} else {
			material = foundProduct.Material
		}

		err = c.clothingRepo.UpdateClothing(productID, brand, size, material)
		if err != nil {
			return nil, err
		}
	}
	// Update Product
	updatedProduct, err := c.product.UpdateProduct(bodyUpdate, productID)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil


}

func NewClothing(product IProduct, clothingRepo repositories.IClothingRepository) IProduct {
	return &clothing{
		product:      product,
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

// //////////////////////////////////
// ///// 	Electronics		////////
// ////////////////////////////////
type electronic struct {
	product        IProduct
	electronicRepo repositories.IElectronicsRepository
}

// UpdateProduct implements IProduct.
func (e *electronic) UpdateProduct(bodyUpdate dtos.ProductUpdateRequest, productID string) (*database.Product, error) {
	panic("unimplemented")
}

func NewElectronic(
	product IProduct,
	electronicRepo repositories.IElectronicsRepository,
) IProduct {
	return &electronic{
		product:        product,
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
	product       IProduct
	furnitureRepo repositories.IFurnitureRepository
}

// UpdateProduct implements IProduct.
func (f *furniture) UpdateProduct(bodyUpdate dtos.ProductUpdateRequest, productID string) (*database.Product, error) {
	panic("unimplemented")
}

func NewFurniture(product IProduct, furnitureRepo repositories.IFurnitureRepository) IProduct {
	return &furniture{
		product:       product,
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
