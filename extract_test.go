package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExtractTime(t *testing.T) {
	obj1 := extractTime("星期三第11-12节{7周,11-13周}<br/>星期三第11-13节{2-6周,8-10周}")
	obj2 := extractTime("星期二第11-13节{1-11周}")
	obj3 := extractTime("星期一第11-13节{1-6周}<br/>星期二第11-13节{1-5周}")

	s1, _ := json.Marshal(obj1)
	s2, _ := json.Marshal(obj2)
	s3, _ := json.Marshal(obj3)

	println(string(s1))
	println(string(s2))
	println(string(s3))
}

func TestExtractAddress(t *testing.T) {
	tests := []string{
		"1区5-403", "2区A309", "1区计-203", "3区枫-101",
		"国软4-304", "4区01101<br/>4区01101", "新珈楼B120",
		"桂园网球场", "星湖操场", "卓尔体育馆", "九一二体育场",
	}
	for _, str := range tests {
		marshal, _ := json.Marshal(extractAddress(str))
		fmt.Println(string(marshal))
	}
}

func TestExtractTeacher(t *testing.T) {
	tests := []string{"00008396/杜莉/教授;00006329/李雪松/教授", "00031202/Matthew P.Lutz/副教授"}
	for _, str := range tests {
		marshal, _ := json.Marshal(extractTeacher(str))
		fmt.Println(string(marshal))
	}
}
