package decorator

// decorator helping add feature or changing behaviors of type via compositional approach

// IComponent the interface for decorator and the type need to be decorated
type IComponent interface {
	Calc() int // calculation
}

// Component the implement of IComponent, the base type
type Component struct {
}

func (c *Component) Calc() int {
	return 0
}

func NewComponent() IComponent {
	return &Component{}
}

// WarpMulComponent the decorator for multiple
type WarpMulComponent struct {
	num int
	IComponent
}

func (m *WarpMulComponent) Calc() int {
	return m.IComponent.Calc() * m.num
}

func NewWarpMulComponent(c IComponent, num int) IComponent {
	return &WarpMulComponent{
		num:        num,
		IComponent: c,
	}
}

// WarpAddComponent the decorator for add
type WarpAddComponent struct {
	num int
	IComponent
}

func (m *WarpAddComponent) Calc() int {
	return m.IComponent.Calc() + m.num
}

func NewWarpAddComponent(c IComponent, num int) IComponent {
	return &WarpAddComponent{
		num:        num,
		IComponent: c,
	}
}
