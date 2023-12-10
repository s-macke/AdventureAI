package backend

type ChatBackend interface {
	GetResponse(input string) (string, int, int)
}
