package services

import (
	"kood/social-network/pkg/api"
	"regexp"
	"strings"
	"time"
)

type ValidatorFunc func(interface{}) []api.ValidationError

var Validators = map[string]ValidatorFunc{
	"registration": validateRegistration,
	"post":         validatePost,
	"comment":      validateComment,
	"group":        validateGroup,
	"event":        validateEvent,
}

func ValidateOperation(operationType string, data interface{}) []api.ValidationError {
	validator, ok := Validators[operationType]
	if !ok {
		return []api.ValidationError{{Field: "", Message: "Unsupported operation type"}}
	}

	return validator(data)
}

func containsWhitespace(s string) bool {
	return strings.Contains(s, " ")
}

func validateRegistration(data interface{}) []api.ValidationError {
	registrForm, ok := data.(api.RegistrationRequest)
	if !ok {
		return []api.ValidationError{{Field: "", Message: "Invalid data type for registration"}}
	}

	var validationErrors []api.ValidationError

	if len(registrForm.Nickname) < 3 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "nickname",
			Message: "Nickname must be at least 3 characters long",
		})
	}
	if containsWhitespace(registrForm.Nickname) {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "nickname",
			Message: "Nickname must not contain whitespace",
		})
	}
	if !isValidEmail(registrForm.Email) {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "email",
			Message: "Invalid email format",
		})
	}
	if !isValidPassword(registrForm.Password) {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "password",
			Message: "Password must be at least 6 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character",
		})
	}
	if registrForm.DateOfBirth == "" {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "age",
			Message: "DateOfBirth is required",
		})
	} else {
		if !isValidDate(registrForm.DateOfBirth) {
			validationErrors = append(validationErrors, api.ValidationError{
				Field:   "dateOfBirth",
				Message: "Invalid date format. Date must be in dd.mm.yyyy format",
			})
		} else {
			ageDate, _ := time.Parse("02.01.2006", registrForm.DateOfBirth)
			today := time.Now().UTC()
			if ageDate.After(today) {
				validationErrors = append(validationErrors, api.ValidationError{
					Field:   "age",
					Message: "Date of birth cannot be in the future",
				})
			}
		}
	}
	if registrForm.FirstName == "" {
		message := "First name is required"
		if containsWhitespace(registrForm.FirstName) {
			message += " and must not contain whitespace"
		}
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "firstName",
			Message: message,
		})
	}
	if registrForm.LastName == "" {
		message := "Last name is required"
		if containsWhitespace(registrForm.LastName) {
			message += " and must not contain whitespace"
		}
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "lastName",
			Message: message,
		})
	}
	registrForm.AboutMe = strings.TrimSpace(registrForm.AboutMe)
	if len(registrForm.AboutMe) > 256 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "aboutMe",
			Message: "Content length cannot exceed 256 characters",
		})
	}

	return validationErrors
}

func isValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case strings.ContainsAny(string(char), "!@#$%^&*()-_=+{}[]|/<>?,."):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-01", dateStr)
	return err == nil
}

func validatePost(data interface{}) []api.ValidationError {
	postData, ok := data.(api.PostCreateRequest)
	if !ok {
		return []api.ValidationError{{Field: "", Message: "Invalid data type for post"}}
	}

	var validationErrors []api.ValidationError

	if len(postData.Title) < 6 || len(postData.Title) > 48 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "title",
			Message: "Title length must be between 6 and 48 characters",
		})
	}

	postData.Content = strings.TrimSpace(postData.Content)

	if len(postData.Content) > 256 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "content",
			Message: "Content length cannot exceed 256 characters",
		})
	}

	return validationErrors
}


func validateComment(data interface{}) []api.ValidationError {
	commentData, ok := data.(api.CommentCreateRequest)
	if !ok {
		return []api.ValidationError{{Field: "", Message: "Invalid data type for comment"}}
	}

	var validationErrors []api.ValidationError

	commentData.Content = strings.TrimSpace(commentData.Content)

	if len(commentData.Content) > 256 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "content",
			Message: "Content length cannot exceed 256 characters",
		})
	}

	return validationErrors
}


func validateGroup(data interface{}) []api.ValidationError {
	groupData, ok := data.(api.GroupCreateRequest)
	if !ok {
		return []api.ValidationError{{Field: "", Message: "Invalid data type for group"}}
	}

	var validationErrors []api.ValidationError

	if len(groupData.Title) < 6 || len(groupData.Title) > 48 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "title",
			Message: "Title length must be between 6 and 48 characters",
		})
	}

	groupData.Description = strings.TrimSpace(groupData.Description)

	if len(groupData.Description) > 256 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "description",
			Message: "Description length cannot exceed 256 characters",
		})
	}

	return validationErrors
}


func validateEvent(data interface{}) []api.ValidationError {
	eventData, ok := data.(api.GroupEventCreate)
	if !ok {
		return []api.ValidationError{{Field: "", Message: "Invalid data type for event"}}
	}

	var validationErrors []api.ValidationError

	if len(eventData.Title) < 6 || len(eventData.Title) > 48 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "title",
			Message: "Title length must be between 6 and 48 characters",
		})
	}

	eventData.Description = strings.TrimSpace(eventData.Description)

	if len(eventData.Description) > 256 {
		validationErrors = append(validationErrors, api.ValidationError{
			Field:   "description",
			Message: "Description length cannot exceed 256 characters",
		})
	}

	return validationErrors
}