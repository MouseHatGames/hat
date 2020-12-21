package store

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type boltStore struct {
	db     *bolt.DB
	bucket []byte
}

const bucketName = "hatstore"

func NewStore(filePath string) (Store, error) {
	db, err := bolt.Open(filePath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	}); err != nil {
		db.Close()
		return nil, fmt.Errorf("create bucket: %w", err)
	}

	return &boltStore{db: db, bucket: []byte(bucketName)}, nil
}

func (s *boltStore) Close() error {
	return s.db.Close()
}

func (s *boltStore) Get(p Path) ([]byte, error) {
	var value []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		value = b.Get(joinPath(p))

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (s *boltStore) Set(p Path, v []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		return b.Put(joinPath(p), v)
	})
}

func (s *boltStore) Del(p Path) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		return b.Delete(joinPath(p))
	})
}
