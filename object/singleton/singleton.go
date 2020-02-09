package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
	v int
}

var instance *singleton
var lock sync.Once


func New(v int) *singleton {

	lock.Do(func() {
		instance = &singleton{v:v}
	})

	return instance
}

func (s *singleton) Show() {
	fmt.Println(fmt.Sprintf("hello world...Value : %d", s.v))
}