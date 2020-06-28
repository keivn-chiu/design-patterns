package flyweight

import (
	"fmt"
	"sync"
)

// flyweight mode to share the data that no need to recreate some same obj

// IData the interface for data shared
type IData interface {
	Name() string
	Date() string
}

// dataImpl the type implement IData
type dataImpl struct {
	name string
	data string
}

func (data *dataImpl) Name() string {
	return data.name
}
func (data *dataImpl) Date() string {
	return fmt.Sprintf("%s => data\n", data.name)
}
func NewData(name string) *dataImpl {
	return &dataImpl{name: name}
}

// factory to create or share dataImpl
type factory struct {
	m map[string]IData
}

func (f *factory) GetData(name string) IData {
	data := f.m[name]
	if data == nil {
		// save the new data here
		data = NewData(name)
		f.m[name] = data
	}
	return data
}

// factory should be singleton
var (
	f   *factory
	one sync.Once
)

func GetFactory() *factory {
	one.Do(func() {
		f = &factory{make(map[string]IData)}
	})
	return f
}

// DataViewer the struct for reading dataImpl
type DataViewer struct {
	IData
}

func (dv *DataViewer) ReadData() {
	fmt.Printf("Name:%s, Data:%s", dv.IData.Name(), dv.IData.Date())
}

func NewDataViewer(name string) *DataViewer {
	return &DataViewer{GetFactory().GetData(name)}
}
