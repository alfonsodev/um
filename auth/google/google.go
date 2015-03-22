package usermanager

func GoogleAuthLogic(code string) string {
	accessToken, idToken, err := googleAuth.Exchange(code)
	gplusID, err := googleAuth.DecodeIdToken(idToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n ADMINS:%s=%s\n", os.Getenv("ADMINS"), gplusID)
	if !strings.Contains(os.Getenv("ADMINS"), gplusID) {
		return ""
	} else {
		user, _ := GetByGoogleId(gplusID)
		if user.Id == "" {
			person := getPersonFromToken(accessToken)
			user := NewUserModel()
			user.Doc.Person = *person
			user.Save()
		}
		return gplusID
	}
}
