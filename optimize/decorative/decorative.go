package decorative

import "fmt"

//人
type Person struct {
	Name string
}

//核心功能
func (p Person) Eat(name string) {
	fmt.Println(p.Name + "吃" + name)
}

//装饰 从 person中脱离出来
func (p Person) Show() {

	fmt.Printf(p.Name + "装扮: ")
}

type BaseInterface interface {
	Show()
}

//装饰
type Dress struct {
	b BaseInterface
}

func (d *Dress) Decorative(bt BaseInterface) {
	d.b = bt
}

func (d *Dress) baseShow() {
	if d.b != nil {
		d.b.Show()
	}
}

type TShirt struct {
	Dress
}

func (t *TShirt) Show() {
	t.baseShow()
	fmt.Printf(" 穿TShirt")
}

type Hat struct {
	Dress
}

func (t *Hat) Show() {
	t.baseShow()
	fmt.Printf(" 戴帽子")
}
