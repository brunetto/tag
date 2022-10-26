package ecomm

type Cart struct{ items []Item }
type Item struct{}

func NewEmptyCart() Cart {
	return Cart{}
}

func NewCartWithItems(items ...Item) Cart {
	return Cart{items: items}
}
