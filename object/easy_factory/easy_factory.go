package easy_factory

import (
	"errors"
)

//工厂
func OperationFactory(op string) Operation {

	switch op {

		case "+" :
			return OperationAdd{}
	 	case "-" :
	 		return OperationSub{}
		case "*":
			return OperationMul{}
		case "/":
			return OperationDiv{}
		default:
			panic(errors.New("undefined operation type"))
	}
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