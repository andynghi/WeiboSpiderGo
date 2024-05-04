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

func SetInfoCallback(getInfoC, getMoreInfoC *colly.Collector) {
	getInfoC.OnResponse(func(r *colly.Response) {
		content := string(r.Body)
		info := mdb.Information{}
		info.CrawlTime = int32(time.Now().Unix())
		info.Id_ = utils.ReParse(`(\d+)/info`, r.Request.URL.String())
		nickName := utils.ReParse(`nickname?[::]?(.*?)<br/>`, content)
		authentication := utils.ReParse(`Authentication?[::]?(.*?)<br/>`, content)
		gender := utils.ReParse(`Gender?[::]?(.*?)<br/>`, content)
		place := utils.ReParse(`Region?[::]?(.*?)<br/>`, content)
		briefIntroduction := utils.ReParse(`Introduction?[::]?(.*?)<br/>`, content)
		birthday := utils.ReParse(`Birthday?[::]?(.*?)<br/>`, content)
		sexOrientation := utils.ReParse(`Orientation?[::]?(.*?)<br/>`, content)
		sentiment := utils.ReParse(`Sentiment status?[::]?(.*?)<br/>`, content)
		vipLevel := utils.ReParse(`Member level?[::]?(.*?) <a`, content)
		labels := utils.ReParseMayLi(`stag=1">(.*?)</a>`, content) //Label
		info.Nickname = nickName
		info.Gender = gender
		placeli := strings.Split(place, " ")
		info.Province = placeli[0]
		if len(placeli) > 1 {
			info.City = placeli[1]
		}
		info.BriefIntroduction = briefIntroduction
		info.Birthday = birthday
		if sexOrientation == gender {
			info.SexOrientation = "Gay"
		} else {
			info.SexOrientation = "Heterosexual"
		}
		info.Sentiment = sentiment
		info.VipLevel = vipLevel
		info.Authentication = authentication
		info.Labels = ""
		for i, labelItem := range labels {
			if i != 0 {
				info.Labels += ","
			}
			info.Labels += labelItem[1]
		}
		r.Ctx.Put("info", info)
		getMoreInfoC.Request("GET", "https://weibo.cn/u/"+info.Id_, nil, r.Ctx, nil)
	})
}

func SetMoreInfoCallback(getMoreInfoC *colly.Collector) {
	getMoreInfoC.OnResponse(func(r *colly.Response) {
		content := string(r.Body)
		info := r.Ctx.GetAny("info").(mdb.Information)
		tweetsNum := utils.ReParse(`Weibo\[(\d+)\]`, content)
		followsNum := utils.ReParse(`follow\[(\d+)\]`, content)
		fansNum := utils.ReParse(`fans\[(\d+)\]`, content)
		if tweetsNum != "" {
			temp, _ := strconv.Atoi(tweetsNum)
			info.TweetsNum = int32(temp)
		}
		if followsNum != "" {
			temp, _ := strconv.Atoi(followsNum)
			info.FollowsNum = int32(temp)
		}
		if fansNum != "" {
			temp, _ := strconv.Atoi(fansNum)
			info.FansNum = int32(temp)
		}
		err := mdb.Insert(dbName, "Information", info)
		if mgo.IsDup(err) {
			//There is duplicate data
			fmt.Println("already scrapy")
		}
	})
}
