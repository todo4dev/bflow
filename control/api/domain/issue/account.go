package issue

type AccountAlreadyActivated struct{}

func (e *AccountAlreadyActivated) Error() string {
	return "Account already activated"
}

type AccountNotFound struct{}

func (e *AccountNotFound) Error() string {
	return "Account not found"
}

type AccountEmailInUse struct{}

func (e *AccountEmailInUse) Error() string {
	return "The provided email is already associated with an account"
}

type AccountInvalidOTP struct{}

func (e *AccountInvalidOTP) Error() string {
	return "Invalid OTP code"
}
