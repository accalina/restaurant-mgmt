package exception

type NotFoundError struct {
	message string
}

func (notFoundError NotFoundError) Error() string {
	return notFoundError.message
}
