package factory_abstract

// 接口
type IADomain interface {
	Ping()
}

type IBDomain interface {
	Pong()
}

type AbsDomainFactory interface {
	CreateADomain() IADomain
	CreateBDomain() IBDomain
}

// 实例 A
type ADomainFactory struct {
	// A factory
}

func (*ADomainFactory) CreateADomain() IADomain {
	return &AADomain{}
}
func (*ADomainFactory) CreateBDomain() IBDomain {
	return &ABDomain{}
}

type AADomain struct {
}

func (*AADomain) Ping() {
	println("AADomain ping")
}

type ABDomain struct {
}

func (*ABDomain) Pong() {
	println("ABDomain pong")
}

// 实例 B
type BDomainFactory struct {
	// B factory
}

func (*BDomainFactory) CreateADomain() IADomain {
	return &BADomain{}
}
func (*BDomainFactory) CreateBDomain() IBDomain {
	return &BBDomain{}
}

type BADomain struct {
}

func (*BADomain) Ping() {
	println("BADomain ping")
}

type BBDomain struct {
}

func (*BBDomain) Pong() {
	println("BBDomain pong")
}
