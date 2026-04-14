package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "yourmomsbestpicture"
	pathname := CASPathTransformFunc(key)

	expectedPathName := "91fa3/16fe1/2e57b/75539/42e90/1fef8/83dd9/515b6"

	if pathname != expectedPathName {
		t.Errorf("have %s want %s", pathname, expectedPathName)
	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeSteam("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
