package facade

import "fmt"

// facade provides apis that hide submodules' apis to make facade package more easy to use

// API for facade package
type API interface {
	Do() string
	Done() string
}

// NewApi to get the instance of API
func NewApi() API {
	return &implAPI{
		a: &AModule{},
		b: &BModule{},
	}
}

// Module A API
type IAModule interface {
	ADo() string
	ADone() string
}

// Module B API
type IBModule interface {
	BDo() string
	BDone() string
}

// implAPI the implement of API, it includes all apis from facade package
type implAPI struct {
	a IAModule
	b IBModule
}

func (i *implAPI) Do() string {
	return fmt.Sprintf("%s-%s", i.a.ADo(), i.b.BDo())
}

func (i *implAPI) Done() string {
	return fmt.Sprintf("%s-%s", i.a.ADone(), i.b.BDone())
}

// AModule the implement of IAModule
type AModule struct {
}

func (*AModule) ADo() string {
	return "A do"
}

func (*AModule) ADone() string {
	return "A done"
}

// BModule the implement of IAModule
type BModule struct {
}

func (*BModule) BDo() string {
	return "B do"
}

func (*BModule) BDone() string {
	return "B done"
}
