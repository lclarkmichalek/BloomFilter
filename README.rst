BloomFilter
===========

An implementation of the BloomFilter datastructure in Go.

Uses the BloomFilter, along with a hash list to check if an object is in the structure. The hash list is only checked when the BloomFilter indicates that the object is in the structure. This is due to the fact that BloomFilter's are vunreble to false posatives, but never false negatives.

The library exposes the following functions and methods

New() *bloomfilter
------------------

Returns a new initialised bloom filter.

(*bloomfilter) In(string) bool
------------------------------

Returns if the string is in the structure.

(*bloomfilter) Add(string)
--------------------------

Adds the string to the structure.