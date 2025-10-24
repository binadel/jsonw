package main

import (
	"fmt"

	"github.com/binadel/jsonw/jsoni"
)

func main() {
	w := jsoni.NewObjectWriter(nil)
	w.Open()
	w.StringField("foo", "bar")
	w.IntegerField("cold", -273)
	nested := w.ObjectField("nested")
	nested.Open()
	nested.NumberField("big", "23565849841318736104")
	nested.BooleanField("flag", true)
	items := nested.ArrayField("items")
	items.Open()
	items.IntegerValue(732)
	items.StringValue("is bigger than")
	items.FloatValue(73.2)
	arr := items.ArrayValue()
	arr.Open()
	arr.AnyValue(map[string]string{"deeply": "hidden"})
	arr.Close()
	items.Close()
	nested.Close()
	w.FloatField("pi", 3.1415)
	same := w.ArrayField("same")
	same.Open()
	same.BooleanValue(false)
	same.NullValue()
	empty := same.ObjectValue()
	empty.Open()
	empty.Close()
	same.NumberValue("0.0")
	same.Close()
	w.AnyField("dynamite", "explodes")
	w.NullField("last")
	w.Close()

	if out, err := w.BuildBytes(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}
