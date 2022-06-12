package model

type student struct { // 结构体名首字母小写，则仅能包内访问
	Name string
	age  int
}

func CreateStudent(name string, age int) *student {
	// 暴露函数名首字母大写的函数，重当构造方法
	return &student{name, age}
}

func (student *student) GetAge() int {
	return student.age
}
