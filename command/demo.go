package command

import "fmt"

// Command mode it collects the functions of a object into a command object,
// the command object is easy to transmit, persistence and use.
// area: batch processing, undo, redo ..

// Browser the type provides functions for packaging
type Browser struct {
}

func NewBrowser() *Browser {
	return &Browser{}
}
func (b *Browser) Refresh() {
	fmt.Println("Browser -> refresh")
}
func (b *Browser) LoadUrl(url string) {
	fmt.Printf("Browser -> Load url:%s\n", url)
}

// ICommand the interface of Command types
type ICommand interface {
	Execute()
}

// RefreshCommand the command wrapper
type RefreshCommand struct {
	*Browser
}

func NewOpenTabCommand(b *Browser) *RefreshCommand {
	return &RefreshCommand{b}
}
func (o *RefreshCommand) Execute() {
	o.Browser.Refresh()
}

// LoadUrlCommand the command wrapper
type LoadUrlCommand struct {
	*Browser
	url string
}

func NewLoadUrlCommand(b *Browser, url string) *LoadUrlCommand {
	return &LoadUrlCommand{b, url}
}
func (l *LoadUrlCommand) Execute() {
	l.Browser.LoadUrl(l.url)
}

// Box the box collects all functions of Browser
type Box struct {
	c1 ICommand
	c2 ICommand
}

func NewBox(c1, c2 ICommand) *Box {
	return &Box{c1, c2}
}
func (b *Box) ExecC1() {
	b.c1.Execute()
}
func (b *Box) ExecC2() {
	b.c2.Execute()
}
