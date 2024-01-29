package container

func Set(key string, value any) {
	_container.Store(key, value)
}

func Get(key string) any {
	value, ok := _container.Load(key)
	if !ok {
		return nil
	}

	return value
}

func Type[T any](key string) *T {
	value, ok := _container.Load(key)
	if !ok {
		return nil
	}

	x := T(value)
	return &x
}
