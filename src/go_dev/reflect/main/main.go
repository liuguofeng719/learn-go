package main

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	name    string `json:"name"`
	age     int8   `json:"age"`
	address string `json:"address"`
}

func (u *User) printInfo() {
	fmt.Printf("name=%v,age=%v,address=%v", u.name, u.age, u.address)
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetAge(age int8) {
	u.age = age
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func NewUser() *User {
	return &User{
		name:    "list",
		age:     20,
		address: "四川省成都市",
	}
}

// 反射
func main() {
	var money float32 = 22.00
	varType := reflect.TypeOf(money)
	value := reflect.ValueOf(money)
	fmt.Println(varType.Kind() == reflect.Float32)
	fmt.Printf("%v = %v \n", varType, value)
	//reflectFieldStruct()
	reflectMethodStruct()
}

func reflectMethodStruct() {
	user := NewUser()
	//typeOf := reflect.TypeOf(user)
	value := reflect.ValueOf(user)

	valMethod := value.NumMethod()
	log.Printf("method size=%v", valMethod)
	methodByName := value.MethodByName("SetName")

	var params []reflect.Value
	params = append(params, reflect.ValueOf("zhangsan"))
	methodByName.Call(params)

	fmt.Printf("设置的值=%v", user.name)

	//method := typeOf.NumMethod()
	//elem := typeOf.Elem()
	//for i := 0; i < method; i++ {
	//	Type := elem.Method(i).Type
	//	methodName := elem.Method(i).Name
	//	if methodName == "SetName" {
	//		reflect.ValueOf(Type).MethodByName(methodName).SetString("ssss")
	//	}
	//	log.Printf("methodName=%v,value= %v", methodName, reflect.ValueOf(Type))
	//}
}

func reflectFieldStruct() {

	var user = &User{}
	typeOf := reflect.TypeOf(user)

	// Elem只支持5种类型 Array、Chan、Map、Ptr、Slice
	elem := typeOf.Elem()
	field := elem.NumField()
	for i := 0; i < field; i++ {
		name := elem.Field(i).Name
		tag := elem.Field(i).Tag
		path := elem.Field(i).PkgPath
		ints := elem.Field(i).Index
		log.Printf("name=%v,tag=%v,path=%v,index=%v \n", name, tag, path, ints)
	}
}
