package role

import (
	"github.com/depsea/dep-core/core"
)

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

// Role -
type Role struct {
	TableName struct{} `sql:"role" json:"-"`

	ID         ID     `json:"id" sql:"id"`
	Name       string `json:"name" sql:"name"`
	Desc       string `json:"desc" sql:"description"`
	CreateTime int    `json:"create_time" sql:"create_time"`
	UpdateTime int    `json:"update_time" sql:"update_time"`
	Status     Status `json:"status" sql:"status"`
}

// Model -
type Model struct{}

// modelInstance -
var modelInstance *Model

// NewModel -
func NewModel() *Model {
	if modelInstance == nil {
		modelInstance = &Model{}
	}
	return modelInstance
}

// Find -
func (m *Model) Find() ([]Role, error) {
	var roles = []Role{}

	err := core.DB.Model(&roles).Select()

	return roles, err
}

// FindOne -
func (m *Model) FindOne(roleID ID) (Role, error) {
	role := &Role{ID: roleID}

	err := core.DB.Select(role)

	return *role, err
}

// Create -
func (m *Model) Create(role *Role) (Role, error) {
	now := core.GetNowTimestamp()
	role.CreateTime = now
	role.UpdateTime = now

	err := core.DB.Insert(role)

	return *role, err
}

// Update -
func (m *Model) Update(roleID ID, role *Role) (Role, error) {
	role.ID = roleID
	role.UpdateTime = core.GetNowTimestamp()

	err := core.DB.Update(role)

	return *role, err
}

// Delete -
func (m *Model) Delete(roleID ID) (Role, error) {
	role := &Role{
		ID:         roleID,
		Status:     StatusUnable,
		UpdateTime: core.GetNowTimestamp(),
	}

	_, err := core.DB.Model(role).Column("status", "update_time").Update()

	return *role, err
}
