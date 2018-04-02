package tag

// ID -
type ID = int

// Tag -
type Tag struct {
	ID         ID     `json:"id" sql:"id"`
	Name       string `json:"name" sql:"name"`
	Desc       string `json:"desc" sql:"description"`
	CreateTime int    `json:"craete_time" sql:"create_time"`
	UpdateTime int    `json:"update_time" sql:"update_time"`
}

// Model -
type Model struct{}
