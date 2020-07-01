package bridge

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	NewCommonMsg("kevin", NewMsgEmail()).SendMsg("kobe", "do something?")
	fmt.Println("----")
	NewCommonMsg("kevin", NewMsgMSM()).SendMsg("kobe", "do something?")
	fmt.Println("----")
	NewUrgentMsg("kobe", NewMsgMSM()).SendMsg("mary", "what?")
	fmt.Println("----")
	NewUrgentMsg("kobe", NewMsgEmail()).SendMsg("mary", "what?")
}

/*
common msg from: kevin
email[ to kobe : do something?]
----
common msg from: kevin
msm[ to kobe : do something?]
----
urgent msg from: kobe
msm[ to mary : what?]
----
urgent msg from: kobe
email[ to mary : what?]
*/
