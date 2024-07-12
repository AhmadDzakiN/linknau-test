package questions

import (
	"context"
	"sync"
)

// Use Person struct from struct.go file
type PersonService interface {
	SavePerson(ctx context.Context, person Person) error
}

type EmailService interface {
	SendRegisterEmail(ctx context.Context, email string) error
}

type AgeValidatorService interface {
	ValidateAge(age uint) error
}

type Registrar struct {
	PersonService PersonService
	EmailService  EmailService
	AgeValidator  AgeValidatorService
}

// RegisterPerson simply will validate Person age. Then save the Person and will send a register email
func (r *Registrar) RegisterPerson(ctx context.Context, person Person) error {
	if err := r.AgeValidator.ValidateAge(person.Age); err != nil {
		return err
	}

	if err := r.PersonService.SavePerson(ctx, person); err != nil {
		return err
	}

	var wg sync.WaitGroup
	var emailErr error
	wg.Add(1)

	go func() {
		defer wg.Done()
		emailErr = r.EmailService.SendRegisterEmail(ctx, "anyemail@email.com")
	}()

	wg.Wait()

	if emailErr != nil {
		return emailErr
	}

	return nil
}
