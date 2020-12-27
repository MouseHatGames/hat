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

func (s *boltStore) Get(p Path) (string, error) {
	var value []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		value = b.Get(joinPath(p))

		return nil
	})
	if err != nil {
		return "", err
	}

	if value == nil {
		return "", ErrKeyNotFound
	}

	return string(value), nil
}

func (s *boltStore) GetBulk(p []Path) ([]*string, error) {
	values := make([]*string, len(p))

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		for i, path := range p {
			data := b.Get(joinPath(path))

			if data != nil {
				s := string(data)
				values[i] = &s
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (s *boltStore) Set(p Path, v string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		return b.Put(joinPath(p), []byte(v))
	})
}

func (s *boltStore) Del(p Path) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucket)

		return b.Delete(joinPath(p))
	})
}
