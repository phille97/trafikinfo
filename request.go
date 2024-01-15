package trafikinfo

import "encoding/xml"

// Request tells the API what we're interested in
//
// It must include the Login information and at least
// one Query.
type Request struct {
	XMLName string  `xml:"REQUEST"`
	Login   *login  `xml:"LOGIN"`
	Queries []Query `xml:"QUERY"`
}

// Query adds one or more queries to the request
func (r *Request) Query(queries ...Query) *Request {
	r.Queries = append(r.Queries, queries...)
	return r
}

// APIKey sets the API key to use for this request
func (r *Request) APIKey(key string) *Request {
	r.Login = &login{AuthenticationKey: key}
	return r
}

// Build returns the XML encoded request as an array of bytes. It
// can be passed as http.NewRequest's body by wrapping it in a call
// to bytes.NewBuffer().
//
// The Build() method is final when used in a fluent API style,
// you can't chain additional methods on it that continue to modify
// the request.
func (r *Request) Build() ([]byte, error) {
	return xml.Marshal(r)
}

// login holds the API authentication key
type login struct {
	AuthenticationKey string `xml:"authenticationkey,attr"`
}

// NewRequest returns a Request using the specified API authentication key
// and the data to be retrieved and filtered by the specified queries. At
// least 1 query needs to be provided.
func NewRequest() *Request {
	return &Request{
		Queries: []Query{},
	}
}
