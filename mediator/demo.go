package mediator

import (
	"fmt"
	"strings"
	"sync"
)

// mediator makes data transfer easier between different types instances,
// the complex transfers are packaged in the type of mediator

// Define the types need data transfer together, here is the demo about playing a movie on a PC
type Mediator struct {
	Cpu       *CPU
	Driver    *UDiskDriver
	AudioCard *AudioCard
	ImageCard *ImageCard
}

func (m *Mediator) Change(obj IPass) {
	switch ins := obj.(type) {
	case *CPU:
		m.AudioCard.SetAudio(ins.Audio)
		m.ImageCard.SetImage(ins.Video)
	case *UDiskDriver:
		m.Cpu.Process(ins.data)
	case *AudioCard:
		ins.Play()
	case *ImageCard:
		ins.Play()
	}
}

// singleton of mediator
var (
	mediator *Mediator
	once     sync.Once
)

func GetMediator() *Mediator {
	once.Do(func() {
		mediator = &Mediator{}
	})
	return mediator
}

// IPass the interface for the types need data transfer via mediator
type IPass interface {
	pass(mediator *Mediator)
}

// UDiskDriver it can provide data to CPU
type UDiskDriver struct {
	data string
}

func (driver *UDiskDriver) ReadData(data string) {
	driver.data = data
	driver.pass(GetMediator())
}
func (driver *UDiskDriver) pass(mediator *Mediator) {
	mediator.Change(driver)
}

// CPU needs to pass the data to AudioCard and VideoCard
type CPU struct {
	Video string
	Audio string
}

func (cpu *CPU) Process(data string) {
	d := strings.Split(data, ",")
	cpu.Video = d[0]
	cpu.Audio = d[1]
	cpu.pass(GetMediator())
}
func (cpu *CPU) pass(mediator *Mediator) {
	mediator.Change(cpu)
}

// AudioCard plays music
type AudioCard struct {
	Audio string
}

func (ac *AudioCard) pass(mediator *Mediator) {
	mediator.Change(ac)
}
func (ac *AudioCard) SetAudio(audio string) {
	ac.Audio = audio
	ac.pass(GetMediator())
}
func (ac *AudioCard) Play() {
	fmt.Println("play audio:", ac.Audio)
}

//ImageCard plays image
type ImageCard struct {
	Image string
}

func (i *ImageCard) pass(mediator *Mediator) {
	mediator.Change(i)
}
func (i *ImageCard) SetImage(data string) {
	i.Image = data
	i.pass(GetMediator())
}
func (i *ImageCard) Play() {
	fmt.Println("play image:", i.Image)
}
