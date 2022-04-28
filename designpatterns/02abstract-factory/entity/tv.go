package entity

import "fmt"

type TV struct {
	Name    string
	Brand   string
	Version string
}

func (t *TV) Play() {
	fmt.Println(t.Name, t.Brand, t.Version)
}
