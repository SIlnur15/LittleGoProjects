package main

import "fmt"

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	slice2 := make([]Child, len(p.Children))
	copy(slice2, p.Children) // теперь у slice2 другая ссылка
	return Parent{Name: p.Name, Children: slice2}
}

func main() {
	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}

	cp := CopyParent(p)

	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	cp.Children[0] = Child{
		Name: "Gosha",
		Age:  30,
	}

	fmt.Println(p.Children)  // -> [{Andy 18}]
	fmt.Println(cp.Children) // -> [{Gosha 30}]

	cp2 := CopyParent(nil) // -> Parent{}
	fmt.Println(cp2)
}
