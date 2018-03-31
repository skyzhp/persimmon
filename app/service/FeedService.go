package service

import (
	"time"
	"html/template"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/utils"
)

type FeedService struct{}

func (this *FeedService) BuildFeed(domain string) (string, error) {
	allOption, optErr := optionService.GetAllOption(false)
	if optErr != nil {
		return "", optErr
	}
	optArr := make(map[string]string)
	for _, val := range allOption {
		optArr[val.OptionName] = val.OptionValue
	}
	optArr["domain"] = domain

	posts, pErr := postService.SearchList(0, "", 15, 1, false)
	if pErr != nil {
		return "", pErr
	}

	return this.makeXml(optArr, posts), nil
}

func (this *FeedService) makeXml(optArr map[string]string, posts []info.Posts) string {
	lastBuildDate := time.Now().Format("Mon, 2 Jan 2006 15:04:05")
	begin := `<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0">
<channel>
<atom:link href="https://` + optArr["domain"] + `/feed" rel="self" type="application/rss+xml" />`

	siteInfo := `<title>` + optArr["site_name"] + `</title>
        <link>http://` + optArr["domain"] + `</link>
        <description>` + optArr["description"] + `</description>
        <language>zh-CN</language>
        <copyright>Copyright (c) ` + optArr["site_name"] + `</copyright>
        <lastBuildDate>` + lastBuildDate + `</lastBuildDate>`

	items := ""
	for _, val := range posts {
		pubDate := utils.Date(utils.WdmyTimeFormat, val.CreatedAt)
		items += `<item>
            <title>` + val.Title + `</title>
            <link>http://` + optArr["domain"] + "/post/" + val.Flag + `</link>
            <description><![CDATA[` + template.HTMLEscapeString(val.Content) + `]]></description>
            <guid>` + val.Flag + `</guid>
            <pubDate>` + pubDate + ` +0000</pubDate>
        </item>`
	}

	end := `</channel></rss>`
	html := begin + siteInfo + items + end
	return html
}
