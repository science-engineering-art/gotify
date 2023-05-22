package structs

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type Storage struct {
	KV map[string]interface{}
}

func (s *Storage) init() {
	s.KV = make(map[string]interface{})
}

func (s *Storage) Create(key []byte, data *[]byte) error {
	id := base64.RawStdEncoding.EncodeToString(key)

	_, exists := s.KV[id]
	if exists {
		return errors.New("the key already exists")
	}

	s.KV[id] = data

	return nil
}

func (s *Storage) Read(key []byte) (*[]byte, error) {
	id := base64.RawStdEncoding.EncodeToString(key)

	v, exists := s.KV[id]
	if !exists {
		return nil, errors.New("the key is not found")
	}

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *Storage) Delete(key []byte) error {
	id := base64.RawStdEncoding.EncodeToString(key)

	_, exists := s.KV[id]
	if !exists {
		return errors.New("the key is not found")
	}

	delete(s.KV, id)
	return nil
}
