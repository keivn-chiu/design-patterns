package command

import "testing"

func TestDemo(t *testing.T) {
	browser := NewBrowser()
	c1 := NewLoadUrlCommand(browser, "www.baidu.com")
	c2 := NewOpenTabCommand(browser)
	b1 := NewBox(c1, c2)
	b2 := NewBox(c2, c1)
	// bi execute functions
	b1.ExecC1()
	b1.ExecC2()
	// b2 execute functions
	b2.ExecC1()
	b2.ExecC2()
}
