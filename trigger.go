package trigger

import (
	"fmt"
	"io/ioutil"
)

func trigger() {
	if _, err := os.Stat(recovery.conf); os.IsNotExist(err) {
		fmt.Printf("I'm not the master shonuff")
		return
	}
	// This file would be to move the recovery.conf file from slave to another slave.
	// make agents to talk to each other and if necessary write recovery.conf
}
