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

	expectedOriginalKey := "91fa316fe12e57b7553942e901fef883dd9515b6"
	expectedPathName := "91fa3/16fe1/2e57b/75539/42e90/1fef8/83dd9/515b6"

	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("have %s want %s", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")

	if err := s.writeSteam(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "momsspecials"
	data := []byte("some jpg bytes")

	if err := s.writeSteam(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
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

	// if err := s.Delete(key); err != nil {
	// 	t.Error(err)
	// }
}
