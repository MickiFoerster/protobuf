package main

import (
	"fmt"

	enumpb "github.com/MickiFoerster/protobuf/go/src/enum_example"
)

func doEnum() {
	em := enumpb.EnumMessage{
		Id:            42,
		DsayOfTheWeek: enumpb.DayOfTheWeek_TUESDAY,
	}

	fmt.Println(em)
}
