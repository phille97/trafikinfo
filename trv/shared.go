package trv

import (
	"encoding/xml"
	"fmt"
	"time"
)

// ObjectType holds metadata about an Object you can query for in
// the Trafikverket API.
//
// The Kind is the object type, the version is the object version
// and the namespace, if necessary, that the object resides in.
// Namespaces are new to the API and are only present for certain
// versions of certain objects.
//
// The ObjectType method in each public package will return this
// object with all information correctly filled in. It can be passed
// as the argument to [code.dny.dev/trafikinfo.NewQuery].
type ObjectType struct {
	Kind      string
	Version   string
	Namespace string
}

// LastModified represents the latest ModifiedTime of any object that
// was included in the response.
type LastModified time.Time

// UnmarshalXML unmarshals into an internal type first, so that we can
// access the datetime attribute and lift it out. Since datetime is an
// attribute and not an element, we can't use an xml:"LASTMODIFIED>datetime"
// tag to get to it.
func (lm *LastModified) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type lastModified struct {
		DateTime *time.Time `xml:"datetime,attr"`
	}
	res := &lastModified{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	*lm = LastModified(*res.DateTime)
	return nil
}

// APIError is the error information that will be present if the API
// responds with an error to your request.
type APIError struct {
	XMLName xml.Name `xml:"ERROR"`
	Source  string   `xml:"SOURCE"`
	Message string   `xml:"MESSAGE"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Source, e.Message)
}
