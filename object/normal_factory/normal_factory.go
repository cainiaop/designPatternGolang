package normal_factory

import "errors"

//工厂
type Factory interface{
	Create() Operation
}

type OperationAddFactory struct {

}

func (o OperationAddFactory) Create() Operation {

	return OperationAdd{}
}

type OperationSubFactory struct {

}

func (o OperationSubFactory) Create() Operation {

	return OperationSub{}
}

type OperationMulFactory struct {

}

func (o OperationMulFactory) Create() Operation {

	return OperationMul{}
}

type OperationDivFactory struct {

}

func (o OperationDivFactory) Create() Operation {

	return OperationDiv{}
}


//产品
type Operation interface {
	Execute(l, r int) int
}

type OperationAdd struct {
}

func (o OperationAdd) Execute(l, r int) int {

	return l + r
}

type OperationSub struct {

}

func (o OperationSub) Execute(l, r int) int {

	return l - r
}

type OperationMul struct {

}

func (o OperationMul) Execute(l, r int) int {

	return l * r
}

type OperationDiv struct {

}

func (o OperationDiv) Execute(l, r int) int {
	if r == 0 {
		panic(errors.New("被除数不能为空"))
	}

	return l / r
}