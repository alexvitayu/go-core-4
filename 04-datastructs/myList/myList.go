package main

import (
	"fmt"
	"strconv"
)

type Elem struct {
	Val  interface{}
	next *Elem
	prev *Elem
}

type List struct {
	head *Elem
	tail *Elem
}

// Создаёт новый элемент
func NewElem(value interface{}) *Elem {
	return &Elem{Val: value}
}

// Создаёт новый двусвязный список
func NewList() *List {
	return &List{}
}

// Добавляет элемент в конец списка
func (l *List) Push(value interface{}) {
	newElem := NewElem(value)
	if l.head == nil {
		l.head = newElem
		l.tail = newElem
	} else {
		l.tail.next = newElem
		newElem.prev = l.tail
		l.tail = newElem
	}
}

// Выводит содержимое списка
func (l *List) String() string {
	var result string
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%v ", current.Val)
		current = current.next
	}
	return result
}

func main() {
	list := NewList()

	// Добавляем разные элементы
	var data interface{}
	for {
		data = promptData()
		if data == "exit" {
			break
		}
		list.Push(data)
	}

	// Выводим содержимое списка
	fmt.Println("Содержимое двусвязного списка:")
	fmt.Println(list.String())
}

func promptData() interface{} {
	fmt.Print("Введите данные:")
	var input string
	fmt.Scan(&input)
	//attemption to determine the type of input
	if intVal, err := strconv.Atoi(input); err == nil {
		return intVal
	} else if flaotValue, err := strconv.ParseFloat(input, 64); err == nil {
		return flaotValue
	}
	return input
}
