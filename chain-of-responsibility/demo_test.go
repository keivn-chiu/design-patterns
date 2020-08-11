package chain_of_responsibility

import "testing"

func TestDemo(t *testing.T) {
	level1 := NewRequestLevel1()
	level2 := NewRequestLevel2()
	level3 := NewRequestLevel3()

	level1.SetSuccessor(level2)
	level2.SetSuccessor(level3)

	var manager Manager = level1

	manager.HandleRequest("peter", 1)
	manager.HandleRequest("pop", 2)
	manager.HandleRequest("may", 3)
	manager.HandleRequest("kevin", 2)
}

/*
result:
peter is level 1, job done
pop is level 2, job done
may is level 3, job done
kevin can't handle the request, level is 2
*/
