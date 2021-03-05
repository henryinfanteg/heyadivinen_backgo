package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// IsNotEmpty devuelve error si un string es vacio
func IsNotEmpty(data string, field string) error {
	if len(data) == 0 {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterThanCeroInt devuelve error si un numero es menor que cero
func IsGreaterThanCeroInt(number int, field string) error {
	if number <= 0 {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterThanCeroFloat64 devuelve error si un numero es menor que cero
func IsGreaterThanCeroFloat64(number float64, field string) error {
	if number <= 0 {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterOrEqualToCeroInt devuelve error si un numero es menor o igual a cero
func IsGreaterOrEqualToCeroInt(number int, field string) error {
	if number < 0 {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterOrEqualToCeroFloat64 devuelve error si un numero es menor o igual a cero
func IsGreaterOrEqualToCeroFloat64(number float64, field string) error {
	if number < 0 {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterThanInt devuelve error si numeroA es menor que numeroB
func IsGreaterThanInt(numberA int, numberB int, field string) error {
	if numberA <= numberB {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterThanFloat64 devuelve error si numeroA es menor que numeroB
func IsGreaterThanFloat64(numberA float64, numberB float64, field string) error {
	if numberA <= numberB {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterOrEqualToInt devuelve error si numeroA es menor o igual que numeroB
func IsGreaterOrEqualToInt(numberA int, numberB int, field string) error {
	if numberA < numberB {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsGreaterOrEqualToFloat64 devuelve error si numeroA es menor o igual que numeroB
func IsGreaterOrEqualToFloat64(numberA float64, numberB float64, field string) error {
	if numberA < numberB {
		resp := fmt.Sprintf("%s%s", field, " is required")
		return errors.New(resp)
	}
	return nil
}

// IsDateEmpty devuelve error si una fecha es vacia
func IsDateEmpty(data time.Time, field string) error {
	// if data == time.Time{} {
	// 	resp := fmt.Sprintf("%s%s", field, " is required")
	// 	return errors.New(resp)
	// }
	return nil
}

var validate *validator.Validate

// ErrorResponse entidad
type ErrorResponse struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

// ValidateStruct valida un struct
func ValidateStruct(object interface{}) error {

	if validate == nil {
		validate = validator.New()
	}

	if err := validate.Struct(object); err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var errs []ErrorResponse
		for _, err2 := range err.(validator.ValidationErrors) {
			switch err2.Tag() {
			case "required":
				// errs += fmt.Sprintf("'%s' %s\n", err2.StructNamespace(), "is required")
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: "is required",
				})
			case "min":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be greater than or equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be greater than or equal to", err2.Param()),
				})
			case "max":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be less than or equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be less than or equal to", err2.Param()),
				})
			case "gt":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be greater than", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be greater than", err2.Param()),
				})
			case "gte":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be greater than or equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be greater than or equal to", err2.Param()),
				})
			case "lt":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be less than", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be less than", err2.Param()),
				})
			case "lte":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be less than or equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be less than or equal to", err2.Param()),
				})
			case "len":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be equal to", err2.Param()),
				})
			case "eq":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be equal to", err2.Param()),
				})
			case "ne":
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), "must be not equal to", err2.Param())
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s %s", "must be not equal to", err2.Param()),
				})
			case "alphanum":
				// errs += fmt.Sprintf("'%s' %s\n", err2.StructNamespace(), " must contain alphanumeric characters only")
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: "must contain alphanumeric characters only",
				})
			case "numeric":
				// errs += fmt.Sprintf("'%s' %s\n", err2.StructNamespace(), " must contain numeric characters only")
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: "must contain numeric characters only",
				})
			case "email":
				// errs += fmt.Sprintf("'%s' %s\n", err2.StructNamespace(), "invalid email format")
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: "invalid email format",
				})
			default:
				// errs += fmt.Sprintf("'%s' %s %s\n", err2.StructNamespace(), " - ", err)
				errs = append(errs, ErrorResponse{
					Field:       err2.StructNamespace(),
					Description: fmt.Sprintf("%s", err),
				})
			}

		}
		arrayBytes, err := json.Marshal(errs)
		if err != nil {
			return err
		}

		return errors.New(string(arrayBytes))
	}

	return nil
}
