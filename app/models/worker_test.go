package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchWorker(t *testing.T) {
	w := FetchWorker(1)
	assert.Equal(t, uint(0), w.ID)
}
