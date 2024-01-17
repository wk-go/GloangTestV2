package main

//在 Go 语言中，解析 XML 是可以处理不成对标签的，比如 `<atom:link href="http://xxx" rel="self" type="application/rss+xml"/>`。这种标签被视为自闭合标签，是 XML 标准的一部分。在 Go 中，可以使用 `encoding/xml` 包来解析这样的 XML 数据。
//在这个例子中，定义了一个 `Link` 结构体来映射 XML 标签。然后使用 `xml.Unmarshal` 函数解析给定的 XML 数据。这个程序将打印出解析后的 `Link` 结构体的内容，包含 `href`, `rel`, 和 `type` 属性的值。这种方式适用于处理 XML 中的自闭合标签。
//这里有一个简单的例子展示如何解析这样的标签：

import (
	"encoding/xml"
	"fmt"
)

type Link struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

func main() {
	data := `<atom:link href="http://xxx" rel="self" type="application/rss+xml"/>`

	var link Link
	err := xml.Unmarshal([]byte(data), &link)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Link: %+v\n", link)
}
