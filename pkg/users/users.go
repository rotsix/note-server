package users

// Login returns a token, empty in case of fail
func Login(user, password string) string {
	if user == "" || password == "" {
		return ""
	}

	return "t0k3n"
}
