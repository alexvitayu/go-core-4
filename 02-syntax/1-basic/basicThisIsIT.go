package basic

import "fmt"

//Pointers
// 1. Обозначение: *Type (*int, *string, *bool)
// 2. Разыменование указателей *somePointer
// 3. Default value любого указателя nil и nil pointer exception
// 4. Передача аргументов функций по значению и ссылке/указателю
// 5. Где использовать указатели?
// а) сайд эффект
// б) признак пустого значения
// в) экономия памяти (чтобы не копировать значения)

func Pointers2() {
	// default value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	//получение not-nil указателей
	//variable
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	//get variable pointer
	var pointerA *int64 = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

	//get pointer via new keyword
	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

}
