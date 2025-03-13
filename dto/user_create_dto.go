package dto

import (
	"fmt"
	"time"
)

type UserCreateDTO struct {
	FirstName string     `json:"first_name" validate:"required,min=4,max=15"`
	LastName  string     `json:"last_name"  validate:"required,min=4,max=15"`
	Email     string     `json:"email"  validate:"required,email"`
	Phone     string     `json:"phone"  validate:"required"`
	DOB       CustomDate `json:"dob" validate:"required"`
	NIC       string     `json:"nic"  validate:"required"`
	Password  string     `json:"password"  validate:"required"`
	Role      string     `json:"role"  validate:"required"`
}

type CustomDate time.Time

const customDateLayout = "2006-01-02"

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	// Remove quotes around the date string
	str = str[1 : len(str)-1]

	parsedTime, err := time.Parse(customDateLayout, str)
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}

	*cd = CustomDate(parsedTime)
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(cd).Format(customDateLayout))), nil
}

func (cd CustomDate) ToTime() time.Time {
	return time.Time(cd)
}
