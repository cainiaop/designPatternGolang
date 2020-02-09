package make_up

//组合模式：用户一致对待单个对象和组合对象
import (
	"fmt"
)

type Merchant interface {
	Add(m Merchant)
	Remove(name string)
	Show(depth int)
	GetName() string
}

//单个对象
type SingleMerchant struct {
	Name string
	list []Merchant
}

func (s *SingleMerchant) Add(m Merchant) {
	fmt.Println("个体户不允许增加子商户")
}

func (s *SingleMerchant) Remove(name string) {

	fmt.Println("个体户没有子商户")
}

func (s *SingleMerchant) Show(depth int) {
	var pre string
	for i := 0; i < depth; i++ {
		pre += "-"
	}
	fmt.Println(fmt.Sprintf("%s%s", pre, s.Name))
}

func (s *SingleMerchant) GetName() string {
	return s.Name
}

//组合对象
type BigMerchant struct {
	Name string
	list []Merchant
}

func (s *BigMerchant) Add(m Merchant) {
	s.list = append(s.list, m)
}

func (s *BigMerchant) Remove(name string) {

	var index int
	for k, v := range s.list {
		if v.GetName() == name {
			index = k
			break
		}
	}

	s.list = append(s.list[:index], s.list[index+1:]...)
}

func (s *BigMerchant) Show(depth int) {

	var pre string
	for i := 0; i < depth; i++ {
		pre += "-"
	}
	fmt.Println(fmt.Sprintf("%s%s", pre, s.Name))

	count := len(s.list)
	for i := 0; i < count; i++ {
		s.list[i].Show(depth+1)
	}
}

func (s *BigMerchant) GetName() string {
	return s.Name
}

