package main

import (
	"fmt"

	"github.com/binadel/jsonw/jsond"
)

func main() {
	obj := jsond.New(
		jsond.String("foo", "bar"),
		jsond.Object("nested",
			jsond.Number("big", "23565849841318736124854886"),
			jsond.Integer("weight", 25),
		),
		jsond.Array("price",
			jsond.FloatItem(25.97),
			jsond.NullItem(),
			jsond.BooleanItem(true),
		),
		jsond.Any("dynamite", "explodes"),
	)

	if out, err := obj.Write(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}
