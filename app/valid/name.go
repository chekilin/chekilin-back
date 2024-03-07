package valid

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func IsNameStartingWithBigLetter(fl validator.FieldLevel) bool { //引数の型、返り値は固定
	Name := fl.Field().String()

	pattern := `(\w+)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	matches := re.FindAllString(Name, -1)
	fmt.Println("Matches found:", matches)

	if matches != nil {
		return true
	}
	return false
}
