package interfaces

type Persistence interface {
	Create(key []byte, data *[]byte) error

	Read(key []byte) (data *[]byte)

	Delete(key []byte) error
}
