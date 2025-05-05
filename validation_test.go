package golang_validation

import (
	"fmt"
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