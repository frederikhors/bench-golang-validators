package main

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-playground/validator/v10"
	buffalo "github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	saddam "github.com/thedevsaddam/govalidator"
)

type User struct {
	FirstName string `json:"firstname" valid:"required" validate:"required"`
	LastName  string `json:"lastname" valid:"required" validate:"required"`
	Age       int    `json:"age" valid:"range(0|130)" validate:"gte=0,lte=130"`
	Email     string `json:"email" valid:"required,email" validate:"required,email"`
	Street    string `json:"street" valid:"required" validate:"required"`
	City      string `json:"city" valid:"required" validate:"required"`
	Planet    string `json:"planet" valid:"required" validate:"required"`
	Phone     string `json:"phone" valid:"required" validate:"required"`
}

func main() {
	user := newUser()
	playgroundValidate := validator.New()

	validateWithAsaskevich(user)
	validateWithPlayground(playgroundValidate, user)
	validateWithSaddam(user)
	validateWithBuffalo(user)
	validateWithOzzo(user)
}

func newUser() *User {
	return &User{
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       80,
		City:      "string",
		Email:     "Badger.Smith@email.com",
		Street:    "Eavesdown Docks",
		Planet:    "Persphone",
		Phone:     "none",
	}
}

func validateWithAsaskevich(user *User) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		println("Asaskevich: " + err.Error())
	}
}

func validateWithPlayground(validate *validator.Validate, user *User) {
	if err := validate.Struct(user); err != nil {
		println("Playground: " + err.Error())
	}
}

func validateWithSaddam(user *User) {
	rules := saddam.MapData{
		"firstname": []string{"required"},
		"lastname":  []string{"required"},
		"age":       []string{"min:0", "max:130"},
		"email":     []string{"required", "email"},
		"street":    []string{"required"},
		"city":      []string{"required"},
		"planet":    []string{"required"},
		"phone":     []string{"required"},
	}
	opts := saddam.Options{
		Data:  user,
		Rules: rules,
	}
	v := saddam.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		data, _ := json.Marshal(e)
		println("Saddam: " + string(data))
	}
}

func validateWithBuffalo(user *User) {
	errors := buffalo.Validate(
		&validators.StringIsPresent{Field: user.FirstName},
		&validators.StringIsPresent{Field: user.LastName},
		&validators.IntIsGreaterThan{Field: user.Age, Compared: 0},
		&validators.IntIsLessThan{Field: user.Age, Compared: 130},
		&validators.EmailIsPresent{Field: user.Email},
		&validators.StringIsPresent{Field: user.Street},
		&validators.StringIsPresent{Field: user.City},
		&validators.StringIsPresent{Field: user.Planet},
		&validators.StringIsPresent{Field: user.Phone},
	)
	if len(errors.Errors) > 0 {
		println(errors.Error())
	}
}

func validateWithOzzo(user *User) {
	err := validation.ValidateStruct(user,
		validation.Field(&user.FirstName, validation.Required),
		validation.Field(&user.LastName, validation.Required),
		validation.Field(&user.Age, validation.Min(0), validation.Max(130)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Street, validation.Required),
		validation.Field(&user.City, validation.Required),
		validation.Field(&user.Planet, validation.Required),
		validation.Field(&user.Phone, validation.Required),
	)
	if err != nil {
		println("Ozzo: " + err.Error())
	}
}
