package main

import "fmt"
import "slices"

func main() {
	link()
}

/*
Переменные после объявления содержать дефолтные значения, даже если не было ничего присвоино
Нельзя создавать ещё переменную с уже существующим именем

int8: представляет целое число от -128 до 127 и занимает в памяти 1 байт (8 бит)
int16: представляет целое число от -32768 до 32767 и занимает в памяти 2 байта (16 бит)
int32: представляет целое число от -2147483648 до 2147483647 и занимает 4 байта (32 бита)
int64: представляет целое число от –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807 и занимает 8 байт (64 бита)
uint8: представляет целое число от 0 до 255 и занимает 1 байт
uint16: представляет целое число от 0 до 65535 и занимает 2 байта
uint32: представляет целое число от 0 до 4294967295 и занимает 4 байта
uint64: представляет целое число от 0 до 18 446 744 073 709 551 615 и занимает 8 байт
byte: синоним типа uint8, представляет целое число от 0 до 255 и занимает 1 байт
rune: синоним типа int32, представляет целое число от -2147483648 до 2147483647 и занимает 4 байта
int: представляет целое число со знаком, которое в зависимости о платформы может занимать либо 4 байта, либо 8 байт. То есть соответствовать либо int32, либо int64.
uint: представляет целое беззнаковое число только без знака, которое, аналогично типу int, в зависимости о платформы может занимать либо 4 байта, либо 8 байт. То есть соответствовать либо uint32, либо uint64.

float32: представляет число с плавающей точкой от 1.4*10-45 до 3.4*1038(для положительных). Занимает в памяти 4 байта (32 бита)
float64: представляет число с плавающей точкой от 4.9*10-324 до 1.8*10308 (для положительных) и занимает 8 байт.

complex64: комплексное число, где вещественная и мнимая части представляют числа float32
complex128: комплексное число, где вещественная и мнимая части представляют числа float64

bool: true (истина) или false (ложь).

string: строка

	\n: переход на новую строку
	\r: возврат каретки
	\t: табуляция
	\": двойная кавычка внутри строк
	\\: обратный слеш
*/
func variables() {
	//var (
	//	a int8   = -1
	//	b uint8  = 2
	//	c byte   = 3 // byte - синоним типа uint8
	//	d int16  = -4
	//	f uint16 = 5
	//	g int32  = -6
	//	h rune   = -7 // rune - синоним типа int32
	//	j uint32 = 8
	//	k int64  = -9
	//	l uint64 = 10
	//	m int    = 102
	//	n uint   = 105
	//
	//	f  float32 = 18
	//	g  float32 = 4.5
	//	d  float64 = 0.23
	//	pi float64 = 3.14
	//	e  float64 = 2.7
	//
	//	str string = "Hello World"
	//)

	var hello string // = "" по умолчанию

	var hello2 string = "Hello2"
	var hello3 = "Hello3"
	hello4 := "Hello4"

	var (
		hello5 string = "Hello5"
		hello6 string = "Hello6"
	)

	var name, age = "Tom", 27

	fmt.Println(hello)
	fmt.Println(hello2)
	fmt.Println(hello3)
	fmt.Println(hello4)
	fmt.Println(hello5)
	fmt.Println(hello6)
	fmt.Println(name)
	fmt.Println(age)

}

func constants() {
	const pi float64 = 3.1415

	const (
		a = 1
		b
		c
		d = 3
		f
	)
}

func arithmetic() {
	//var m float32 = 10 / 4 // 2
	//var m2 float32 = 10 / 4.0 //2.5
}

func array() {
	var numbers0 []string                       // []
	var numbers [5]int                          //[0 0 0 0 0]
	var numbers2 [5]int = [5]int{1, 2, 3, 4, 5} //[1 2 3 4 5]
	var numbers3 [5]int = [5]int{1, 2}          // [1 2 0 0 0]

	numbers4 := [5]int{1, 2, 3, 4, 5} //[1 2 3 4 5]
	numbers5 := [...]int{1, 2, 3}     //[1 2 3]

	fmt.Println(numbers[0])
	fmt.Println(numbers0, numbers, numbers2, numbers3, numbers4, numbers5)

	//var numbers10 [3]int = [3]int{1, 2, 3}
	//var numbers211[4]int = [4]int{1, 2, 3, 4}
	//numbers = numbers2  // ! Ошибка

	colors := [...]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue
}

/*
срезы могут менять свой размер (как arrayList)
make () создаёт slice и заполняет его по дефолту
range () - раскрывает массивы/срезы, возвращая ключ и значение (для for и тд)
append(slice, value) - добавить в слайз новый элемент
*/
func slice() {
	//var users []string
	//fmt.Println(users)

	var users []string = make([]string, 46)
	users[0] = "Tom"
	users[1] = "Alice"
	users[2] = "Bob"

	fmt.Println(users)

	//удаляем 4-й элемент
	users2 := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	users2 = append(users2[:n], users2[n+1:]...)
	fmt.Println(users2) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

	//удалаем через библиотеку
	fmt.Println(slices.Delete(users2, 0, 3))
}

/*
хэш мапа или аасоциативный масси
delete удалает элемент
*/
func mapTest() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people["Bob"]) // 1

	if val, ok := people["Alice"]; ok { //проверка на ключ
		fmt.Println(val)
	}

	var people2 = map[string]int{"Tom": 1, "Bob": 2, "Sam": 8}
	delete(people2, "Bob")
	fmt.Println(people2)
}

func logic(param ...int) int { // неопределённое кол-во параметров и возвращает тип int
	a := 9
	switch a {
	case 9:
		fmt.Println("a = 9")
		fallthrough //так же запускает кейс снизу
	case 0:
		fmt.Println("a = 0 или 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	case 6, 5, 4:
		fmt.Println("a = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной a не определено")

	}

	for i := 1; i < param[0]; i++ {
		fmt.Println(i * i)
	}

	var users = [3]string{"Tom", "Alice", "Kate"}
	for _, value := range users {
		fmt.Println(value)
	}

	for _, value := range [3]string{"Tom2", "Alice2", "Kate2"} {
		fmt.Println(value)
	}

	return 5
}

func functions(i int) (z int, fullName string) {
	z = i
	fullName = "Full Name"

	return
}

func functions2() {

	//func add(x int, y int) int {
	//
	//	return x + y
	//}
	//func multiply(x int, y int) int {
	//
	//	return x * y
	//}
	//func action(n1 int, n2 int, operation func(int, int) int){
	//
	//	result := operation(n1, n2)
	//	fmt.Println(result)
	//}
	//func main() {
	//
	//	action(10, 25, add)     // 35
	//	action(5, 6, multiply)  // 30
	//}
}
func function3() {
	//func action(n1 int, n2 int, operation func(int, int) int){
	//
	//	result := operation(n1, n2)
	//	fmt.Println(result)
	//}
	//func main() {
	//
	//	action(10, 25, func (x int, y int) int { return x + y })    // 35
	//	action(5, 6, func (x int, y int) int { return x * y })      // 30
	//}
}
func function4() {
	//func square() func() int {
	//	var x int = 2
	//	return func() int {
	//		x++
	//		return x * x
	//	}
	//}
	//
	//func main() {
	//
	//	f := square()
	//	fmt.Println(f())        // 9
	//	fmt.Println(f())        // 16
	//	fmt.Println(f())        // 25
	//}
}

func factorial() {
	//func factorial(n uint) uint{
	//
	//	if n == 0{
	//	return 1
	//}
	//	return n * factorial(n - 1)
	//}
	//func main() {
	//
	//	fmt.Println(factorial(4))   // 24
	//	fmt.Println(factorial(5))   // 120
	//	fmt.Println(factorial(6))   // 720
	//}
}

func link() {
	var x int = 4
	var p *int = &x            // указатель получает адрес переменной
	fmt.Println("Address:", p) // значение указателя - адрес переменной x
	fmt.Println("Value:", *p)  // значение переменной x

	*p = 25
	fmt.Println(x)
	fmt.Println("Address:", p)

	//может быть nil, потому лучше проверять!
	var pf *float64
	if pf != nil {
		fmt.Println("Value:", *pf)
	}

	//Можно передать адресс переменной в другую функцию, где переменная будет та же и будет изменена без return
	//func changeValue(x *int){
	//	*x = (*x) * (*x)
	//}
	//func main() {
	//
	//	d := 5
	//	fmt.Println("d before:", d)     // 5
	//	changeValue(&d)                 // изменяем значение
	//	fmt.Println("d after:", d)      // 25 - значение изменилось!
	//}

}

/*
Именнованный тип, стогая типизация
*/
func typeTest() {

	//type mile uint
	//type kilometer uint
	//
	//func distanceToEnemy(distance mile) {
	//
	//	fmt.Println("расстояние для противника:")
	//	fmt.Println(distance, "миль")
	//}

	//var distance mile = 5
	//distanceToEnemy(distance) //норм

	// var distance1 uint = 5
	// distanceToEnemy(distance1)   // !Ошибка

	// var distance2 kilometer = 5
	// distanceToEnemy(distance2)   // ! ошибка
}

/*
псевдноним, нестрогая типизация
*/
func typeTest2() {
	//type mile = uint
	//type kilometer = uint
	//
	//func distanceToEnemy (distance mile){
	//
	//	fmt.Println("расстояние для противника:")
	//	fmt.Println(distance, "миль")
	//}
	//
	//func main() {
	//
	//	var distance mile = 5
	//	distanceToEnemy(distance)
	//
	//	var distance1 uint = 5
	//	distanceToEnemy(distance1)   // норм
	//
	//	var distance2 kilometer = 5
	//	distanceToEnemy(distance2)   // норм
	//}
}

/* выполняется всегда в конце выполнения функции, в котором было вызвано */
func deferTest() {
	printFinish := func() { fmt.Println("Finish") }

	printFinish()
	fmt.Println("Start work")
	fmt.Println("Working")
}

/* Создание исключений */
func panicTest(a, b int) int {
	if b == 0 {
		panic("Division by zero!")
	}
	return a / b
}

/*
Создаёт экземляр класса, но отдаёт указатель (ссылку)
*/
func newTest() {
	p := new(int)             // p - тут это указатель
	fmt.Println("Value:", *p) // Value: 0 - значение по умолчанию
	*p = 8                    // изменяем значение
	fmt.Println("Value:", *p) // Value: 8
}
