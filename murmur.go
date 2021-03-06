package hyperloglog

import "unsafe"

// This file implements the murmur3 32-bit hash on 32bit and 64bit integers
// for little endian machines only with no heap allocation.  If you are using
// HLL to count integer IDs on intel machines, this is your huckleberry.

// MurmurString implements a fast version of the murmur hash function for strings
// for little endian machines.  Suitable for adding strings to HLL counter.
func MurmurString(key string) uint32 {
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	var h, k uint32

	bkey := *(*[]byte)(unsafe.Pointer(&key))
	blen := len(bkey)
	l := blen / 4 // chunk length

	// encode each 4 byte chunk of `key'
	for i := 0; i < l; i++ {
		k = *(*uint32)(unsafe.Pointer(&bkey[i*4]))
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
		h = (h << 13) | (h >> (32 - 13))
		h = (h * 5) + 0xe6546b64
	}

	k = 0
	//remainder
	if mod := blen % 4; mod > 0 {
		btail := *(*uint32)(unsafe.Pointer(&bkey[l*4]))
		switch mod {
		case 3:
			k ^= ((btail >> 16) & 0x0000FFFF) << 16
			fallthrough
		case 2:
			k ^= ((btail >> 8) & 0x000000FF) << 8
			fallthrough
		case 1:
			k ^= btail & 0x000000FF
			k *= c1
			k = (k << 15) | (k >> (32 - 15))
			k *= c2
			h ^= k
		}
	}

	h ^= uint32(blen)
	h ^= (h >> 16)
	h *= 0x85ebca6b
	h ^= (h >> 13)
	h *= 0xc2b2ae35
	h ^= (h >> 16)
	return h
}

// Murmur32 implements a fast version of the murmur hash function for uint32 for
// little endian machines.  Suitable for adding 32bit integers to a HLL counter.
func Murmur32(i uint32) uint32 {
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	var h, k uint32
	k = i
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// second part
	h ^= 4
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

// Murmur64 implements a fast version of the murmur hash function for uint64 for
// little endian machines.  Suitable for adding 64bit integers to a HLL counter.
func Murmur64(i uint64) uint32 {
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	var h, k uint32
	//first 4-byte chunk
	k = uint32(i)
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// second 4-byte chunk
	k = uint32(i >> 32)
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// second part
	h ^= 8
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

// Murmur128 implements a fast version of the murmur hash function for two uint64s
// for little endian machines.  Suitable for adding a 128bit value to an HLL counter.
func Murmur128(i, j uint64) uint32 {
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	var h, k uint32
	//first 4-byte chunk
	k = uint32(i)
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// second 4-byte chunk
	k = uint32(i >> 32)
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// third 4-byte chunk
	k = uint32(j)
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// fourth 4-byte chunk
	k = uint32(j >> 32)
	k *= c1
	k = (k << 15) | (k >> (32 - 15))
	k *= c2
	h ^= k
	h = (h << 13) | (h >> (32 - 13))
	h = (h * 5) + 0xe6546b64
	// second part
	h ^= 16
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h

}
