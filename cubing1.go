package main

import (
	"fmt"

	"github.com/gitdlam/rtreego"
)

type ItemType struct {
	ID           int64
	pointChanged bool
	dim1         float64
	dim2         float64
	dim3         float64
	BBox         rtreego.Rect
}

func (item *ItemType) Bounds() *rtreego.Rect {
	return &item.BBox
}

func (item *ItemType) Setup(x float64, y float64, z float64, dim1 float64, dim2 float64, dim3 float64) {
	item.dim1 = dim1
	item.dim2 = dim2
	item.dim3 = dim3
	item.BBox, _ = rtreego.NewRect(rtreego.Point{x, y, z}, [3]float64{dim1 + 0.0001, dim2 + 0.0001, dim3 + 0.0001})
}

func (item *ItemType) Move(x float64, y float64, z float64) *ItemType {
	var newItem ItemType
	newItem.Setup(x, y, z, item.dim1, item.dim2, item.dim3)
	return &newItem
}

func main() {
	rt := rtreego.NewTree(25, 50)
	b1 := ItemType{}
	b1.Setup(0, 0, 0, 5, 10, 20)

	fmt.Println(rt, b1)
}
