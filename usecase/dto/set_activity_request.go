package dto

type SetActivityRequest struct {
	Activities []Activity
}

type Activity struct {
	Date string `json:"date"`
	Schedule string `json:"schedule"`
	SecondFloorTeacherName string `json:"second_floor_teacher_name"`
	ThirdFloorTeacherName string `json:"third_floor_teacher_name"`
	FourthFloorTeacherName string `json:"forth_floor_teacher_name"`
}
