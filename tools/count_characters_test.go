package tools

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCountCharacters(t *testing.T) {
	t.Run("纯中文", func(t *testing.T) {
		lens := CountCharacters("天街小雨润如酥")
		assert.Equal(t, lens, 7)
	})
	t.Run("中英文", func(t *testing.T) {
		lens := CountCharacters("天街小雨润如酥werty")
		assert.Equal(t, lens, 12)
	})
}
