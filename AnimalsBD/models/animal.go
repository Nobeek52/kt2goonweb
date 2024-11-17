package animal

type Animal interface {
	Eat() string
	Sound() string
	Move() string
	Age() int
}

type Zebra struct {
	AgeValue int
}

func (g *Zebra) Sound() string {
	return "Зебра фыркает и ревет"
}

func (g *Zebra) Move() string {
	return "Зебра шагает"
}

func (g *Zebra) Age() int {
	return g.AgeValue
}

func (g *Zebra) Eat() string {
	return "Зебра ест траву"
}

type Tiger struct {
	AgeValue int
}

func (c *Tiger) Sound() string {
	return "Тигр рычит"
}

func (c *Tiger) Move() string {
	return "Тигр шагает"
}

func (c *Tiger) Age() int {
	return c.AgeValue
}

func (c *Tiger) Eat() string {
	return "Тигр ест животных"
}

type Panda struct {
	AgeValue int
}

func (k *Panda) Sound() string {
	return "Панда мычит"
}

func (k *Panda) Move() string {
	return "Панда шагает"
}

func (k *Panda) Age() int {
	return k.AgeValue
}

func (k *Panda) Eat() string {
	return "Панда ест бамбук"
}
