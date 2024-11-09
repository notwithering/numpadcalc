package main

import (
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func listen(conn *xgb.Conn) {
	for {
		ev, err := conn.WaitForEvent()
		if err != nil {
			log.Fatal(err)
		}

		switch event := ev.(type) {
		case xproto.KeyPressEvent:
			keyCh <- event.Detail
		}
	}
}
