package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func getHref(t html.Token) string {
	var href string
	for _, attr := range t.Attr {
		if attr.Key == "href" {
			href = attr.Val
			break
		}
	}

	return href
}

func getInnerText(z *html.Tokenizer) string {
	textNodes := make([]string, 1)

	for t := z.Token(); !(t.Type == html.EndTagToken && t.DataAtom == atom.A); t = z.Token() {
		if t.Type == html.TextToken {
			textNodes = append(textNodes, strings.TrimSpace(t.Data))
		}

		z.Next()
	}
	return strings.TrimSpace(strings.Join(textNodes, " "))
}

func Parse(r io.Reader) ([]Link, error) {
	links := make([]Link, 0)

	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		t := z.Token()
		if t.Type == html.StartTagToken && t.DataAtom == atom.A && len(t.Attr) > 0 {
			href := getHref(t)

			if href != "" {
				z.Next()

				text := getInnerText(z)
				links = append(links, Link{
					href,
					text,
				})
			}
		}
	}

	return links, nil
}
