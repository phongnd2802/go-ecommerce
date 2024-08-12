//go:build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/phongnd2802/go-ecommerce/internal/controllers"
	database "github.com/phongnd2802/go-ecommerce/internal/database/sqlc"
	"github.com/phongnd2802/go-ecommerce/internal/repositories"
	"github.com/phongnd2802/go-ecommerce/internal/services"
)

func InitProductRouterHandler(db *sql.DB) (*controllers.ProductController, error) {
	wire.Build(
		database.NewStore,
		repositories.NewFurnitureRepository,
		repositories.NewElectronicRepository,
		repositories.NewClothingRepository,
		repositories.NewProductReposiroty,
		services.NewProductFactory,
		controllers.NewProductController,
	)

	return new(controllers.ProductController), nil
}