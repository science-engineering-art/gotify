package persistence

type MetadataStorage struct {
}

func NewMetadataStorage() *MetadataStorage {
	s := &MetadataStorage{}
	return s
}

func (*MetadataStorage) Create(key []byte, data *[]byte) error {
	return nil
}

func (*MetadataStorage) Read(key []byte) (data *[]byte, err error) {
	return nil, nil
}

func (*MetadataStorage) Delete(key []byte) error {
	return nil
}
