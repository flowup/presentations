package auth

import (
	"testing"

	"github.com/flowup/presentations/go_mocks/auth"
	"github.com/flowup/presentations/go_mocks/email"
	"github.com/golang/mock/gomock"
)

// test valid password
func TestRegisterConfirmationEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockEmail := email.NewMockSender(mockCtrl)
	authService := &auth.AuthService{EmailSender: mockEmail}

	mockEmail.EXPECT().SendEmail("my@email.com", "peter.malina@flowup.cz", "Welcome!")
	authService.Register("peter.malina@flowup.cz", "12345678")
}
