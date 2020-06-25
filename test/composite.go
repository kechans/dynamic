package main

import "fmt"

//type people struct {
//	name string
//	age  int
//}
//
//type people1 struct {
//	name string
//	age  int
//}
//
//func (p *people) showa() {
//	fmt.Println("p showa")
//	p.showb()
//}
//func (p *people) showb() {
//	fmt.Println("p showb")
//}
//
//type teacher struct {
//	people
//}

//func (t *teacher) showa() {
//	fmt.Println("t showa")
//}
type student struct {
	Name string
	Age  int
}

func pase_student() map[string]student {
	m := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Printf("%p\n", &stu)
		m[stu.Name] = stu
	}
	return m
}

type User struct {
	name string
	age  int
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer : %p, %v \n", self, self)
}
func (self User) TestValue() {
	fmt.Printf("TestValue : %p, %v", &self, self)
}
func main() {
	u := User{name: "aaa", age: 12}
	u.TestPointer()
	u.TestValue()
	//students := pase_student()
	//for k, v := range students {
	//	fmt.Printf("key=%s,value=%v \n", k, v)
	//}
	//t := teacher{}
	//t.showa()
	//arr := []int{1, 23, 4, 5, 0}
	//sli := arr[2:4]
	//sli[1] = 100
	//
	//m := make(map[int]string, 10)
	//m[0] = "aa"
	//m[1] = "bb"
	//m2 := make(map[string]int, 10)
	//m2["aa"] = 1
	//m2["bb"] = 2
	//m2["cc"] = 3
	//for k, v := range m2 {
	//	println(k, v)
	//}
	//arr1 := make([]int, 6, 8)
	//fmt.Printf("%p\n", &arr1)
	//arr1 = append(arr1, 23)
	//fmt.Printf("%p\n", &arr1)
	//fmt.Println(len(sli), cap(sli))
	//t := teacher{}
	//t.showa()
}
