package log

import (
	"fmt"
	"testing"
)

type Info1 struct {
	Name string
	Age  int
}

func (p *Info1) GetName() string {
	return p.Name
}

func (p Info1) GetAge() int {
	return p.Age
}

type Mc struct {
	Info1
}

type Mc1 struct {
	*Info1
}

func TestLog(t *testing.T) {
	a := &Mc1{&Info1{
		"mc",
		12,
	}}

	fmt.Println(a.GetAge())
	fmt.Println(a.GetName())
}
