package abstract_factory

import "fmt"

//工厂类
type Factory interface{
	CreateWife() Wife
	CreateHusband() Husband
}

type ShanghaiFactory struct {

}

func (f ShanghaiFactory) CreateWife() Wife {

	return ShanghaiWife{}
}

func (f ShanghaiFactory) CreateHusband() Husband {

	return ShanghaiHusband{}
}


type ChengduFactory struct {

}

func (f ChengduFactory) CreateWife() Wife {

	return ChengduWife{}
}

func (f ChengduFactory) CreateHusband() Husband {

	return ChengduHusband{}
}




//产品类

//妻子
type Wife interface {
	Eat()
}

type ShanghaiWife struct {

}

func (w ShanghaiWife) Eat() {

	fmt.Println("上海妻子吃起来很精致")
}

type ChengduWife struct {

}

func (u ChengduWife) Eat() {
	fmt.Println("成都妻子吃起来很豪放")
}


//丈夫
type Husband interface {
	Eat()
}

type ShanghaiHusband struct {

}

func (u ShanghaiHusband) Eat() {
	fmt.Println("上海丈夫吃起来很猥琐")
}

type ChengduHusband struct {

}

func (u ChengduHusband) Eat() {
	fmt.Println("成都丈夫吃起来很耙耳朵")
}

