package builder

// interfaces of builder
type IBuilder interface {
	Part1()
	Part2()
	Part3()
}

// the builder
type Director struct {
	IBuilder
}

// func to get a director with builder
func NewDirector(builder IBuilder) *Director {
	return &Director{builder}
}

// start building the obj according to builder
func (d *Director) Construct() {
	d.IBuilder.Part1()
	d.IBuilder.Part2()
	d.IBuilder.Part3()
}

// struct implement builder
type Builder struct {
	result string
}

func (b *Builder) Part1() {
	b.result += "a"
}
func (b *Builder) Part2() {
	b.result += "b"
}
func (b *Builder) Part3() {
	b.result += "c"
}
