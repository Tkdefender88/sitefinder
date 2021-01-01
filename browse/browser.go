package browse

type Browser interface {
	Browse(url string) error
}
