package panicer

func fromAny(err any) string {
	if err == nil {
		return ""
	}

	switch err.(type) {
	case string:
		return err.(string)
	case error:
		return err.(error).Error()
	}

	return ""
}
