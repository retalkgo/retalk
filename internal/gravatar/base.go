package gravatar

func GetGravatarURL(baseURL string, email string) string {
	return baseURL + email
}
