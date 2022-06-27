package spider

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// 使用Go官方包http begin
// 抓取页面
func FetchByHttp(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	req.Header.Set("Cookie", "_ga=GA1.2.224942325.1626580594; _gid=GA1.2.1333806371.1656230940; __atuvc=4|26; __atuvs=62b8141c09e9f01a003")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("http status code:", resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err:", err)
		return ""
	}
	return string(body)
}

func Parse(html string) {
	// 替换空格
	html = strings.Replace(html, "\n", "", -1)

	// 边栏侧内容正则
	regex_sidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	// 找到边栏内容块
	sidebar := regex_sidebar.FindString(html)
	// 链接正则
	regex_link := regexp.MustCompile(`href="(.*?)"`)
	//找到所有链接
	links := regex_link.FindAllString(sidebar, -1)

	base_url := "https://gorm.io/zh_CN/docs/"
	for i, v := range links {
		if i == (len(links)-1) || i == (len(links)-2) || i == (len(links)-3) {
			continue
		}
		s := v[6 : len(v)-1]
		url := base_url + s
		body := FetchByHttp(url)
		go ParseContent(body)
	}
}

// end

func ParseContent(body string) {
	// 替换空格
	body = strings.Replace(body, "\n", "", -1)

	// 页面内容正则
	regex_content := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	// 找到页面内容
	content := regex_content.FindString(body)

	// 标题正则
	regex_title := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
	//找到标题内容
	title := regex_title.FindString(content)
	fmt.Println("title:", title)

	//找到具体
	title = title[42 : len(title)-5]
	fmt.Println("title:", title)
	//do save
}

// 基于goqery实现 begin
func GoQuery() {
	url := "https://gorm.io/zh_CN/docs/"
	d, _ := goquery.NewDocument(url)
	//找到sidebar-link级别
	d.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
		// 获取sidebar-link级别的属性
		href, _ := s.Attr("href")
		fmt.Println("href", href)
	})
}

/**
选择器：常用的有元素名称选择器、ID选择器、class选择器
元素名称：d.Find("div")
ID选择器：d.Find("#div")
class选择器：d.Find(".div")
*/

// end

//基于colly实现 begin
func Colly() {
	c := colly.NewCollector()
	// goquery class selector
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url", r.URL)
	})
	c.Visit("https://gorm.io/zh_CN/docs/")
}

func Spider(ctx context.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.baidu.com"),
		colly.MaxDepth(1),
	)
	var visited = map[string]bool{}
	detailRegex, _ := regexp.Compile(`\?p=\d+$`)
	listRegex, _ := regexp.Compile(`/t/\d+#\w+`)
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if visited[link] && (detailRegex.Match([]byte(link)) || listRegex.Match([]byte(link))) {
			return
		}
		// 那么不是我们关心的内容，要跳过
		if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
			println("not match", link)
			return
		}
		// 所以爬虫逻辑中应该有 sleep 逻辑以避免被封杀
		time.Sleep(time.Second)
		println("match", link)

		visited[link] = true
		time.Sleep(time.Millisecond * 2)
		collector.Visit(e.Request.AbsoluteURL(link))
	})
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	err := collector.Visit("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
}

//end
