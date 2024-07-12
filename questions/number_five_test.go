package questions

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// MockPersonService is a mock of the PersonService interface
type MockPersonService struct {
	SavePersonFunc func(ctx context.Context, person Person) error
}

func (m *MockPersonService) SavePerson(ctx context.Context, person Person) error {
	return m.SavePersonFunc(ctx, person)
}

// MockEmailService is a mock of the EmailService interface
type MockEmailService struct {
	SendWelcomeEmailFunc func(ctx context.Context, email string) error
}

func (m *MockEmailService) SendRegisterEmail(ctx context.Context, email string) error {
	return m.SendWelcomeEmailFunc(ctx, email)
}

// MockAgeValidator is a mock of the AgeValidator interface
type MockAgeValidator struct {
	ValidateAgeFunc func(age uint) error
}

func (m *MockAgeValidator) ValidateAge(age uint) error {
	return m.ValidateAgeFunc(age)
}

type NumberFiveTestSuite struct {
	suite.Suite
	mockPersonSvc    PersonService
	mockEmailSvc     EmailService
	mockAgeValidator AgeValidatorService
	registrar        Registrar
}

func TestNumberFiveSuite(t *testing.T) {
	suite.Run(t, new(NumberFiveTestSuite))
}

func (n *NumberFiveTestSuite) TestRegisterPerson() {
	type fields struct {
		mock func(ctx context.Context, person Person)
	}

	type args struct {
		ctx    context.Context
		person Person
	}

	tests := []struct {
		name        string
		args        args
		fields      fields
		registrar   Registrar
		expectedErr error
	}{
		{
			name: "invalid age",
			args: args{
				ctx:    context.Background(),
				person: Person{Name: "Ahmad", Age: 21},
			},
			fields: fields{func(ctx context.Context, person Person) {
				n.mockAgeValidator = &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return errors.New("invalid age")
					},
				}
				n.mockPersonSvc = &MockPersonService{}
				n.mockEmailSvc = &MockEmailService{}
			}},
			registrar: Registrar{
				PersonService: &MockPersonService{},
				EmailService:  &MockEmailService{},
				AgeValidator: &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return errors.New("invalid age")
					},
				},
			},
			expectedErr: errors.New("invalid age"),
		},
		{
			name: "error while saving person",
			args: args{
				ctx:    context.Background(),
				person: Person{Name: "Ahmad", Age: 21},
			},
			fields: fields{mock: func(ctx context.Context, person Person) {
				n.mockAgeValidator = &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return nil
					},
				}
				n.mockPersonSvc = &MockPersonService{
					SavePersonFunc: func(ctx context.Context, person Person) error {
						return errors.New("save person error")
					},
				}
				n.mockEmailSvc = &MockEmailService{}
			}},
			registrar: Registrar{
				PersonService: &MockPersonService{
					SavePersonFunc: func(ctx context.Context, person Person) error {
						return errors.New("save person error")
					},
				},
				EmailService: &MockEmailService{},
				AgeValidator: &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return nil
					},
				},
			},
			expectedErr: errors.New("save person error"),
		},
		{
			name: "error while sending register email",
			args: args{
				ctx:    context.Background(),
				person: Person{Name: "Ahmad", Age: 21},
			},
			fields: fields{func(ctx context.Context, person Person) {
				n.mockAgeValidator = &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return nil
					},
				}
				n.mockPersonSvc = &MockPersonService{
					SavePersonFunc: func(ctx context.Context, person Person) error {
						return nil
					},
				}
				n.mockEmailSvc = &MockEmailService{
					SendWelcomeEmailFunc: func(ctx context.Context, email string) error {
						return errors.New("email error")
					},
				}
			}},
			registrar: Registrar{
				PersonService: &MockPersonService{
					SavePersonFunc: func(ctx context.Context, person Person) error {
						return nil
					},
				},
				EmailService: &MockEmailService{
					SendWelcomeEmailFunc: func(ctx context.Context, email string) error {
						return errors.New("email error")
					},
				},
				AgeValidator: &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return nil
					},
				},
			},
			expectedErr: errors.New("email error"),
		},
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				person: Person{Name: "Ahmad", Age: 21},
			},
			fields: fields{func(ctx context.Context, person Person) {
				n.mockAgeValidator = &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return nil
					},
				}
				n.mockPersonSvc = &MockPersonService{
					SavePersonFunc: func(ctx context.Context, person Person) error {
						return nil
					},
				}
				n.mockEmailSvc = &MockEmailService{
					SendWelcomeEmailFunc: func(ctx context.Context, email string) error {
						return nil
					},
				}
			}},
			registrar: Registrar{
				PersonService: &MockPersonService{
					SavePersonFunc: func(ctx context.Context, person Person) error {
						return nil
					},
				},
				EmailService: &MockEmailService{
					SendWelcomeEmailFunc: func(ctx context.Context, email string) error {
						return nil
					},
				},
				AgeValidator: &MockAgeValidator{
					ValidateAgeFunc: func(age uint) error {
						return nil
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		n.Suite.Run(test.name, func() {
			test.fields.mock(test.args.ctx, test.args.person)

			actualErr := test.registrar.RegisterPerson(test.args.ctx, test.args.person)

			assert.Equal(n.T(), test.expectedErr, actualErr)
		})
	}
}
