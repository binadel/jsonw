package main

import (
	"fmt"

	json "github.com/binadel/jsonw/jsondf" // or jsondi, or jsonds
)

func main() {
	obj := json.New(
		json.String("foo", "bar"),
		json.Integer("cold", -273),
		json.Object("nested",
			json.Number("big", "23565849841318736104"),
			json.Boolean("flag", true),
			json.Array("items",
				json.IntegerItem(732),
				json.StringItem("is bigger than"),
				json.FloatItem(73.2),
				json.ArrayItem(
					json.AnyItem(map[string]string{"deeply": "hidden"}),
				),
			),
		),
		json.Float("pi", 3.1415),
		json.Array("same",
			json.BooleanItem(false),
			json.NullItem(),
			json.ObjectItem(),
			json.NumberItem("0.0"),
		),
		json.Any("dynamite", "explodes"),
		json.Null("last"),
	)

	if out, err := obj.Build(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}
