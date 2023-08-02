package password

func Compare(password, rePassword string) error {
	if password == rePassword {
		return nil
	}

	return ErrorPasswordsNotEqual()
}
