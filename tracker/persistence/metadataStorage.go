package persistence

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
)

type MetadataStorage struct {
	KV map[string][][]byte
}

func NewMetadataStorage() *MetadataStorage {
	s := &MetadataStorage{}
	s.KV = make(map[string][][]byte)
	return s
}

func (s *MetadataStorage) Create(key []byte, data *[]byte) error {
	id := base64.RawStdEncoding.EncodeToString(key)

	// _, exists := s.KV[id]
	// if !exists {
	// }

	s.KV[id] = append(s.KV[id], *data)
	//fmt.Println(s.KV[id])
	return nil
}

func (s *MetadataStorage) Read(key []byte, start int32, end int32) (data *[]byte, err error) {
	id := base64.RawStdEncoding.EncodeToString(key)

	v, exists := s.KV[id]
	if !exists {
		return nil, errors.New("the key is not found")
	}
	fmt.Println("Retrived v:", v)

	flattenByteArray := getFlattenByteArray(v)
	fmt.Println("Flatten array:", flattenByteArray)
	return &flattenByteArray, nil
}

func (s *MetadataStorage) Delete(key []byte) error {
	id := base64.RawStdEncoding.EncodeToString(key)

	_, exists := s.KV[id]
	if !exists {
		return errors.New("the key is not found")
	}

	delete(s.KV, id)
	return nil
}

func getFlattenByteArray(data [][]byte) []byte {
	flatByteArray := []byte{}

	for _, elem := range data {
		elemLen := len(elem)
		byteLen := make([]byte, 4)
		binary.LittleEndian.PutUint32(byteLen, uint32(elemLen))

		flatByteArray = append(flatByteArray, byteLen...)
		flatByteArray = append(flatByteArray, elem...)
	}

	return flatByteArray
}
