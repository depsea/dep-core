package role

import (
	"time"

	"github.com/depsea/api/core"
	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
)

// ID -
type ID = int

// Role -
type Role struct {
	ID         ID     `json:"id" sql:"id"`
	Name       string `json:"name" sql:"name"`
	Desc       string `json:"desc" sql:"description"`
	CreateTime int    `json:"create_time" sql:"create_time"`
	UpdateTime int    `json:"update_time" sql:"update_time"`
}

// BeforeInsert -
func (r *Role) BeforeInsert(db orm.DB) error {
	now := int(time.Now().Unix())
	r.CreateTime = now
	r.UpdateTime = now
	return nil
}

// BeforeUpdate -
func (r *Role) BeforeUpdate(db orm.DB) error {
	now := int(time.Now().Unix())
	r.UpdateTime = now
	return nil
}

// Model -
type Model struct {
	DB *pg.DB
}

// NewModel -
func NewModel() *Model {
	db := core.DB
	return &Model{
		DB: db,
	}
}

// Find -
func (m *Model) Find() ([]Role, error) {
	roles := []Role{}

	err := m.DB.Model(&roles).Select()

	return roles, err
}

// FindOne -
func (m *Model) FindOne(roleID ID) (Role, error) {
	role := &Role{ID: roleID}
	err := m.DB.Select(role)

	return *role, err
}

// Create -
func (m *Model) Create(role *Role) error {
	return m.DB.Insert(role)
}

// Update -
func (m *Model) Update(roleID ID, role *Role) error {
	return nil
}

// Delete -
func (m *Model) Delete(roleID ID) (Role, error) {
	return Role{}, nil
}
