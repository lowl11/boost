package destroyer

func (destroyer *Destroyer) AddFunction(destroyFunc func()) *Destroyer {
	destroyer.mutex.Lock()
	defer destroyer.mutex.Unlock()

	destroyer.functions = append(destroyer.functions, destroyFunc)
	return destroyer
}

func (destroyer *Destroyer) Destroy() {
	destroyer.mutex.Lock()
	defer destroyer.mutex.Unlock()

	for _, destroyFunc := range destroyer.functions {
		destroyer.runFunc(destroyFunc)
	}
}
