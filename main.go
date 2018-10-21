package main

import (
	"fmt"
	"net/http"

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

	// s := cardKingdom()

	webPage := `https://www.cardkingdom.com/catalog/view?filter%5Bipp%5D=60&filter%5Bsort%5D=most_popular&filter%5Bsearch%5D=mtg_advanced&filter%5Bcategory_id%5D=0&filter%5Bmulti%5D%5B0%5D=1&filter%5Btype_mode%5D=any&filter%5Bmanaprod_select%5D=any`
	res, err := http.Get(webPage)
	if err != nil {
		fmt.Println(err)
	}
	doc, err := html.Parse(res.Body)
	// doc, err := html.Parse(strings.NewReader(s))
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
	// webPage := `https://www.cardkingdom.com/catalog/view?filter%5Bipp%5D=60&filter%5Bsort%5D=most_popular&filter%5Bsearch%5D=mtg_advanced&filter%5Bcategory_id%5D=0&filter%5Bmulti%5D%5B0%5D=1&filter%5Btype_mode%5D=any&filter%5Bmanaprod_select%5D=any`
	// res, err := http.Get(webPage)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// z := html.NewTokenizer(res.Body)
	// for {
	// 	tt := z.Next()
	// 	token := z.Token()
	//
	// 	switch {
	// 	case tt == html.ErrorToken:
	// 		return
	// 	case tt == html.TextToken && strings.TrimSpace(token.String()) == "":
	// 		fmt.Println("text token: ", token, len(token.String()))
	// 		s := token.String()
	// 		var a [20]byte
	// 		copy(a[:], s)
	// 		fmt.Println("s:", []byte(s), "a:", a)
	// 	default:
	// 		fmt.Println(token)
	// 	}

}
func cardKingdom() string {
	return `<div class="itemContentWrapper">
  <table>
  <tbody><tr class="detailWrapper">
  <td>
  <span class="productDetailTitle"><a href="https://www.cardkingdom.com/mtg/core-set-2019/resplendent-angel">Resplendent Angel</a></span>
  <div class="productDetailSet">
  Core Set 2019 (M) <div>
  </div></div></td>
  <td class="productDetailDrillIn">
  <div class="productDetailCastCost"><img src="/media/images/web/mana_symbols/mana_1.png" style="padding-right:1px;"><img src="/media/images/web/mana_symbols/mana_w.png" style="padding-right:1px;"><img src="/media/images/web/mana_symbols/mana_w.png" style="padding-right:1px;"></div>
  <div class="productDetailType">
  3/3 Creature - Angel </div>
  </td>
  </tr>
  <tr class="detailFlavortext">
  <td colspan="2">
  Flying<br>
  At the beginning of each end step, if you gained 5 or more life this turn, create a 4/4 white Angel creature token with flying and vigilance.<br>
  <img style="margin:1px;" src="/media/images/web/mana_symbols/mana_3.png"><img style="margin:1px;" src="/media/images/web/mana_symbols/mana_w.png"><img style="margin:1px;" src="/media/images/web/mana_symbols/mana_w.png"><img style="margin:1px;" src="/media/images/web/mana_symbols/mana_w.png">: Until end of turn, Resplendent Angel gets +2/+2 and gains lifelink. </td>
  </tr>
  <tr>
  <td colspan="2">

  </td>
  </tr>
  </tbody></table>

  <style>
  .table-borderless tbody tr td, .table-borderless tbody tr th, .table-borderless thead tr th {
    border: none;
  }
  </style>
  <div class="addToCartWrapper hasQty">
  <ul class="cardTypeList">




  <li class="NM active">NM</li><li class="EX disabled">EX</li><li class="VG disabled">VG</li><li class="G disabled">G</li>
  </ul>
  <ul class="addToCartByType">
  <li class="itemAddToCart  NM active">
  <form action="/cart/add" method="get" class="addToCartForm">
  <input type="hidden" class="product_id" name="product_id[0]" value="219913">
  <input type="hidden" class="style" name="style[0]" value="NM">
  <input type="hidden" class="maxQty" name="maxQty" value="8">
  <div class="amtAndPrice"><span class="styleQty">8</span> <span class="styleQtyAvailText">available </span> <span class="styleAt">@</span> <span class="stylePrice"> $20.99
  </span>
  </div>
  <div class="dropdown">
  <button class="selectQty btn btn-default btn-lg dropdown-toggle col-xs-6" data-toggle="dropdown" aria-expanded="false" style="height:46px;">
  <span class="addToCartButton">Add to Cart<input type="text" readonly="" disabled="" class="qty incart" style="display:none;" name="qty[0]" size="2" value="0"></span> <span class="glyphicon glyphicon-triangle-bottom">
  </span></button>
  <ul class="dropdown-menu qtyList twoRow">
  <li><span class="glyphicon glyphicon-trash"></span></li>
  <li>1</li><li>2</li><li>3</li><li>4</li><li>5</li><li>6</li><li>7</li><li>8</li>
  </ul>
  </div>

  </form>
  </li>
  <li class="itemAddToCart  outOfStock  EX ">
  <form action="/cart/add" method="get" class="addToCartForm maxxed noInventory">
  <input type="hidden" class="product_id" name="product_id[1]" value="219913">
  <input type="hidden" class="style" name="style[1]" value="EX">
  <input type="hidden" class="maxQty" name="maxQty" value="0">
  <div class="amtAndPrice"><span class="styleQty">0</span> <span class="styleQtyAvailText">available </span> <span class="styleAt">@</span> <span class="stylePrice"> $17.84
  </span>
  </div>
  <div class="outOfStockNotice">Out of stock.</div>

  </form>
  </li>
  <li class="itemAddToCart  outOfStock  VG ">
  <form action="/cart/add" method="get" class="addToCartForm maxxed noInventory">
  <input type="hidden" class="product_id" name="product_id[2]" value="219913">
  <input type="hidden" class="style" name="style[2]" value="VG">
  <input type="hidden" class="maxQty" name="maxQty" value="0">
  <div class="amtAndPrice"><span class="styleQty">0</span> <span class="styleQtyAvailText">available </span> <span class="styleAt">@</span> <span class="stylePrice"> $14.69
  </span>
  </div>
  <div class="outOfStockNotice">Out of stock.</div>

  </form>
  </li>
  <li class="itemAddToCart  outOfStock  G ">
  <form action="/cart/add" method="get" class="addToCartForm maxxed noInventory">
  <input type="hidden" class="product_id" name="product_id[3]" value="219913">
  <input type="hidden" class="style" name="style[3]" value="G">
  <input type="hidden" class="maxQty" name="maxQty" value="0">
  <div class="amtAndPrice"><span class="styleQty">0</span> <span class="styleQtyAvailText">available </span> <span class="styleAt">@</span> <span class="stylePrice"> $10.49
  </span>
  </div>
  <div class="outOfStockNotice">Out of stock.</div>

  </form>
  </li>
  </ul>
  </div>
  <input type="hidden" name="redirect" value="1">
  <script type="text/javascript">

  </script>
  </div>`
}
