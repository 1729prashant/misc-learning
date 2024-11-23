package main
import "errors"

func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	} else if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	} else {
		return nil		
	}
}

