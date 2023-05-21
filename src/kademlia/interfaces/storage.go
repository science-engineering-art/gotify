package interfaces

type Persistence interface {
	Create(data *[]byte) (key []byte)

	Read(key []byte) (data *[]byte)

	Delete(key []byte) error
}
