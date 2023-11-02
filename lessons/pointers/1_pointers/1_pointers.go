package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	fmt.Println("v.X inside Scale function: ", v.X)
	v.Y = v.Y * f
	fmt.Println("v.Y inside Scale function: ", v.Y)
}

func (v Vertex) ScaleWithoutPointerReceiver(f float64) {
	v.X = v.X * f
	fmt.Println("v.X inside ScaleWithoutPointerReceiver function: ", v.X)
	v.Y = v.Y * f
	fmt.Println("v.Y inside ScaleWithoutPointerReceiver function: ", v.Y)
}

func main() {
	v := Vertex{3, 4}
	//TODO - do we need to use & in this case?
	//v := &Vertex{3, 4}

	fmt.Println("v.X inside Main before Scale function: ", v.X)
	fmt.Println("v.Y inside Main before Scale function: ", v.Y)

	v.Scale(10)

	fmt.Println("v.X inside Main after Scale function: ", v.X) // RESULT will be 30, SO IT STRUCT FIELD WAS CHANGED THANKS TO A POINTER
	fmt.Println("v.Y inside Main after Scale function: ", v.Y) // RESULT will be 40, SO IT STRUCT FIELD WAS CHANGED THANKS TO A POINTER

	v.ScaleWithoutPointerReceiver(10)
	fmt.Println("v.X inside Main after ScaleWithoutPointerReceiver function: ", v.X) // RESULT will be 3, SO IT STRUCT FIELD WAS NOT CHANGED
	fmt.Println("v.Y inside Main after ScaleWithoutPointerReceiver function: ", v.Y) // RESULT will be 4, SO IT STRUCT FIELD WAS NOT CHANGED

	//fmt.Println(v.Abs())
}
