package zermelo

// Does a radix sort in place using supplied buffer space. len(r) must equal len(buffer)
func rsortUint32BYOB(r []uint32, buffer []uint32) {
	sliceCheckUint32(r, buffer)
	copy(buffer, r)

	// Radix is a byte, 4 bytes to a uint32
	for pass := uint(0); pass < 4; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			rsortPassUint32(r[:], buffer[:], pass)
		} else {
			rsortPassUint32(buffer[:], r[:], pass)
		}
	}
}

func rsortPassUint32(from []uint32, to []uint32, pass uint) {
	byteOffset := pass * rSortRadix
	byteMask := uint32(0xFF << byteOffset)
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

func sliceCheckUint32(a []uint32, b []uint32) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}
