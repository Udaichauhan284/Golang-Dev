package auth

//now for getting the session, i want it, private, so make it in small letters

//This is Private, only access in this file and in this package, not outisde
func extractSession() string {
	return "still logged it"
}

//This is public can be used outside of package
func GetSession() string {
	return extractSession();
}