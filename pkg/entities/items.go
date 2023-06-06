package entities

type Item struct{}

var ItemIndexes map[string]int

func (i *Item) Init() {
	ItemIndexes["Arrow"] = 0
	ItemIndexes["FireArrow"] = 1
	ItemIndexes["Heart"] = 2
	ItemIndexes["HalfHeart"] = 3
	ItemIndexes["Meat"] = 4
	ItemIndexes["Shield"] = 5
	ItemIndexes["Spear"] = 6
	ItemIndexes["Longbow"] = 7
	ItemIndexes["Longbow2"] = 8
	ItemIndexes["Longbow3"] = 9
	ItemIndexes["Dagger"] = 10
	ItemIndexes["Torch"] = 11
}
