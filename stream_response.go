package trafikinfo

import (
	"encoding/xml"
	"io"
	"iter"

	"github.com/phille97/trafikinfo/trv"
)

func StreamResponse[T trv.Object](r io.Reader) iter.Seq2[T, error] {
	var zero T
	xmlName := zero.XMLName()
	return func(yield func(T, error) bool) {
		decoder := xml.NewDecoder(r)
		for {
			token, err := decoder.Token()
			if err == io.EOF {
				return
			}
			if err != nil {
				yield(zero, err)
				return
			}
			se, ok := token.(xml.StartElement)
			if !ok {
				continue
			}

			switch se.Name {
			case xml.Name{Local: "ERROR"}:
				var apiErr trv.APIError
				if err := decoder.DecodeElement(&apiErr, &se); err != nil {
					yield(zero, err)
					return
				}
				if !yield(zero, &apiErr) {
					return
				}
			case xmlName:
				var u T
				if err := decoder.DecodeElement(&u, &se); err != nil {
					yield(u, err)
					return
				}
				if !yield(u, nil) {
					return
				}
			}
		}
	}
}
