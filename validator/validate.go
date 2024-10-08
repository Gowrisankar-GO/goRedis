// Package validator provides a straight forward way to validate basic payloads passed to every request handler functions
package validator

import (
	"fmt"
	"errors"
	"redis_user_management/info"
	"redis_user_management/models"
	"reflect"
	"regexp"
	"strings"
)

// 	Function ValidateStruct used to validate struct typed variables.It returns error if the struct doesnot meets the required validations, otherwise it retuns nil.
func ValidateStruct(data interface{}) error {

	v := reflect.ValueOf(data)

	// confirming whether the provided data is a struct type or not
	if v.Kind() != reflect.Struct {

		return info.ErrStructType
	}

	t := reflect.TypeOf(data)

	for i := 0; i < v.NumField(); i++ {

		fieldValue := v.Field(i)

		FieldType := t.Field(i)

		FieldName := FieldType.Name

		tagValue := FieldType.Tag.Get(models.ValidatorKeys["Tag"])

		tagData := strings.Split(tagValue, ",")

		for _, key := range tagData {

			switch key {

			case models.ValidatorKeys["Required"]:

				emptyData := fmt.Sprintf("%v", fieldValue.Interface())

				if emptyData == "" || emptyData == "0" || emptyData == "nil" {

					err := fmt.Sprintf("Field:%v, Required:%v",FieldName,info.ErrReqFld)

					return errors.New(err)
				}

			case models.ValidatorKeys["Omit"]:

				continue

			case models.ValidatorKeys["FirstName"]:

				pattern := `^[A-Za-z]{3,15}$`

				regex := regexp.MustCompile(pattern)

				if !regex.MatchString(fieldValue.Interface().(string)){

					err := fmt.Sprintf("Field:%v, FirstName:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				}

			case models.ValidatorKeys["LastName"]:

				pattern := `^[A-Za-z]{1,10}$`

				regex := regexp.MustCompile(pattern)

				if !regex.MatchString(fieldValue.String()){

					err := fmt.Sprintf("Field:%v, LastName:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				}

			case models.ValidatorKeys["Email"]:

				pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2,})*$`

				regex := regexp.MustCompile(pattern)

				if !regex.MatchString(fieldValue.String()){

					err := fmt.Sprintf("Field:%v, Email:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				}

			case models.ValidatorKeys["Mobile"]:

				pattern := `^[6-9]\d{9}$`

				regex := regexp.MustCompile(pattern)

				if !regex.MatchString(fieldValue.String()){

					err := fmt.Sprintf("Field:%v, Email:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				}

			case models.ValidatorKeys["Age"]:

				if fieldValue.Interface().(uint)<1 || fieldValue.Interface().(uint)>130{

					err := fmt.Sprintf("Field:%v, Age:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				} 

			case models.ValidatorKeys["Role"]:

				if fieldValue.Interface().(uint)<1{

					err := fmt.Sprintf("Field:%v, Role:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				} 

			case models.ValidatorKeys["Status"]:

				if fieldValue.Interface().(uint)!=0 && fieldValue.Interface().(uint)!=1{

					err := fmt.Sprintf("Field:%v, Status:%v",FieldName,info.ErrFldVal)

					return errors.New(err)
				} 

			default:

				fmt.Printf("the field: %v  has validation errors",FieldName)

			}

		}

	}

	return nil

}
