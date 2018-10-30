package validator

import "goleaky/repo"

func ValidateAndSave(report string) {
	valid := validate(report)
	if valid == true {
		repo.SaveReport(report)
	}
}

func validate(report string) bool {
	return true
}
