package kv

// PK a primary key
type PK string

// SK a secondary key
type SK string

type API interface {
	Get(pk PK, sk SK) ([]byte, error)
	Set(pk PK, sk SK, content []byte) error
	Delete(pk PK, sk SK) error
}
