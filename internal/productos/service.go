package productos

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int64, nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	UpdateNyP(id int64, nombre string, precio float64) (Producto, error)
	Delete(id int64) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetAll() ([]Producto, error) {
	productos, err := s.repo.GetAll()
	if err != nil {
		panic(err)
	}
	return productos, nil
}

func (s *service) Store(nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	lastId, err := s.repo.LastID()
	if err != nil {
		return Producto{}, err
	}
	lastId++
	prod, err := s.repo.Store(lastId, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		panic(err)
	}
	return prod, nil
}

func (s *service) Update(id int64, nombre string, color string, precio float64, stock int64, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	return s.repo.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) UpdateNyP(id int64, nombre string, precio float64) (Producto, error) {
	return s.repo.UpdateNyP(id, nombre, precio)
}
func (s *service) Delete(id int64) error {
	return s.repo.Delete(id)
}
