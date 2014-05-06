package ZermeloGo

// Sorts a []uint32 using a Radix sort.  This uses O(n) extra memory

// Does a radix sort in place (but uses O(n) extra memory)
func SortUint32(r []uint32) {
	buffer := make([]uint32, len(r))
	rsortUint32BYOB(r, buffer)
}

// Sorts a []uint64 using a Radix sort.  This uses O(n) extra memory
func SortUint64(r []uint64) {
	buffer := make([]uint64, len(r))
	rsortUint64BYOB(r, buffer)
}

// Does a radix sort in place using supplied buffer space. len(r) must equal len(buffer)
func rsortUint32BYOB(r []uint32, buffer []uint32) {
	if len(r) != len(buffer) {
		panic("You can't use a buffer of a different size")
	}
	copy(buffer, r)

	// Radix is a byte, 4 bytes to a uint32
	for pass := 0; pass < 4; pass++ {
		// Radix offset and mask
		byteOffset := uint(pass * 8)
		byteMask := uint32(0xFF << byteOffset)
		// Keep track of the number of elements for each kind of byte
		var counts [256]int
		// Keep track of where room is made for byte groups in the buffer
		var offset [256]int
		// To save allocations, switch source and buffer roles back and forth
		toBuff := pass%2 == 0

		var passByte uint8 // Current byte value
		for i := 0; i < len(r); i++ {
			// For each elem to sort, fetch the byte at current radix
			if toBuff {
				passByte = uint8((r[i] & byteMask) >> byteOffset)
			} else {
				passByte = uint8((buffer[i] & byteMask) >> byteOffset)
			}
			// inc count of bytes of this type
			counts[passByte]++
		}
		// Make room for each group of bytes by calculating offset of each
		offset[0] = 0
		for i := 1; i < len(offset); i++ {
			offset[i] = offset[i-1] + counts[i-1]
		}
		// Swap values between the buffers by radix
		for i := 0; i < len(r); i++ {
			if toBuff {
				// Get the byte of each element at the radix
				passByte = uint8((r[i] & byteMask) >> byteOffset)
				// Copy the element depending on byte offsets
				buffer[offset[passByte]] = r[i]
			} else {
				passByte = uint8((buffer[i] & byteMask) >> byteOffset)
				r[offset[passByte]] = buffer[i]
			}
			// One less space empty in that byte groups reserved area, so move the offset
			offset[passByte]++
		}
	}
}

// Does a radix sort in place using supplied buffer space. len(r) must equal len(buffer)
func rsortUint64BYOB(r []uint64, buffer []uint64) {
	if len(r) != len(buffer) {
		panic("You can't use a buffer of a different size")
	}
	copy(buffer, r)

	// Radix is a byte, 8 bytes to a uint64
	for pass := 0; pass < 8; pass++ {
		// Radix offset and mask
		byteOffset := uint(pass * 8)
		byteMask := uint64(0xFF << byteOffset)
		// Keep track of the number of elements for each kind of byte
		var counts [256]int
		// Keep track of where room is made for byte groups in the buffer
		var offset [256]int
		// To save allocations, switch source and buffer roles back and forth
		toBuff := pass%2 == 0

		var passByte uint8 // Current byte value
		for i := 0; i < len(r); i++ {
			// For each elem to sort, fetch the byte at current radix
			if toBuff {
				passByte = uint8((r[i] & byteMask) >> byteOffset)
			} else {
				passByte = uint8((buffer[i] & byteMask) >> byteOffset)
			}
			// inc count of bytes of this type
			counts[passByte]++
		}
		// Make room for each group of bytes by calculating offset of each
		offset[0] = 0
		for i := 1; i < len(offset); i++ {
			offset[i] = offset[i-1] + counts[i-1]
		}
		// Swap values between the buffers by radix
		for i := 0; i < len(r); i++ {
			if toBuff {
				// Get the byte of each element at the radix
				passByte = uint8((r[i] & byteMask) >> byteOffset)
				// Copy the element depending on byte offsets
				buffer[offset[passByte]] = r[i]
			} else {
				passByte = uint8((buffer[i] & byteMask) >> byteOffset)
				r[offset[passByte]] = buffer[i]
			}
			// One less space empty in that byte groups reserved area, so move the offset
			offset[passByte]++
		}
	}
}
