package tag

import "github.com/depsea/dep-core/core"

// ID -
type ID = int

// Status -
type Status = int

const (
	// StatusUnable -
	StatusUnable Status = iota

	// StatusAble -
	StatusAble
)

// Tag -
type Tag struct {
	ID         ID     `json:"id" sql:"id"`
	Name       string `json:"name" sql:"name"`
	Desc       string `json:"desc" sql:"description"`
	CreateTime int    `json:"craete_time" sql:"create_time"`
	UpdateTime int    `json:"update_time" sql:"update_time"`
	Status     Status `json:"status" sql:"status"`
}

// Model -
type Model struct{}

var modelInstance *Model

// NewModel -
func NewModel() *Model {
	if modelInstance == nil {
		modelInstance = &Model{}
	}
	return modelInstance
}

// Find -
func (m *Model) Find() {

}

// FindByID -
func (m *Model) FindByID(id ID) {

}

// Create -
func (m *Model) Create(tag *Tag) {
	now := core.GetNowTimestamp()
	tag.CreateTime = now
	tag.UpdateTime = now
	tag.Status = StatusAble
}

// Update -
func (m *Model) Update(id ID, tag Tag) {

}

// Delete -
func (m *Model) Delete(id ID) {
	tag := &Tag{
		ID:         id,
		UpdateTime: core.GetNowTimestamp(),
		Status:     StatusUnable,
	}

	print(tag)
}
