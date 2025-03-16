package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)

	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"

	if pathKey.PathName != expectedPathName {
		t.Errorf("expected %s, got %s", expectedPathName, pathKey.PathName)
	}

	if pathKey.FileName != expectedOriginalKey {
		t.Errorf("expected %s, got %s", expectedOriginalKey, pathKey.FileName)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "timepass"
	data := []byte("some data third version")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)

	fmt.Printf("Data in the file : %s", string(b))

	if string(b) != string(data) {
		t.Errorf("want %s, got %s", data, b)
	}
}
