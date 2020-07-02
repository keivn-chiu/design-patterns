package mediator

import (
	"testing"
)

func TestName(t *testing.T) {
	m := GetMediator()
	m.ImageCard = &ImageCard{}
	m.AudioCard = &AudioCard{}
	m.Cpu = &CPU{}
	m.Driver = &UDiskDriver{}

	// trigger
	m.Driver.ReadData("image,audio")
}
