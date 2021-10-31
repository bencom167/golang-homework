package test

import (
	"gin-postgres/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockData(t *testing.T) {
	err := repo.MockData()
	assert.Nil(t, err)
}
