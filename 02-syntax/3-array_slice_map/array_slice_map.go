package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// Объявляем массив из 5 элементов int64
	var arr [5]int64 = [5]int64{10, 20, 30, 40, 50}

	// Получаем адрес третьего элемента (индекс 2, так как индексация с 0)
	firstElementAddr := &arr[0]
	thirdElementAddr := &arr[2]

	fmt.Printf("Адрес первого элемента: %p\n", firstElementAddr)
	fmt.Printf("Значение первого элемента: %d\n", *firstElementAddr)
	fmt.Printf("Адрес третьего элемента: %p\n", thirdElementAddr)
	fmt.Printf("Значение третьего элемента: %d\n", *thirdElementAddr)

	// Альтернативно: через арифметику указателей (используя unsafe)
	arrPtr := uintptr(unsafe.Pointer(&arr[0]))
	thirdAddr := unsafe.Pointer(arrPtr + 2*unsafe.Sizeof(arr[0]))
	fmt.Printf("Адрес через unsafe: %p\n", thirdAddr)

	var array = [...]int{1, 2, 3, 4, 5} // массив
	slice := []int{0, 1, 2, 3, 4, 5}    // слайс
	slice = append(slice, 10)

	fmt.Println(slice[1:4])

	_ = slice[len(slice)-1]
	fmt.Println("array", array, len(array), cap(array))
	fmt.Println("slice", slice, len(slice), cap(slice))

	slice = make([]int, 10, 20)
	fmt.Println("\nslice после инициализации:", slice, len(slice), cap(slice))

	// ассоциативный массив
	var m map[int]string
	//m[3] = "ABC" // panic: assignment to entry in nil map
	//m := make(map[int]string)
	m = map[int]string{}
	m[3] = "ABC"
	m[6] = "DEF"
	for k, v := range m {
		fmt.Printf("ключ %d, значение %s\n", k, v)
	}
}

func sliceDimensions() {
	var slice []int
	c := cap(slice)
	for i := 0; i < 1_000_000; i++ {
		slice = append(slice, i)
		if cap(slice) != c {
			fmt.Printf("длина: %v\tемкость:%v\tкоэффициент:%v\n", len(slice), cap(slice), float64(cap(slice))/float64(c))
			c = cap(slice)
		}
	}
}
