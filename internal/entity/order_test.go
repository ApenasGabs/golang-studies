package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := &Order{}
	assert.Error(t, order.Validate(), "ID is required")
}
