package basic

import "fmt"

func PointersUsage() {
	num := 3
	//num = square(num)
	fmt.Println(num)

	squarePointer(&num)
	fmt.Println(num)

	//empty value flag
	var wallet *int
	fmt.Println(hasWallet(wallet))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))

}

/*func square(num int) int {
	return num * num
}*/

func squarePointer(num *int) {
	*num *= *num
	//the same
	//*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}
