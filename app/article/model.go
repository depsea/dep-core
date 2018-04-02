package article

import (
	"github.com/depsea/api/app/tag"
)

// ID -
type ID = int

// Status -
type Status = int

const (
	// Disabled -
	Disabled Status = iota

	// Publish -
	Publish
)

// Article -
type Article struct {
	ID          ID       `json:"id" sql:"id"`
	Title       string   `json:"title" sql:"title"`
	Desc        string   `json:"desc" sql:"description"`
	Body        string   `json:"body" sql:"body"`
	Tags        []tag.ID `json:"tags" sql:"tags"`
	CreateTime  int      `json:"create_time" sql:"create_time"`
	UpdateTime  int      `json:"update_time" sql:"update_time"`
	PublishTime int      `josn:"publish_time" sql:"publish_time"`
	Status      Status   `json:"status" sql:"status"`
}

// Model -
type Model struct{}

// Find -
func (m *Model) Find() {}

// FindOne -
func (m *Model) FindOne() {}

// Create -
func (m *Model) Create(article *Article) {}

// Update -
func (m *Model) Update() {}

// Delete -
func (m *Model) Delete() {}
