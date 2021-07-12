package dto

type SetStudentRequest struct {
	Students []Student
}

type Student struct {
	Num  string `json:"num"`
	Name string `json:"name"`
}
