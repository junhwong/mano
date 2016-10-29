package opc

import (
	"testing"

	"github.com/junhwong/mano/otpl/common"
)

func getFunc() common.TemplateFunc {
	return func(params ...interface{}) (r interface{}, err error) {
		return nil, nil
	}
}

type myStruct struct {
	name string
}

func TestRef(t *testing.T) {
	//tempFunc := getFunc()

	//t.Fatal(reflect.TypeOf(tempFunc)) //common.TemplateFunc

	my := &myStruct{
		name: "张三",
	}

	// t.Fatal(reflect.TypeOf(my)) //*opc.myStruct

	// t.Fatal(reflect.TypeOf(my)) //*opc.myStruct

	// tp:=reflect.TypeOf(my)

	// fd,ok:=tp.FieldByName("")

	// val:=reflect.ValueOf(my)
	// val.Call()

	arr := []int{1, 2, 3}

	if arr != nil && my != nil {

	}

	m := make(map[int]string, 0)
	m[1] = "11111"
	m[3] = "33333"
	//val := reflect.ValueOf(m)
	it := newIterator(m, nil)

	t.Fatal(it)
}
