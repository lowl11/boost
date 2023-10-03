package container_helper

func Remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
