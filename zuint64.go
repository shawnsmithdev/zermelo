package zermelo

// Does a radix sort in place using supplied buffer space. len(r) must equal len(buffer)
func rsortUint64BYOB(r []uint64, buffer []uint64) {
	sliceCheckUint64(r, buffer)
	copy(buffer, r)

	// Radix is a byte, 8 bytes to a uint64
	for pass := uint(0); pass < 8; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			rsortPassUint64(r[:], buffer[:], pass)
		} else {
			rsortPassUint64(buffer[:], r[:], pass)
		}
	}
}

func rsortPassUint64(from []uint64, to []uint64, pass uint) {
	byteOffset := pass * rSortRadix
	byteMask := uint64(0xFF << byteOffset)
	var counts [256]int // Keep track of the number of elements for each kind of byte
	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	var passByte uint8  // Current byte value

	for _, elem := range from {
		// For each elem to sort, fetch the byte at current radix
		passByte = uint8((elem & byteMask) >> byteOffset)
		// inc count of bytes of this type
		counts[passByte]++
	}

	// Make room for each group of bytes by calculating offset of each
	offset[0] = 0
	for i := 1; i < len(offset); i++ {
		offset[i] = offset[i-1] + counts[i-1]
	}

	// Swap values between the buffers by radix
	for _, elem := range from {
		passByte = uint8((elem & byteMask) >> byteOffset) // Get the byte of each element at the radix
		to[offset[passByte]] = elem                       // Copy the element depending on byte offsets
		offset[passByte]++                                // One less space, move the offset
	}
}

func sliceCheckUint64(a []uint64, b []uint64) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}
