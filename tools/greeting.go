package tools

import "fmt"

func Greeting(name string, age int) string {
	var result string
	if age < 12 {
		result = fmt.Sprintf("welcome %s , to my world you are kide ", name)
	} else if age > 12 && age < 20 {
		result = fmt.Sprintf("welcome %s , to my world you are teenager ", name)
	} else {
		result = fmt.Sprintf("welcome %s , to my world you are %d years old ", name, age)
	}
	return result
}
