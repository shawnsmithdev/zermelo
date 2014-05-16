package zermelo

// Does a radix sort in place using supplied buffer space. len(r) must equal len(buffer)
func rsortInt32BYOB(r []int32, buffer []int32) {
	sliceCheckInt32(r, buffer)
	copy(buffer, r)

	// Radix is a byte, 4 bytes to a uint32
	for pass := uint(0); pass < 4; pass++ {
		if pass%2 == 0 { // swap back and forth between buffers to save allocations
			rsortPassInt32(r[:], buffer[:], pass)
		} else {
			rsortPassInt32(buffer[:], r[:], pass)
		}
	}
}

func rsortPassInt32(from []int32, to []int32, pass uint) {
	byteOffset := pass * rSortRadix
	byteMask := int32(0xFF << byteOffset)
	var counts [256]int // Keep track of the number of elements for each kind of byte
	var offset [256]int // Keep track of where room is made for byte groups in the buffer
	var passByte uint8  // Current byte value
	for _, elem := range from {
		// For each elem to sort, fetch the byte at current radix
		passByte = uint8((elem & byteMask) >> byteOffset)
		// inc count of bytes of this type
		counts[passByte]++
	}

	if pass == 3 {
		// Handle signed values
		// Count negative elements (last 128 counts)
		negCnt := 0
		for i := 128; i < 256; i++ {
			negCnt += counts[i]
		}

		offset[0] = negCnt // Start of positives
		offset[128] = 0    // Start of negatives
		for i := 1; i < 128; i++ {
			// Positive values
			offset[i] = offset[i-1] + counts[i-1]
			// Negative values
			offset[i+128] = offset[i+127] + counts[i+127]
		}
	} else {
		offset[0] = 0
		for i := 1; i < len(offset); i++ {
			offset[i] = offset[i-1] + counts[i-1]
		}
	}

	// Swap values between the buffers by radix
	for _, elem := range from {
		passByte = uint8((elem & byteMask) >> byteOffset) // Get the byte of each element at the radix
		to[offset[passByte]] = elem                       // Copy the element depending on byte offsets
		offset[passByte]++
	}
}

func sliceCheckInt32(a []int32, b []int32) {
	if a == nil || b == nil || len(a) != len(b) {
		panic("Slices must be the same size and not nil")
	}
}
