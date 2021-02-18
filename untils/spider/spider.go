package spider

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/panjf2000/ants/v2"
	"iwara/models"
	"iwara/untils"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const Host = "ecchi.iwara.tv"
const attempts int = 3
const sleep time.Duration = 1 * time.Second

var wg sync.WaitGroup
var total int

type Spider struct {
}

func (s *Spider) Start() {
	go Start()
}

func Start() {
	// todo: 爬虫逻辑
	list()
}

func list() {
	pool, _ := ants.NewPool(10)
	defer ants.Release()
	total := Total()
	c := NewCollector()
	c.OnHTML("div.node.node-video.node-teaser.node-teaser", func(e *colly.HTMLElement) {
		// id, _ := strconv.ParseUint(e.Attr("id")[5:], 10, 32)
		viewStr := e.ChildText("div.left-icon.likes-icon")
		var view int
		if strings.Contains(viewStr, "k") {
			viewFloat, _ := strconv.ParseFloat(viewStr[:len(viewStr)-1], 2)
			view = int(viewFloat * 1000)
		} else {
			view, _ = strconv.Atoi(viewStr)
		}
		star, _ := strconv.Atoi(e.ChildText("div.right-icon.likes-icon"))
		video := &models.Video{
			Title:  e.ChildText("h3.title a"),
			Url:    fmt.Sprintf("https://ecchi.iwara.tv%s", e.ChildAttr("h3.title a", "href")),
			Poster: fmt.Sprintf("https:%s", e.ChildAttr("div.field-item.even img", "src")),
			View:   view,
			Star:   star,
		}
		log.Println(video.Title)
	})
	for i := 1; i <= total; i++ {
		c.UserAgent = RandomUserAgent()
		url := fmt.Sprintf("https://ecchi.iwara.tv/videos?sort=likes&page=%d", i)
		err := untils.Retry(attempts, sleep, func() error {
			wg.Add(1)
			return pool.Submit(func() {
				log.Printf("当前正在访问：%s", url)
				_ = c.Visit(url)
				wg.Done()
			})
		})
		if err != nil {
			log.Printf(err.Error())
		}
	}
	wg.Wait()
}

func Total() int {
	c := NewCollector()
	c.OnHTML("li.pager-last.last", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")
		log.Println(href)
		index := strings.Index(href, "page=")
		pageStr := href[index+5:]
		page, _ := strconv.Atoi(pageStr)
		total = page
	})
	err := c.Visit("https://ecchi.iwara.tv/videos")
	log.Println(err)

	log.Printf("一共搜索到%d页", total)

	if total ==0 {
		return Total()
	}

	return total
}

func Detail(id int) {
	err := untils.Retry(attempts, sleep, func() error {
		return nil
	})
	if err != nil {
		log.Printf(err.Error())
	}
}

func NewCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(Host),
		colly.UserAgent(RandomUserAgent()),
		// colly.Async(true),
	)
	_ = c.SetProxy("http://127.0.0.1:1087")
	return c
}
