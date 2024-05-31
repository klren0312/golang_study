package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "you are sb"
	} else {
		talk = "hello"
	}
	return
}

type SbNo struct {
	ID     int `json:"id"`
	Gender string
	name   string
}

func main() {
	s1 := SbNo{
		ID:     1,
		Gender: "男",
		name:   "ttest",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak((think)))

	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 11
	fmt.Printf("%#v\n", user)
	var p1 person
	p1.name = "test"
	p1.city = "sh"
	p1.age = 11
	fmt.Printf("p1 = %v\n", p1)
	fmt.Printf("p1 = %#v\n", p1)
	sourceMap := make(map[string]int, 8)
	sourceMap["张1"] = 90
	sourceMap["明1"] = 100
	sourceMap["张2"] = 90
	sourceMap["明2"] = 100
	sourceMap["张3"] = 90
	sourceMap["明3"] = 100
	sourceMap["张4"] = 90
	sourceMap["明4"] = 100
	sourceMap["张5"] = 90
	sourceMap["明5"] = 100
	sourceMap["张6"] = 90
	sourceMap["明6"] = 100
	sourceMap["张7"] = 90
	sourceMap["明7"] = 1000000000000000000
	fmt.Println(sourceMap)
	delete(sourceMap, "明1")
	fmt.Println(sourceMap["明1"])
	fmt.Printf("type of a:%T\n", sourceMap)

	v, ok := sourceMap["张"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("无")
	}

	for k, v := range sourceMap {
		fmt.Println(k, v)
	}

}
