// пример реализации двусвязного списка
package main

import "fmt"

// элемент двусвязного списка
type Elem struct {
	Val  interface{}
	next *Elem
	prev *Elem
}

// двусвязный список
type List struct {
	root *Elem
}

// создаём новый элемент
func NewElem(val interface{}) *Elem {
	return &Elem{Val: val}
}

// создаём новый двусвязный список
func NewList() *List {
	root := &Elem{}
	root.next = root
	root.prev = root
	return &List{root: root}
}

// вставляем элемент в начало списка
func (l *List) PushForward(val interface{}) {
	newElem := NewElem(val)    //создаём новый элемент
	newElem.prev = l.root      //указываем, что предыдущий для нового - корень
	newElem.next = l.root.next //следующий элемент для нового  - текущий первый элемент
	l.root.next.prev = newElem //теперь для старого первого предыдущий - новый элемент
	l.root.next = newElem      //корень указывает на новый элемент
}

// вставляем элемент в конец списка
func (l *List) PushBack(val interface{}) {
	newElem := NewElem(val)    //создали
	newElem.prev = l.root.prev //предыдущий элемент нового - предыдущий эл корня
	newElem.next = l.root      //следующий для нового - корень
	l.root.prev.next = newElem //теперь следующий для последнего - новый элем.
	l.root.prev = newElem      //теперь для корня предыдущий -новый элемент
}

func (l *List) Print() {
	current := l.root.next
	for current != l.root {
		fmt.Printf("%v ", current.Val)
		current = current.next
	}
	fmt.Println()
}

// выводим содержимое списка
func main() {
	l := NewList()
	l.PushForward(10)
	l.PushForward(20)
	l.PushForward(30)
	l.PushForward(40)
	l.PushForward(50)
	l.PushForward(60)
	l.Print() //выводит 60, 50, 40, 30, 20, 10

	q := NewList()
	q.PushBack(10)
	q.PushBack(20)
	q.PushBack(30)
	q.PushBack(40)
	q.PushBack(50)
	q.PushBack(60)

	q.Print() //выводит 10, 20, 30, 40, 50 ,60
}
