package dto

import "time"

type UserGetDTO struct {
	ID         uint      `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	DOB        time.Time `json:"dob"`
	NIC        string    `json:"nic"`
	EmployeeID string    `json:"employee_id"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserUpdateDTO struct {
	FirstName string     `json:"first_name" binding:"required,min=4,max=15"`
	LastName  string     `json:"last_name"  binding:"required,min=4,max=15"`
	Email     string     `json:"email"  binding:"required,email"`
	Phone     string     `json:"phone"  binding:"required"`
	DOB       CustomDate `json:"dob" binding:"required"`
	NIC       string     `json:"nic"  binding:"required"`
	Role      string     `json:"role"  binding:"required"`
}

type UsersGetDTO struct {
	Users []UserGetDTO `json:"users"`
}
