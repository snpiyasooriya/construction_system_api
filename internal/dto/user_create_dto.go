package dto

import (
	"fmt"
	"time"
)

type Date time.Time

// MarshalJSON converts the Date to a string in the "YYYY-MM-DD" format
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(d).Format("2006-01-02"))), nil
}

// UnmarshalJSON parses a string in the "YYYY-MM-DD" format to a Date
func (d *Date) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	*d = Date(parsedTime)
	return nil
}

// ToTime converts the custom Date type back to a time.Time
func (d Date) ToTime() time.Time {
	return time.Time(d)
}

type UserCreateDTO struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	DOB       time.Time `json:"dob"`
	NIC       string    `json:"nic"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
}
