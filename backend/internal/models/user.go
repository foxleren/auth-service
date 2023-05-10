package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}

type UserNameInfo struct {
	Id         int    `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	FirstName  string `json:"first_name" db:"first_name"`
	SecondName string `json:"second_name" db:"second_name"`
	ThirdName  string `json:"third_name" db:"third_name"`
}

type UserIn struct {
	Id    int    `json:"id" db:"id"`
	Email string `json:"email" db:"email" binding:"required"`
	//Password    string      `json:"password" db:"password_hash" binding:"required"`
	FirstName       string `json:"first_name" db:"first_name"`
	SecondName      string `json:"second_name" db:"second_name"`
	ThirdName       string `json:"third_name" db:"third_name"`
	UserRole        string `json:"user_role" db:"user_role"`
	Status          string `json:"status" db:"status"`
	SpecialtyIdList []int  `json:"specialty_id_list"`
}

type UserForCreate struct {
	Id              int    `json:"id" db:"id"`
	Email           string `json:"email" db:"email" binding:"required"`
	Password        string `json:"password" db:"password_hash" binding:"required"`
	FirstName       string `json:"first_name" db:"first_name"`
	SecondName      string `json:"second_name" db:"second_name"`
	ThirdName       string `json:"third_name" db:"third_name"`
	UserRole        string `json:"user_role" db:"user_role"`
	Status          string `json:"status" db:"status"`
	SpecialtyIdList []int  `json:"specialty_id_list"`
}

type UserDataFromToken struct {
	Id   int
	Role string
}

const (
	UserRole    = "user"
	AdminRole   = "admin"
	ManagerRole = "manager"
)
