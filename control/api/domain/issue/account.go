package issue

type AccountNotFound struct{}

func (e *AccountNotFound) Error() string {
	return "Account not found"
}

type AccountAlreadyActivated struct{}

func (e *AccountAlreadyActivated) Error() string {
	return "Account already activated"
}

type AccountEmailInUse struct{}

func (e *AccountEmailInUse) Error() string {
	return "The provided email is already associated with an account"
}

type AccountInvalidOTP struct{}

func (e *AccountInvalidOTP) Error() string {
	return "Invalid OTP code"
}

type AccountInvalidCredentials struct{}

func (e *AccountInvalidCredentials) Error() string {
	return "Invalid email or password"
}

type AccountNotVerified struct{}

func (e *AccountNotVerified) Error() string {
	return "Account validation is required"
}

type AccountDeactivated struct{}

func (e *AccountDeactivated) Error() string {
	return "Account is deactivated"
}
