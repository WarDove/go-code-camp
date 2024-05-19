package main

import "fmt"

func intOrStringSum(a, b any) any {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aString, aIsInt := a.(string)
	bString, bIsInt := b.(string)
	if aIsInt && bIsInt {
		return aString + bString
	}
	return nil
}

// just to compare what generic saves us from doing ^^
func intOrStringSumGeneric[T int | string](a, b T) T {
	return a + b
}

func main() {

	// typeSwitch example
	var myVar interface{}

	switch myVar.(type) {
	case int:
		println("myVar is an int")
	case string:
		println("myVar is a string")
	case bool:
		println("myVar is a bool")
	default:
		println("myVar is another type")
	}

	myVarInt, ok := myVar.(int)
	if ok {
		fmt.Println("myVar is an int:", myVarInt)
	} else {
		fmt.Println("myVar is not an int")
	}

}
