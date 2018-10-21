package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Card struct {
	CardName string
	CardUrl  string
}

func findInAttributes(n []html.Attribute, t string) bool {
	for _, v := range n {
		if t == v.Val {
			return true
		}
	}
	return false
}

func parseCardNode() {

}

func main() {
	var cards []Card

	s := FirstPageHtml()

	// webPage := `https://www.cardkingdom.com/catalog/view?filter%5Bipp%5D=60&filter%5Bsort%5D=most_popular&filter%5Bsearch%5D=mtg_advanced&filter%5Bcategory_id%5D=0&filter%5Bmulti%5D%5B0%5D=1&filter%5Btype_mode%5D=any&filter%5Bmanaprod_select%5D=any`
	// res, err := http.Get(webPage)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// doc, err := html.Parse(res.Body)
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		// fmt.Println("Type: ", n.Type, "dataAtom: ", n.DataAtom, " Data:", n.Data, "Attr:", n.Attr)
		if n.Data == "span" && findInAttributes(n.Attr, "productDetailTitle") {
			//fmt.Println(n.Data, n.Attr)
			var card Card
			c := n.FirstChild
			for _, attr := range c.Attr {
				// fmt.Print(attr.Val, " ")
				card.CardUrl = attr.Val
			}
			c = c.FirstChild
			card.CardName = c.Data
			cards = append(cards, card)
			// fmt.Println()
			for c = c.NextSibling; c != nil; c = c.NextSibling {
				f(c)
			}
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	for i, v := range cards {
		fmt.Println(i, v)
	}

}
