package helpers

func CheckErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err.Error())
		}
	}
}