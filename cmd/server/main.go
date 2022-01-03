package main

import (
	"os"

	"github.com/DavidVidalML/ProyectoInternal/cmd/server/handler"
	"github.com/DavidVidalML/ProyectoInternal/docs"
	"github.com/DavidVidalML/ProyectoInternal/internal/productos"
	"github.com/DavidVidalML/ProyectoInternal/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API GO
// @version 1.0
// @description API para productos de MELI Bootcamp
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name David
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "products.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	p := handler.NewProducto(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNyP())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
