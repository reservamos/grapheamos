package sql

import (
	"html"
	"net/url"
	"strconv"

	"github.com/k3a/html2text"
	"github.com/reservamos/search/internal/config"
)

// CmsSiteContent built to get background and text
type CmsSiteContent struct {
	ID      int    `dl:"id"`
	Slug    string `dl:"slug"`
	Bg      string `dl:"bg"`
	Content string `dl:"content"`
}

// BgURL background url
func (s *CmsSiteContent) BgURL() *string {
	if s.Bg != "" {
		result := config.App.ImageRepo + "/uploads/cms/site_content/bg/" + strconv.Itoa(s.ID) + "/" + url.QueryEscape(s.Bg)
		return &result
	}
	return nil
}

// ContentText html
func (s *CmsSiteContent) ContentText() *string {
	if s.Content != "" {
		result := html.UnescapeString(html2text.HTML2Text(s.Content))
		return &result
	}
	return nil
}
