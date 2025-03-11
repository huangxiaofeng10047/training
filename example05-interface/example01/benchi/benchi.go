package benchi

type Benchi struct {
	name     string
	price    float64
	discount float64
	color    string
}

func (l Benchi) Name() string {
	return l.name
}

func (l Benchi) Price() float64 {
	return l.price * l.discount
}

func (l Benchi) Discount() float64 {
	return l.discount
}

func (l Benchi) Color() string {
	return l.color
}

// NewCar contructure
func NewCar(name string, price float64, discount float64, color string) *Benchi {
	return &Benchi{
		name:     name,
		price:    price,
		discount: discount,
		color:    color,
	}
}
