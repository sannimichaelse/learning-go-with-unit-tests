package shapes

import "fmt"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64{
	return rectangle.Width * rectangle.Height
}

func main(){
	 fmt.Printf("Hello there")
}