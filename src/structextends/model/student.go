package model

type Student struct {
	Name string
	Age  int
}

type Graduate struct {
	Student // 匿名结构体
	Major   string
}

func (student *Student) GetAge() int {
	return student.Age
}

func (graduate *Graduate) GetMajor() string {
	return graduate.Major
}
