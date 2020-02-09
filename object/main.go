package main

import (
	"fmt"
	"test/design/object/abstract_factory"
	"test/design/object/easy_factory"
	"test/design/object/normal_factory"
	"test/design/object/singleton"
)

func main() {

	//===简单工厂模式===
	operation := "/"
	//实例化产品
	var opObject easy_factory.Operation
	opObject = easy_factory.OperationFactory(operation)
	fmt.Println(opObject.Execute(10, 1))

	//===工厂模式===
	//实例化工厂
	var opFactory normal_factory.Factory
	opFactory = normal_factory.OperationAddFactory{}
	//实例化产品
	var opObject1 normal_factory.Operation
	opObject1 = opFactory.Create()
	fmt.Println(opObject1.Execute(10, 1))

	//===抽象工厂模式===
	//工厂
	var cdFactory, shFactory abstract_factory.Factory
	cdFactory = abstract_factory.ChengduFactory{}
	shFactory = abstract_factory.ShanghaiFactory{}
	//产品
	var husband abstract_factory.Husband
	var wife abstract_factory.Wife
	//成都系列
	husband = cdFactory.CreateHusband()
	wife = cdFactory.CreateWife()
	husband.Eat()
	wife.Eat()
	//上海系列
	husband = shFactory.CreateHusband()
	wife = shFactory.CreateWife()
	husband.Eat()
	wife.Eat()

	//===单例模式===
	object1 := singleton.New(10)
	object1.Show()
	object2 := singleton.New(100)
	object2.Show()
	fmt.Println(object1 == object2)





}