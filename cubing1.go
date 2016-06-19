package main

import (
	"fmt"

	"github.com/gitdlam/rtreego"
)

type ItemType struct {
	ID   int64
	Point rtreego.Point
	b1   float64
	b2   float64
	b3   float64
	Box  rtreego.NewRect
	BBbox rtreego.NewRect
}

func (item *ItemType) Bounds() *rtreego.Rect {
	item.BBox.
	return &item.BBox
}

func (item *ItemType) Setup(dim1 float64, dim2 float64, dim3 float64) {
 item.Point = rtreego.Point{0,0,0}
 item.Box = rtreego.NewRect(item.Point, [3]float64{dim1,dim2,dim3})
 item.BBox = rtreego.NewRect(item.Point, [3]float64{dim1+0.0001,dim2+0.0001,dim3+0.0001})
}

func main() {
	rt := rtreego.NewTree(25, 50)
	b1 := Item{}
	b1.Setup(rtreego.Point{5,10,20}
	
	fmt.Println(rt, b1)
}
