package trafikinfo

// Request tells the API what we're interested in
//
// It must include the Login information and at least
// one Query.
type Request struct {
	XMLName string   `xml:"REQUEST"`
	Login   *login   `xml:"LOGIN"`
	Queries []*Query `xml:"QUERY"`
}

// Query adds one or more queries to the request
func (r *Request) Query(query *Query, rest ...*Query) *Request {
	r.Queries = append(r.Queries, query)
	r.Queries = append(r.Queries, rest...)
	return r
}

// APIKey sets the API key to use for this request
func (r *Request) APIKey(key string) *Request {
	r.Login = &login{AuthenticationKey: key}
	return r
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
		Queries: []*Query{},
	}
}
