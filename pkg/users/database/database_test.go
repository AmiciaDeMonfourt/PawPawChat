package db_test

import (
	db "pawpawchat/pkg/users/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB_New(t *testing.T) {
	_, err := db.New()
	assert.NoError(t, err)
}
