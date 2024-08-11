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


func InitAccessRouterHandler(db *sql.DB) (*controllers.AccessController, error) {
	wire.Build(
		database.NewStore,
		repositories.NewTokenRepository,
		repositories.NewShopRepository,
		services.NewAccessService,
		controllers.NewAccessController,
	)
	return new(controllers.AccessController), nil
}