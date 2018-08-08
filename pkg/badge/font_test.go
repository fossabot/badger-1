package badge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeTextWidth(t *testing.T) {
	expectedLength := 153
	width := computeTextWidth("Lorem ipsum dolor sit amet")

	assert.Equal(t, width, expectedLength)
}

func TestComputeTextWidthWithSpecialChar(t *testing.T) {
	expectedLength := 60
	width := computeTextWidth("★★★★☆")

	assert.Equal(t, width, expectedLength)
}
