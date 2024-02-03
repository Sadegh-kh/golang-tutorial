// statement
package main

import (
	_ "code/setup"
	"fmt"
)

func init() {
	fmt.Println("init main go 0")
}

func squence(num int) int {
	if num > 100 && num < 120 {
		return num
	}
	return num * num
}

func main() {

}

/*
// inheritance

type person struct {
	Name string
}

func (p person) GetName() string {
	return p.Name
}

type student struct {
	person
	Number int
}

	fmt.Printf("type of %v is %v \n", 1, reflect.TypeOf(1))
	s := student{
		Number: 2342,
		person: person{Name: "ali"},
	}
	fmt.Println(s.GetName())
// type assertion
func AnimalCheck(a animal.Animal) {
	// type assertion

	// l, ok := a.(animal.Lion)
	// if ok {
	// 	fmt.Println("this animal is lion and level is", l.Level)
	// } else {
	// 	d, ok := a.(animal.Donkey)
	// 	if ok {
	// 		fmt.Println("this donkey and name is", d.Name)
	// 	} else {
	// 		fmt.Println("it's unknown Animal")
	// 	}
	// }

	switch myAniml := a.(type) {
	case animal.Lion:
		fmt.Println("this animal is lion and level is", myAniml.Level)
	case animal.Donkey:
		fmt.Println("this is dockey and its name", myAniml.Name)
	default:
		fmt.Println("this is unkonwn animal")
	}

}
myLion := animal.Lion{
		Name:  "risy",
		Type:  "denger",
		Level: 9,
	}

	AnimalCheck(myLion)
	myDonkey := animal.Donkey{
		Name: "rostam",
		Type: "nader",
	}
	AnimalCheck(myDonkey)
// interface
AppUser := app.User{
		Name: "Sadegh",
		Age:  21,
	}
	storageUserMap := map[string]app.User{}
	memory := storage.Memory{
		Users: storageUserMap,
	}

	application := app.App{
		Name:    "Bazar",
		Storage: &memory,
	}
	application.CreateUser(AppUser)
	fmt.Println("my user in memory : ", application.GetUser("Sadegh"))

	application.Storage = storage.MyFile{}
	// application.CreateUser(AppUser)
	fmt.Println("my user in memory : ", application.GetUser("Sadegh"))

// new method
user := User{ID: 1,
		name: "Sadegh"}

	fmt.Println("my user : ", user)
	fmt.Printf("my user: %T \n", user)
	user.ID = 2
	fmt.Println("my user : ", user)

	u := new(User)
	fmt.Printf("my user1: %T \n", u)
	fmt.Printf("my user1: %v \n", u)
	u.name = "Ali"
	fmt.Printf("my user1: %v \n", *u)

	i := new(int)

	fmt.Printf("my i: %T \n", i)
	fmt.Printf("my i: %v \n", *i)

	var uu User
	uu.name = "ali"
	fmt.Printf("uu type %T \n", uu)
	fmt.Printf("uu type %v \n", uu)

// goto
	i := 0
start:
	fmt.Println("i :", i)

	i++
	if i < 10 {
		goto start
	}
// fmt.Println(Greeting("Sadegh",22))
	// fmt.Println(Greeting("Sadegh",10))
	// fmt.Println(Greeting("Sadegh",15))
	// fmt.Printf("wellcome %s","ali")

	// array and slice
	array := []string{"al8i", "sadegh", "abas"}
	array = append(array, "saeed")
	// fmt.Print(array[:])
	// for index,name := range {}
	// if not using index must be _
	// for _,name := range array{
	// 	println(name)
	// }
	// second formt of for
	// for i:=1 ; i<100 ; i++{
	// 	if i % 7 == 0 {
	// 		println(i)
	// 	}
	// }

	// dictionary in go (map)

	// products := map[string]int{
	// 	"shoes":2000,
	// 	"shirt":1500,
	// 	"pants":300,
	// 	"T-shirt":3000,
	// }
	// coat , ok :=products["coat"]
	// if ok{
	// 	println("price of shoes ",coat)
	// }else{
	// 	println("this prodect not exist")
	// }
	// products["coat"] = 10000
	// for product,price := range products{
	// 	fmt.Println(product,price)
	// }

	// practice
	// var name = "hello world"
	// count_word := tls.Counter(name)
	// for key, value := range count_word {
	// 	println(key, value)
	// }
	// fmt.Println(count_word)

	// x, y := "hello", "world"
	// fmt.Println(x, y)
	// x, y = swap(x, y)
	// fmt.Println(x, y)

	// switch
	a := 122
	switch {
	case a < 10:
		fmt.Println("mardud")
	case 10 <= a && a < 15:
		fmt.Println("khub")
	case 15 <= a && a <= 20:
		fmt.Println("awli")
	default:
		fmt.Println("invalid number")
	}

	var student1 structuers.Student = structuers.Student{
		ID:        2123123,
		FirstName: "Sadegh",
		LastName:  "Khoshbayan",
		Age:       22,
	}
	fmt.Println(student1)

	// point of slice (capisity and lengh)

	MySlice := []int{1, 2, 3}
	fmt.Println(cap(MySlice), len(MySlice))
	MySlice = append(MySlice, 15)
	fmt.Println(cap(MySlice), len(MySlice))


	// os args
	// if len(os.Args) <= 1 {
	// 	fmt.Println("error , not have input")
	// 	return
	// }
	// name := os.Args[1]
	// fmt.Printf("Hello , %s", name)

	// flags
	// name := flag.String("name", "default name", "user name")
	// age := flag.Int("age", 23, "user age")
	// flag.Parse()
	// fmt.Printf("hello %s you are %d years old ", *name, *age)
	// // remaining argument
	// rFlag := flag.Args()
	// fmt.Println(rFlag)

	// var num1, num2, num3, num4 int

	// fmt.Scanln(&num1, &num2, &num3, &num4)
	// fmt.Println(num1, num2, num3, num4)

	// map
	// var myMap = map[string]string{}
	// myMap["apple"] = "red"
	// value, ok := myMap["bnana"]
	// fmt.Printf("bnana value %s and ok is %t \n", value, ok)
	// value, ok = myMap["apple"]
	// fmt.Printf("apple value %s and ok is %t \n", value, ok)

	// function value
	// myComputer := computer(3, power)
	// fmt.Println(myComputer)

	// anonymous function
	// var myFunc = func(x int) int {
	// 	return x * 2
	// }
	// fmt.Println(myFunc(3))

	// variadic function
	// var myGrade = adjusted(20, 18, 19, 15, 14, 17)
	// println(myGrade)

	// method
	// var student2 = structuers.Student{
	// 	ID:        1231231,
	// 	FirstName: "Ali",
	// 	LastName:  "Abdi",
	// 	Age:       24,
	// 	Scores:    []float32{12.3, 20, 10.5, 7.9, 16},
	// }
	// fmt.Printf("student %s with Avg scoures %.2f is %t", student2.FirstName+student2.LastName, (student2.ScoresAvg()), student2.PassedStudent())

	// pointer
	// var i = 13
	// var p = &i
	// fmt.Println("address i:", &i)
	// fmt.Printf("address p:%p value p:%d value i:%d\n", p, *p, i)
	// *p = 134
	// fmt.Println("i:", i)

	// pointer in method
	// fmt.Printf("firstname %s and lastname %s\n", student1.FirstName, student1.LastName)
	// student1.SetName("Abas", "Pildare")
	// fmt.Println(student1.GetName())

	// var myArray []int
	// myArray = append(myArray, 3)
	// fmt.Println(myArray)

	// sleep
	// time.Sleep(10 * time.Second)
*/

func adjusted(values ...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum / len(values)
}
func power(x, t int) int {
	var result = x
	for i := 1; i <= t; i++ {
		result *= i
	}
	return result
}

func computer(x int, power func(x, t int) int) int {
	myNumber := x
	return power(myNumber, 4)

}

func swap(x, y string) (string, string) {

	return y, x

}
