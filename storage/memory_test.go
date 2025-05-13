package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorageCRUD(t *testing.T) {
	Reset() // reset global state

	u := CreateUser("Bob", "bob@example.com")
	assert.Equal(t, 1, u.ID)

	fetched, found := GetUser(u.ID)
	assert.True(t, found)
	assert.Equal(t, "Bob", fetched.Name)

	updated, ok := UpdateUser(u.ID, "Bobby", "bobby@example.com")
	assert.True(t, ok)
	assert.Equal(t, "Bobby", updated.Name)

	ok = DeleteUser(u.ID)
	assert.True(t, ok)

	_, found = GetUser(u.ID)
	assert.False(t, found)
}
