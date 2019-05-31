package data

// ResponseData struct
type ResponseData struct {
	Neighbours     *Neighbours     `json:"neighbours"`
	Nodeinfo       *Nodeinfo       `json:"nodeinfo"`
	Statistics     *Statistics     `json:"statistics"`
	DomainDirector *DomainDirector `json:"domain_director,omitempty"`
}
