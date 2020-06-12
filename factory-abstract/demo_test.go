package factory_abstract

import "testing"

func TestDemo(t *testing.T) {
	var absFactory AbsDomainFactory
	absFactory = &ADomainFactory{}
	absFactory.CreateADomain().Ping()
	absFactory.CreateBDomain().Pong()
	println("-----------------------")
	absFactory = &BDomainFactory{}
	absFactory.CreateADomain().Ping()
	absFactory.CreateBDomain().Pong()
}

/*
=== RUN   TestDemo
AADomain ping
ABDomain pong
-----------------------
BADomain ping
BBDomain pong
--- PASS: TestDemo (0.00s)
PASS
*/
