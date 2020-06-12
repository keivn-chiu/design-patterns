package factroy_general

import "testing"

func TestDemo(t *testing.T) {
	// 使用ADomain
	a := &ADomainFactory{}
	a.CreateDomain().DoSomething()
	// 使用BDomain
	b := &BDomainFactory{}
	b.CreateDomain().DoSomething()
}

/*
运行结果
=== RUN   TestDemo
A do something
B do something
--- PASS: TestDemo (0.00s)
PASS
ok      kevin-chiu/design-patterns/factroy-general      0.006s
*/