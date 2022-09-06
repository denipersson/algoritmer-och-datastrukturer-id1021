package main

type Item_type int

const (
	ADD Item_type = iota
	SUB
	MUL
	DIV
	VALUE
	MOD
)

type Item struct {
	item_type Item_type
	value     int
}

func New_Item(item_type Item_type, value int) Item {
	return Item{item_type, value}
}

func Get_Type(i Item) Item_type {
	return i.item_type
}
func Get_Value(i Item) int {
	return i.value
}
