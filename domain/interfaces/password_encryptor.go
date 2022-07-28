package interfaces

type Encryptor interface {
	Encrypt(value string) string
}
