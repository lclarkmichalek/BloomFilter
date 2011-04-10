package bloom

import (
	"testing"
	"rand"
)

const (
	STORE_RUNS int = 200
	RETRIEVE_RUNS int = 200
)

func TestStore(t *testing.T) {
	bloom := New()

	for i := 0; i < STORE_RUNS; i++ {
		bloom.Add(string(rand.Uint32()))
	}
}

func TestRetrieveTrue(t *testing.T) {
	bloom := New()
	
	var toAdd [RETRIEVE_RUNS]string
	for i := 0; i < RETRIEVE_RUNS; i++ {
		toAdd[i] = string(rand.Uint32())
		bloom.Add(toAdd[i])
	}
	for i := 0; i < RETRIEVE_RUNS; i++ {
		if !bloom.In(toAdd[i]) {
			t.Error("BloomFilter.In returns false posative for string ", toAdd[i])
		}
	}
}

func TestRetrieveFalse(t *testing.T) {
	bloom := New()
	
	var toTest [RETRIEVE_RUNS]string
	for i := 0; i < RETRIEVE_RUNS; i++ {
		toTest[i] = string(rand.Uint32())
	}

	for i := 0; i < RETRIEVE_RUNS; i++ {
		if bloom.In(toTest[i]) {
			t.Error("BloomFilter.In returns false negative for string ", toTest[i])
		}
	}
}