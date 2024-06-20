package helpers

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidatePhone(phone string) error {
	re := regexp.MustCompile(`^[+][9][9][8]\d{9}$`)
	if !re.MatchString(phone) {
		return errors.New("invalid phone number")
	}

	return nil
}

func Validatetype(types string) error {

	var validTypes = map[string]bool{
		"self_pickup": true,
		"delivery":    true,
	}
	if !validTypes[types] {
		return errors.New("invalid type, you have to input 'self_pickup' or 'delivery'")
	}
	return nil
}

func Validatepayment_type(payment_type string) error {
	if payment_type != "uzum" && payment_type != "cash" && payment_type != "terminal" {

		return errors.New("invalid payment_type  three payment_type 'uzum','cash' and 'terminal'  ")
	}
	return nil
}

func Validstatus(types string) error {

	var validstatus = []string{"waiting_for_payment", "collecting", "delivery", "waiting_on_branch", "finished", "cancelled"}

	for _, validstatus := range validstatus {
		if types == validstatus {
			return nil
		}
	}
	return errors.New("invalid type, you have to input 'waiting_for_payment, collecting, delivery, waiting_on_branch, finished, cancelled' ")
}

// Define valid statuses
var validStatuses = []string{"waiting_for_payment", "collecting", "delivery", "waiting_on_branch", "finished", "cancelled"}

// Define valid transitions
var validTransitions = map[string][]string{
	"waiting_for_payment": {"collecting"},
	"collecting":          {"delivery"},
	"delivery":            {"waiting_on_branch"},
	"waiting_on_branch":   {"finished", "cancelled"},
	"finished":            {},
	"cancelled":           {},
}

// Function to validate status transition
func Validstatusorder(currentStatus, newStatus string) error {
	// Check if the new status is valid
	isValidStatus := false
	for _, status := range validStatuses {
		if newStatus == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		return errors.New("invalid type, you have to input 'waiting_for_payment, collecting, delivery, waiting_on_branch, finished, cancelled'")
	}

	// Check if the transition is valid
	validNextStatuses, exists := validTransitions[currentStatus]
	if !exists {
		return errors.New("invalid current status")
	}

	for _, status := range validNextStatuses {
		if newStatus == status {
			return nil
		}
	}

	return fmt.Errorf(fmt.Sprintf("invalid transition from '%s' to '%s'", currentStatus, newStatus))
}
