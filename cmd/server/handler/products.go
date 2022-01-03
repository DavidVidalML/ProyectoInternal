package handler

//VALIDACIONES DEL REQUEST
import (
	"fmt"
	"os"
	"strconv"

	"github.com/DavidVidalML/ProyectoInternal/internal/productos"
	"github.com/DavidVidalML/ProyectoInternal/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int64   `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}
type Producto struct {
	service productos.Service
}

func NewProducto(prod productos.Service) *Producto {
	return &Producto{
		service: prod,
	}
}

// ListProducts godoc
// @Summary Listar Productos de API
// @Tags Products
// @Description getAll products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary cargar productos a la API
// @Tags Products
// @Description post products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		p, err := c.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// UpdateProduct godoc
// @Summary Update products
// @Tags Products
// @Description updatear productos
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/id [put]
func (c *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Token invalido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		//VALIDACIONES DE NULOS
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del proudcto es requerido"))
			return
		}
		if req.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}

		p, err := c.service.Update(int64(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// UpdateNyP godoc
// @Summary UpdateNyP products
// @Tags Products
// @Description updatear nombre y precio productos
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to update NyP"
// @Success 200 {object} web.Response
// @Router /products/id [patch]
func (c *Producto) UpdateNyP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Token invalido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		//VALIDACIONES
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}
		p, err := c.service.UpdateNyP(int64(id), req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(400, p, ""))
	}
}

// Delete godoc
// @Summary Delete products
// @Tags Products
// @Description deletear  productos
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to delete"
// @Success 200 {object} web.Response
// @Router /products/id [delete]
func (c *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		err = c.service.Delete(int64(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto con el ID %d ha sido eliminado", id), ""))
	}

}
