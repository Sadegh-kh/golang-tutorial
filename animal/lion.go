package animal

import "fmt"

type Lion struct {
	Name  string
	Type  string
	Level int
}

func (l Lion) Breath() {
	fmt.Println("breathing")
}
func (l Lion) Born() string {
	return fmt.Sprintf("%s is Born", l.Name)
}
func (l Lion) Died() string {
	return fmt.Sprintf("%s is Diad", l.Name)
}
