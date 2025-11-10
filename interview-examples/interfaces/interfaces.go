package main

import (
	"fmt"
	"math"
)

// 1. Basic Interface Definition
// Interfaces define behavior through method signatures
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Implicit Interface Implementation
// Go interfaces are satisfied implicitly - no "implements" keyword needed
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 3. Using interfaces for polymorphism
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// 4. Empty interface (interface{} or any)
func DescribeValue(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

// 5. Type assertion
func GetAreaIfShape(i interface{}) {
	// Type assertion with "comma, ok" idiom
	if shape, ok := i.(Shape); ok {
		fmt.Printf("It's a shape! Area: %.2f\n", shape.Area())
	} else {
		fmt.Println("Not a shape")
	}
}

// 6. Type switch
func ClassifyType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case Shape:
		fmt.Printf("Shape with area: %.2f\n", v.Area())
	case nil:
		fmt.Println("Nil value")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// 7. Interface embedding/composition
type Named interface {
	Name() string
}

type NamedShape interface {
	Shape // Embedded interface
	Named // Embedded interface
}

type LabeledRectangle struct {
	Rectangle
	Label string
}

func (lr LabeledRectangle) Name() string {
	return lr.Label
}

// 8. Nil interface vs nil value
type Printer interface {
	Print()
}

type MyPrinter struct {
	Message string
}

func (mp *MyPrinter) Print() {
	if mp == nil {
		fmt.Println("Nil pointer printer")
		return
	}
	fmt.Println(mp.Message)
}

func main() {
	fmt.Println("=== 1. Basic Interface Implementation ===")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	PrintShapeInfo(rect)
	PrintShapeInfo(circle)

	fmt.Println("\n=== 2. Interfaces in Slice (Polymorphism) ===")
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Rectangle{Width: 2, Height: 8},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		PrintShapeInfo(shape)
	}

	fmt.Println("\n=== 3. Empty Interface ===")
	DescribeValue(42)
	DescribeValue("hello")
	DescribeValue(rect)
	DescribeValue(true)

	fmt.Println("\n=== 4. Type Assertion ===")
	var emptyInterface interface{} = circle
	GetAreaIfShape(emptyInterface)
	GetAreaIfShape("not a shape")

	// Unsafe type assertion (panics if wrong type)
	// circleValue := emptyInterface.(Circle)
	// fmt.Printf("Circle radius: %.2f\n", circleValue.Radius)

	// Safe type assertion
	if circleValue, ok := emptyInterface.(Circle); ok {
		fmt.Printf("Circle radius: %.2f\n", circleValue.Radius)
	}

	fmt.Println("\n=== 5. Type Switch ===")
	ClassifyType(42)
	ClassifyType("Go")
	ClassifyType(rect)
	ClassifyType(nil)
	ClassifyType(3.14)

	fmt.Println("\n=== 6. Interface Embedding ===")
	labeledRect := LabeledRectangle{
		Rectangle: Rectangle{Width: 6, Height: 8},
		Label:     "MyRectangle",
	}
	var namedShape NamedShape = labeledRect
	fmt.Printf("Name: %s, Area: %.2f, Perimeter: %.2f\n",
		namedShape.Name(), namedShape.Area(), namedShape.Perimeter())

	fmt.Println("\n=== 7. Nil Interface vs Nil Value ===")

	// Case 1: Nil interface
	var p1 Printer
	fmt.Printf("p1 == nil: %v\n", p1 == nil)
	// p1.Print() // This would panic - cannot call method on nil interface

	// Case 2: Non-nil interface with nil value
	var p2 Printer = (*MyPrinter)(nil)
	fmt.Printf("p2 == nil: %v\n", p2 == nil) // false! Interface is not nil
	p2.Print()                               // Works! Method can handle nil receiver

	// Case 3: Non-nil interface with non-nil value
	var p3 Printer = &MyPrinter{Message: "Hello"}
	fmt.Printf("p3 == nil: %v\n", p3 == nil)
	p3.Print()

	fmt.Println("\n=== 8. Interface Conversion ===")
	var i interface{} = Rectangle{Width: 5, Height: 10}

	// Convert empty interface to specific interface
	if s, ok := i.(Shape); ok {
		fmt.Printf("Converted to Shape: Area = %.2f\n", s.Area())
	}

	fmt.Println("\n=== 9. Zero Value Interface ===")
	var zeroShape Shape
	fmt.Printf("Zero value interface == nil: %v\n", zeroShape == nil)
	// zeroShape.Area() // Would panic - calling method on nil interface
}
