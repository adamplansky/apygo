package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"golang.org/x/net/html"
)

type Card struct {
	Id         string
	CardName   string
	CardUrl    string
	Price      string
	Qty        string
	Style      string
	Rarity     string
	Collection string
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
func removeWhiteSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}
func isfindInAttributes(n []html.Attribute, t string) bool {
	for _, v := range n {
		clsVals := strings.SplitAfter(v.Val, " ")
		for _, vv := range clsVals {
			if t == removeWhiteSpace(vv) {
				return true
			}
		}
	}
	return false
}

func parseProductDetail(n *html.Node) (title string, href string) {
	var t, h string
	c := n.FirstChild
	if c != nil && len(c.Attr) > 0 {
		h = c.Attr[0].Val
	}
	x := c.FirstChild
	if x != nil {
		t = x.Data
	}
	return t, h
}
func getValue(n *html.Node) string {
	var val string
	if n != nil {
		for _, v := range n.Attr {
			if v.Key == "value" {
				val = v.Val
			}
		}
	}
	return val
}
func getText(n *html.Node) string {
	var val string
	if n != nil {
		x := n.FirstChild
		if x != nil {
			val = x.Data
		}
	}
	return val
}
func parseStyle(c *html.Node) (string, string, string, string) {
	var id, qty, price, cardType string
	_cardType := traverse(c, "style")
	_productID := traverse(c, "product_id")
	_qty := traverse(c, "styleQty")
	_price := traverse(c, "stylePrice")

	price = getText(_price)
	qty = getText(_qty)
	cardType = getValue(_cardType)
	id = getValue(_productID)
	return id, qty, price, cardType

}

// func traverseCardTypes()
func parseCardNode(c *html.Node) {
	title, href := parseProductDetail(traverse(c, "productDetailTitle"))
	pds := getText(traverse(c, "productDetailSet"))
	pds = strings.TrimSpace(pds)
	collection := pds[0 : len(pds)-3]
	// fmt.Println(title, href)
	// fmt.Println(collection)
	rarity := pds[len(pds)-3:][1:2]
	// fmt.Println("rarity", rarity)

	cardTypes := traverse(c, "addToCartByType")
	li := cardTypes.FirstChild
	for c := li; c != nil; c = c.NextSibling {
		id, qty, price, cardType := parseStyle(c)
		if id == "" || qty == "" {
			continue
		}
		c := Card{
			Id:         id,
			CardName:   title,
			CardUrl:    href,
			Price:      price,
			Qty:        qty,
			Style:      cardType,
			Rarity:     rarity,
			Collection: collection,
		}
		cards = append(cards, c)
	}

	// x := c.NextSibling
	// fmt.Println(x, x.Data, x.DataAtom)
	// var card Card
	//
	// for _, attr := range c.Attr {
	// 	// fmt.Print(attr.Val, " ")
	// 	card.CardUrl = attr.Val
	// }
	// c = c.FirstChild
	// card.CardName = c.Data
	// cards = append(cards, card)
	// _ = findClass(c, "amdAndPrice")
	// fmt.Println(x)
}
func traverse(n *html.Node, cls string) *html.Node {
	if isfindInAttributes(n.Attr, cls) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, cls)
		if result != nil {
			return result
		}
	}
	return nil
}
func parsePage(n *html.Node) {
	if n.Data == "div" && (isfindInAttributes(n.Attr, "productItemWrapper") == true) {
		cardDetail := traverse(n, "itemContentWrapper")
		parseCardNode(cardDetail)
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parsePage(c)
	}
}

var cards []Card

// func (c Card) String() string {
// 	return
// }

func parallelRequest(url string, page string) {
	runtime.Gosched()
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	parsePage(doc)
	// for i, card := range cards {
	// 	fmt.Printf("%v %+v\n", i, card)
	// }
	fmt.Println("cards imported: ", len(cards), page)
	fmt.Println("time: ", time.Since(start))
	// wg.Done()

}

var wg sync.WaitGroup
var start time.Time

func main() {
	start = time.Now()
	// s := FirstPageHtml()
	// doc, err := html.Parse(strings.NewReader(s))
	// const gs int = 1
	// wg.Add(gs)
	fmt.Println("start importing Card Kingdom")

	semaphore := make(chan struct{}, 700)
	for i := 1; i <= 625; i++ {
		wg.Add(1)
		page := strconv.Itoa(i)
		webPage := "https://www.cardkingdom.com/catalog/view?filter%5Bipp%5D=60&filter%5Bsort%5D=most_popular&filter%5Bsearch%5D=mtg_advanced&filter%5Bcategory_id%5D=0&filter%5Bmulti%5D%5B0%5D=1&filter%5Btype_mode%5D=any&filter%5Bmanaprod_select%5D=any&page="
		url := webPage + page
		go func(url, page string) {
			defer wg.Done()

			semaphore <- struct{}{} // Lock
			defer func() {
				<-semaphore // Unlock
			}()

			// Resource intensive work goes here
			parallelRequest(url, page)
		}(url, page)

	}

	wg.Wait()
	// for i, card := range cards {
	// 	fmt.Printf("%v %+v\n", i, card)
	// }
	fmt.Println("cards imported: ", len(cards))
	fmt.Println("time: ", time.Since(start))

}
