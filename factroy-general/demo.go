package factroy_general

// 实例接口
type Domain interface {
	DoSomething()
}

// A Domain
type ADomain struct {
}

func (*ADomain) DoSomething() {
	println("A do something")
}

// B Domain
type BDomain struct {
}

func (*BDomain) DoSomething() {
	println("B do something")
}

// 工厂抽象类
type DomainFactory interface {
	CreateDomain() Domain
}

// 工厂实例类
type ADomainFactory struct {
	// A Domain 的工厂实例
}

func (*ADomainFactory) CreateDomain() Domain {
	return &ADomain{}
}

type BDomainFactory struct {
	// B Domain 的工厂实例
}

func (*BDomainFactory) CreateDomain() Domain {
	return &BDomain{}
}
