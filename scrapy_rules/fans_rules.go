package scrapy_rules

import (
	"WeiboSpiderGo/mdb"
	"WeiboSpiderGo/utils"
	"fmt"
	"github.com/gocolly/colly"
	"gopkg.in/mgo.v2"
	"strconv"
	"strings"
	"time"
)

func SetFansCallback(getFansC *colly.Collector) {
	getFansC.OnResponse(func(r *colly.Response) {
		content := string(r.Body)
		uid := utils.ReParse(`(\d+)/fans`, r.Request.URL.String())
		if strings.Contains(r.Request.URL.String(), "page=1") {
			allPage := utils.ReParse(`/> 1/(\d+)page</div>`, content)
			pageNum, _ := strconv.Atoi(allPage)
			for i := 2; i < (pageNum + 1); i++ {
				link := fmt.Sprintf("%s/%s/fans?page=%d", BaseUrl, uid, i)
				getFansC.Visit(link)
			}
		}
	})
	getFansC.OnXML(`//a[text()="Follow him" or text()="Follow her" or text()="Remove"]/@href`, func(element *colly.XMLElement) {
		followUrl := element.Text
		uid := utils.ReParse(`uid=(\d+)`, followUrl)
		ID := utils.ReParse(`(\d+)/fans`, element.Request.URL.String())
		relationship := mdb.Relationships{}
		relationship.CrawlTime = int32(time.Now().Unix())
		relationship.FanId = uid
		relationship.FollowedId = ID
		relationship.Id_ = uid + "-" + ID
		err := mdb.Insert(dbName, "Relationships", relationship)
		if mgo.IsDup(err) {
			//There is duplicate data
			fmt.Println("already scrapy")
		}
	})
}

func GetFansUrl(uid string) string {
	return fmt.Sprintf("%s/%s/fans?page=1", BaseUrl, uid)
}
