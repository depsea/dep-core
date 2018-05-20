package user

import "github.com/depsea/dep-core/app/role"

// ID -
type ID int

// User -
type User struct {
	ID         ID      `json:"id" sql:"id"`
	Roles      role.ID `json:"roles" sql:"roles"`
	Username   string  `json:"username" sql:"username"`
	Password   string  `json:"password" sql:"password"`
	CreateTime int     `json:"create_time" sql:"create_time"`
	UpdateTime int     `json:"update_time" sql:"update_time"`
}

// Model -
type Model struct{}

// Find -
func (m *Model) Find() {}

// FindOne -
func (m *Model) FindOne() {}

// Create -
func (m *Model) Create() {}

// Update -
func (m *Model) Update() {}

// Delete -
func (m *Model) Delete() {}
