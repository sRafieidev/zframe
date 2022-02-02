package libtest

import "fmt"

func TestPointerStructingo() {

	fmt.Println("========================================================================================")
	fmt.Println("Pointer Management In Golang Test")
	// sample for pointer management in Go
	// of integer type
	var V = 100
	// taking a pointer
	// of integer type
	var pt1 = &V
	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2
	var pt2 = &pt1
	fmt.Println("The Value of Variable V is = ", V)
	fmt.Println("Address of variable V is = ", &V)
	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)
	fmt.Println("The value of pt2 is = ", pt2)
	// Dereferencing the
	// pointer to pointer
	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)
	// double pointer will give the value of variable V
	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
	fmt.Println("========================================================================================")
}

func Testshift() {
	var t, i uint
	t, i = 1, 1

	for i = 1; i < 10; i++ {
		fmt.Printf("%d << %d = %d \n", t, i, t<<i)
	}

	fmt.Println()

	t = 512
	for i = 1; i < 10; i++ {
		fmt.Printf("%d >> %d = %d \n", t, i, t>>i)
	}

}
