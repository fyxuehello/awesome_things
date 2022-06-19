package tools

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestGenCsvFile(t *testing.T) {
	t.Run("生成csv文件", func(t *testing.T) {
		err := GenCsvFile()
		assert.Equal(t, err, nil)
	})

	t.Run("解析csv文件", func(t *testing.T) {
		err := ParseCsvFile()
		assert.Equal(t, err, nil)
	})
}
