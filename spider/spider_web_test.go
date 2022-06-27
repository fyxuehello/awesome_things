package spider

import (
	"testing"
)

func TestSpider(t *testing.T) {
	t.Run("http", func(t *testing.T) {
		url := "https://gorm.io/zh_CN/docs/"
		html := FetchByHttp(url)
		Parse(html)
	})
	t.Run("colly", func(t *testing.T) {
		Colly()
	})
}
