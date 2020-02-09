package main

import (
	"test/design/optimize/decorative"
	"test/design/optimize/make_up"
)

func main() {

	//===组合模式===
	var sMerchant,bigMerchant make_up.Merchant

	sMerchant = &make_up.SingleMerchant{Name: "pwr"}
	//sMerchant.Show(1)

	bigMerchant = &make_up.BigMerchant{Name: "hx"}
	bigMerchant.Add(sMerchant)
	//bigMerchant.Show(1)

	lastMerchant := &make_up.BigMerchant{Name: "pjm"}
	lastMerchant.Add(bigMerchant)
	//lastMerchant.Show(0)

	superMerchant := &make_up.BigMerchant{Name: "pjh"}
	superMerchant.Add(lastMerchant)
	//superMerchant.Show(0)

	last2Merchant := &make_up.BigMerchant{Name: "pzm"}

	superMerchant.Add(last2Merchant)
	superMerchant.Show(0)


	superMerchant.Remove("pzm")
	superMerchant.Show(0)


	//===装饰模式===
	var person decorative.Person
	person = decorative.Person{Name:"pjm"}
	person.Eat("noodle")

	tshirt := decorative.TShirt{}
	tshirt.Decorative(person)
	//tshirt.Show()

	hat := decorative.Hat{}
	hat.Decorative(&tshirt)
	hat.Show()



}
