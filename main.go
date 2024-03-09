package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Hello struct {
	Name  string `validation:"not_empty"`
	Email string `validation:"not_empty"`
}

func main() {
	hello := Hello{
		Name:  "M. Raj",
		Email: "raj@gmail.com",
	}

	err := validate(&hello)

	if len(err) > 0 {
		for _, er := range err {
			fmt.Println(er.Error())
		}
	} else {
		fmt.Println("Validation success.")
	}
}

type Validation struct {
	function    func(string) bool
	err_message string
}

func validate(strct interface{}) []error {
	validations := map[string]Validation{
		"not_empty": {
			function: func(value string) bool {
				return value != ""
			},
			err_message: "Field Should Not Left Empty",
		},
	}

	num_of_fields := reflect.TypeOf(strct).Elem().NumField()
	errors := []error{}

	for i := 0; i < num_of_fields; i++ {
		field := reflect.TypeOf(strct).Elem().FieldByIndex([]int{i})
		value := reflect.ValueOf(strct).Elem().FieldByName(field.Name).String()
		if assertions, ok := field.Tag.Lookup("validation"); ok {
			validation_tags := strings.Split(assertions, ",")
			if len(validation_tags) > 0 {
				for _, v := range validation_tags {
					if !(validations[v].function(value)) {
						errors = append(errors, fmt.Errorf("%s : %s", validations[v].err_message, field.Name))
					}
				}
			}
		}
	}

	return errors
}
