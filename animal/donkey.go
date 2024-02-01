package animal

import "fmt"

type Donkey struct {
	Name string
	Type string
}

func (d Donkey) Breath() {
	fmt.Println("breathing")
}
func (d Donkey) Born() string {
	return fmt.Sprintf("%s is Born", d.Name)
}
func (d Donkey) Died() string {
	return fmt.Sprintf("%s is Diad", d.Name)
}
