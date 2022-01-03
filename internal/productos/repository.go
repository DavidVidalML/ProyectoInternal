package productos

import (
	"fmt"

	"github.com/DavidVidalML/ProyectoInternal/pkg/store"
)

type Producto struct {
	Id            int64   `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int64   `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var productos []Producto
var lastID int64

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int64, nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	LastID() (int64, error)
	Update(id int64, nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	UpdateNyP(id int64, nombre string, precio float64) (Producto, error)
	Delete(id int64) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Producto, error) {
	var prods []Producto
	r.db.Read(&prods)
	return prods, nil
}

func (r *repository) LastID() (int64, error) {
	var prods []Producto
	if err := r.db.Read(&prods); err != nil {
		return 0, err
	}
	if len(prods) == 0 {
		return 0, nil
	}
	return prods[len(prods)-1].Id, nil
}

func (r *repository) Store(id int64, nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	var prods []Producto
	prod := Producto{}
	prod.Id = id
	prod.Nombre = nombre
	prod.Color = color
	prod.Precio = precio
	prod.Stock = stock
	prod.Codigo = codigo
	prod.Publicado = publicado
	prod.FechaCreacion = fechaCreacion
	lastID = prod.Id
	r.db.Read(&prods)
	prods = append(prods, prod)
	if err := r.db.Write(prods); err != nil {
		return Producto{}, err
	}
	return prod, nil
}

func (r *repository) Update(id int64, nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	var prods []Producto
	prod := Producto{}
	prod.Nombre = nombre
	prod.Color = color
	prod.Precio = precio
	prod.Stock = stock
	prod.Codigo = codigo
	prod.Publicado = publicado
	prod.FechaCreacion = fechaCreacion
	updated := false
	r.db.Read(&prods)
	for i := range prods {
		if prods[i].Id == id {
			prod.Id = id
			prods[i] = prod
			updated = true
		}
	}
	if err := r.db.Write(prods); err != nil {
		return Producto{}, err
	}
	if !updated {
		return Producto{}, fmt.Errorf("No se encontro el producto con el id %d", id)
	}
	return prod, nil
}

func (r *repository) UpdateNyP(id int64, nombre string, precio float64) (Producto, error) {
	var prods []Producto
	prod := Producto{}
	updated := false
	r.db.Read(&prods)
	for i := range prods {
		if prods[i].Id == id {
			prods[i].Nombre = nombre
			prods[i].Precio = precio
			updated = true
			prod = prods[i]
		}
	}
	if err := r.db.Write(prods); err != nil {
		return Producto{}, err
	}
	if !updated {
		return Producto{}, fmt.Errorf("No se encontro el producto con el id %d", id)
	}
	return prod, nil
}

func (r *repository) Delete(id int64) error {
	var prods []Producto
	deleted := false
	var index int
	r.db.Read(&prods)
	for i := range prods {
		if prods[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("No se encontro el producto con el id %d", id)
	}
	prods = append(prods[:index], prods[index+1:]...)
	/*
		if err := os.Truncate("./products.json", 0); err != nil {
			return err
		}
	*/
	if err := r.db.Write(prods); err != nil {
		return err
	}
	return nil

}
