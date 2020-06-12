package factory_simple

// 接口
type domain interface {
	do()
}

// 实例 A
type ADomain struct {
}

func (*ADomain) do() {
	// a domain do
}

// 实例 B
type BDomain struct {
}

func (*BDomain) do() {
	// b domain do
}

// 工厂方法
func factoryCreate(name string) domain {
	switch name {
	case "a":
		return &ADomain{}
	case "b":
		return &BDomain{}
	default:
		return nil
	}
}
