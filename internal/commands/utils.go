package commands

func To[T any](t T) *T {
	return &t
}