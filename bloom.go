package bloom

import ("crypto/sha512")

type hash [64]uint8

type bloomfilter struct {
	hbytes hash
	data   []hash
}

func New() *bloomfilter {
	bf := new(bloomfilter)
	bf.data = make([]hash, 1)
	return bf
}

func (b *bloomfilter) In(in string) bool {
	h := sha512.New()
	h.Write([]byte(in))
	hashslice := h.Sum()

	if len(hashslice) != 64 {
		return false
	}
	var hsh hash
	for i := 0; i < 64; i++ {
		hsh[i] = hashslice[i]
	}

	if !equals(or(b.hbytes, hsh), b.hbytes) {
		return false
	}
	for i := 0; i < len(b.data); i++ {
		if equals(b.data[i], hsh) {
			return true
		}
	}
	return false
}

func (b *bloomfilter) Add(in string) {
	h := sha512.New()
	h.Write([]uint8(in))
        hashslice := h.Sum()
	
	if len(hashslice) != 64 {
		return
	}
	var hsh hash
	for i := 0; i < 64; i++ {
		hsh[i] = hashslice[i]
	}

	b.hbytes = or(b.hbytes, hsh)
	b.data = append(b.data, hsh)
}

func or(first hash, second hash) hash {
	var ored hash
	for i := 0; i < len(ored); i++ {
		ored[i] = first[i] | second[i]
	}
	return ored
}

func equals(first hash, second hash) bool {
	for i := 0; i < 64; i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
