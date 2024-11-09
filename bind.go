package main

import (
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

var keyMap = map[xproto.Keycode]string{
	90:  "0",
	87:  "1",
	88:  "2",
	89:  "3",
	83:  "4",
	84:  "5",
	85:  "6",
	79:  "7",
	80:  "8",
	81:  "9",
	91:  ".",
	86:  "+",
	82:  "-",
	63:  "*",
	106: "/",
	104: "enter",
}

func bind() *xgb.Conn {
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root

	for keycode := range keyMap {
		err = xproto.GrabKeyChecked(
			conn,
			true,
			root,
			xproto.ModMaskAny,
			keycode,
			xproto.GrabModeAsync,
			xproto.GrabModeAsync,
		).Check()
		if err != nil {
			log.Fatal(err)
		}
	}

	return conn
}
