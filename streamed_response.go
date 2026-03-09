package trafikinfo

import (
	"encoding/xml"
	"io"
	"iter"

	"github.com/phille97/trafikinfo/trv"
)

type StreamedResult[T trv.Object] struct {
	Data  T
	Error *trv.APIError
}

func StreamResult[T trv.Object](r io.Reader) iter.Seq2[StreamedResult[T], error] {
	var u T
	xmlLocalName := u.T().Kind
	return func(yield func(StreamedResult[T], error) bool) {
		decoder := xml.NewDecoder(r)
		for {
			token, err := decoder.Token()
			if err == io.EOF {
				return
			}
			if err != nil {
				yield(StreamedResult[T]{}, err)
				return
			}
			se, ok := token.(xml.StartElement)
			if !ok {
				continue
			}

			switch se.Name.Local {
			case "ERROR":
				var apiErr trv.APIError
				if err := decoder.DecodeElement(&apiErr, &se); err != nil {
					yield(StreamedResult[T]{}, err)
					return
				}
				if !yield(StreamedResult[T]{Error: &apiErr}, nil) {
					return
				}
			case xmlLocalName:
				var u T
				if err := decoder.DecodeElement(&u, &se); err != nil {
					if !yield(StreamedResult[T]{}, err) {
						return
					}
					continue
				}
				if !yield(StreamedResult[T]{Data: u}, nil) {
					return
				}
			}
		}
	}
}
