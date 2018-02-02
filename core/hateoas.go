package core

import (
	"encoding/json"
	"path"
)

// Link hatoas返回结构
type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
	Title  string `json:"title"`
}

//Slash hatoas添加href
func (l *Link) Slash(s string) *Link {
	l.Href = path.Join(l.Href, s)
	return l
}

//WithRel hatoas添加hrel
func (l *Link) WithRel(s string) *Link {
	l.Rel = s
	return l
}

//WithMethod hatoas添加method
func (l *Link) WithMethod(s string) *Link {
	l.Method = s
	return l
}

//WithTitle hatoas添加title
func (l *Link) WithTitle(s string) *Link {
	l.Title = s
	return l
}

//Links hatoas集合
type Links struct {
	links []*Link
}

//Add hatoas集合添加link
func (l *Links) Add(link *Link) *Links {
	if link != nil {
		l.links = append(l.links, link)
	}
	return l
}

//MarshalJSON hatoas集合json格式化
func (l *Links) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.links)
}

//Hateoas 最终返回结果
type Hateoas struct {
	Links Links `json:"_links"`
}

//AddLinks 添加link集合
func (h *Hateoas) AddLinks(l Links) *Hateoas {
	h.Links = l
	return h
}

//HateoasTemplate hateoas模版
type HateoasTemplate struct {
	Links Links `json:"_link_template"`
}

//AddLinks 为模版添加link集合
func (h *HateoasTemplate) AddLinks(l Links) *HateoasTemplate {
	h.Links = l
	return h
}

//LinkTo 创建 link
func LinkTo(href string, rel string, method string, title string) *Link {
	if href != "" {
		return &Link{Rel: rel, Href: href, Method: method, Title: title}
	}
	return &Link{Rel: "self", Href: "/"}
}
