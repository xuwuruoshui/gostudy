package entity

import "fmt"

type Phone struct {
	Name    string
	Brand   string
	Version string
}

func (p *Phone) Call() {
	fmt.Println(p.Name, p.Brand, p.Version)
}
