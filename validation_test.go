package golang_validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T)  {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T)  {
	validate := validator.New()
	user := "fahril"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariable(t *testing.T)  {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T)  {
	validate := validator.New()
	user := "12345"

	err := validate.Var(user, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T)  {
	validate := validator.New()
	user := "999999"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T)  {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "fahril@gmail.com",
		Password: "fahril",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T)  {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "abu",
		Password: "abu",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossField(t *testing.T)  {
	type RegisterUser struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	request := RegisterUser{
		Username: "fahril@gmail.com",
		Password: "123456",
		ConfirmPassword: "123456",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T)  {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id string `validate:"required"`
		Name string `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	request := User{
		Id: "",
		Name: "",
		Address: Address{
			City: "",
			Country: "",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T)  {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id string `validate:"required"`
		Name string `validate:"required"`
		Addresses []Address `validate:"required,dive"` // isi dari collection harus divalidasi juga (dive)
	}

	validate := validator.New()
	request := User{
		Id: "",
		Name: "",
		Addresses: []Address{
			{
				City: "",
				Country: "",
			},
			{
				City: "",
				Country: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollection(t *testing.T)  {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id string `validate:"required"`
		Name string `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies []string `validate:"dive,required,min=3"`
	}

	validate := validator.New()
	request := User{
		Id: "",
		Name: "",
		Addresses: []Address{
			{
				City: "",
				Country: "",
			},
			{
				City: "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T)  {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id string `validate:"required"`
		Name string `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies []string `validate:"dive,required,min=3"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()
	request := User{
		Id: "",
		Name: "",
		Addresses: []Address{
			{
				City: "",
				Country: "",
			},
			{
				City: "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD" : {
				Name: "SD Indonesia",
			},
			"SMP" : {
				Name: "",
			},
			"" : {
				Name: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMap(t *testing.T)  {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id string `validate:"required"`
		Name string `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies []string `validate:"dive,required,min=3"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallet map[string]int `validate:"dive,keys,required,endkeys,required,gt=0"`
	}

	validate := validator.New()
	request := User{
		Id: "",
		Name: "",
		Addresses: []Address{
			{
				City: "",
				Country: "",
			},
			{
				City: "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD" : {
				Name: "SD Indonesia",
			},
			"SMP" : {
				Name: "",
			},
			"" : {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA" : 1000000,
			"MANDIRI" : 0,
			"" : 1000,
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAlias(t *testing.T)  {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id string `validate:"varchar,min=5"`
		Name string `validate:"varchar"`
		Owner string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id: "123",
		Name: "",
		Owner: "",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if  value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidationFunction(t *testing.T)  {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "FAHRIL",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidationParameter(t *testing.T)  {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "0823232823828",
		Pin: "123456",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T)  {
	type Login struct{
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := Login{
		Username: "fahril@gmail.com",
		Password: "fahril",
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T)  {
	validate := validator.New()
	validate.RegisterValidation("fields_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,fields_equals_ignore_case=Email|fields_equals_ignore_case=Phone"`
		Email string `validate:"required,email"`
		Phone string `validate:"required,numeric"`
		Name string `validate:"required"`
	}

	user := User{
		Username: "fahril@gmail.com",
		Email: "fahril@gmail.com",
		Phone: "2343536",
		Name: "Fahril",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email string `validate:"required,email"`
	Phone string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel)  {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// sukses
	} else {
		// gagal
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T)  {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "089283928392",
		Email: "fahril@gmail.com",
		Phone: "089283928392",
		Password: "rahasia",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}