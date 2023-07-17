package panicer

func Handle(err any) error {
	if err == nil {
		return nil
	}

	return fromAny(err)
}
