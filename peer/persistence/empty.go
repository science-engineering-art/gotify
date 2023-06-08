package persistence

type Empty struct {
}

func NewEmpty() *Empty {
	return &Empty{}
}

func (e *Empty) Create(key []byte, data *[]byte) error {
	return nil
}

func (e *Empty) Read(key []byte, start int32, end int32) (data *[]byte, err error) {
	return nil, nil
}

func (e *Empty) Delete(key []byte) error {
	return nil
}
