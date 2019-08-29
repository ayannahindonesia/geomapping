package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thedevsaddam/govalidator"
)

type AsiraValidator struct {
	DB *gorm.DB `json:"db"`
}

func (a *AsiraValidator) CustomValidatorRules() {
	// unique value on each column. format : []string{"unique:[table],[column],[exclude_column],[excluded_value]"}
	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		var (
			queryRow *gorm.DB
			total    int
		)

		query := `SELECT COUNT(*) as total FROM %s WHERE %s = ?`
		params := strings.Split(strings.TrimPrefix(rule, fmt.Sprintf("%s:", "unique")), ",")

		if len(params) == 2 {
			query = fmt.Sprintf(query, params[0], params[1])
			queryRow = a.DB.Raw(query, value)
		} else if len(params) == 4 {
			query += ` AND %s != ?`
			query = fmt.Sprintf(query, params[0], params[1], params[2])
			queryRow = a.DB.Raw(query, value, params[3])
		} else {
			return fmt.Errorf("Arguments not enough")
		}

		queryRow.Row().Scan(&total)

		if total > 0 {
			if message != "" {
				return errors.New(message)
			}

			return fmt.Errorf("The %s has already been taken", field)
		}

		return nil
	})

	// bank type. must be a listed bank type id.
	govalidator.AddCustomRule("bank_type_id", func(field string, rule string, message string, value interface{}) error {
		var (
			queryRow *gorm.DB
			total    int
		)

		query := `SELECT COUNT(*) as total FROM bank_types WHERE id = ?`
		queryRow = a.DB.Raw(query, value)
		queryRow.Row().Scan(&total)

		if total < 1 {
			return fmt.Errorf(fmt.Sprint("bank type %v is not found.", value), field)
		}
		return nil
	})

	// active / inactive string only.
	govalidator.AddCustomRule("active_inactive", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if strings.ToLower(val) != "active" && strings.ToLower(val) != "inactive" {
			return fmt.Errorf("The %s field must be contain word: active or inactive", field)
		}
		return nil
	})

	// validator for pagination
	govalidator.AddCustomRule("asc_desc", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if strings.ToUpper(val) != "ASC" && strings.ToUpper(val) != "DESC" {
			return fmt.Errorf("The %s field must be contain word: asc or desc", field)
		}
		return nil
	})

	// validator for loans
	govalidator.AddCustomRule("loan_statuses", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "approved" && val != "rejected" && val != "processing" {
			return fmt.Errorf("The %s field must be contain either: approved, rejected, or processing", field)
		}
		return nil
	})

	// validator for otp entity types
	govalidator.AddCustomRule("otp_entity_types", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "loan" && val != "borrower" {
			return fmt.Errorf("The %s field must be contain either: loan or borrower", field)
		}
		return nil
	})

	// validator for indonesia phone number
	govalidator.AddCustomRule("id_phonenumber", func(field string, rule string, message string, value interface{}) error {
		reg := regexp.MustCompile(`\+?([ -]?\d+)+|\(\d+\)([ -]\d+)`)
		if value == nil {
			return fmt.Errorf("no value")
		}
		val := value.(string)
		if !reg.MatchString(val) {
			return fmt.Errorf("The %s field is not a valid indonesia phone number", field)
		}
		return nil
	})
}
