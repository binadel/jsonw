package jsond

import "github.com/binadel/jsonw/jsoni"

type Field interface {
	Write(writer *jsoni.ObjectWriter)
}

type Value interface {
	Write(writer *jsoni.ArrayWriter)
}
