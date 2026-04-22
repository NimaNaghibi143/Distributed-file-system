package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "yourmomsbestpicture"
	pathKey := CASPathTransformFunc(key)

	expectedFileName := "91fa316fe12e57b7553942e901fef883dd9515b6"
	expectedPathName := "91fa3/16fe1/2e57b/75539/42e90/1fef8/83dd9/515b6"

	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectedFileName {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFileName)
	}
}

// func TestStoreDeleteKey(t *testing.T) {
// 	opts := StoreOpts{
// 		PathTransformFunc: CASPathTransformFunc,
// 	}

// 	s := NewStore(opts)
// 	key := "momsspecials"
// 	data := []byte("some jpg bytes")

// 	if err := s.writeSteam(key, bytes.NewReader(data)); err != nil {
// 		t.Error(err)
// 	}

// 	if err := s.Delete(key); err != nil {
// 		t.Error(err)
// 	}

// }

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")

		if err := s.writeSteam(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}
		fmt.Println(string(b))

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("expected to NOT have key %s", key)
		}

	}

}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
