package ccitt

var (
	white byte = 1
	black byte = 0
)

type Encoder struct {
	K                      int
	EndOfLine              bool
	EncodedByteAlign       bool
	Columns                int
	Rows                   int
	EndOfBlock             bool
	BlackIs1               bool
	DamagedRowsBeforeError int
}

func (e *Encoder) Encode(pixels [][]byte) []byte {
	if e.BlackIs1 {
		white = 0
		black = 1
	} else {
		white = 1
		black = 0
	}

	if e.K == 0 {
		// do G31D
		return e.encodeG31D(pixels)
	}

	if e.K > 0 {
		var encoded []byte

		bitPos := 0
		prevBitPos := 0
		for i := 0; i < len(pixels); i += e.K {
			var encodedRow []byte

			encodedRow, bitPos = encodeRow(pixels[i], bitPos, EOL1)

			if len(encoded) > 0 && prevBitPos != 0 {
				// no byte aligning
				// TODO: change to be dependent on the ByteAlign option
				encoded[len(encoded)-1] = encoded[len(encoded)-1] | encodedRow[0]

				encoded = append(encoded, encodedRow[1:]...)
			} else {
				encoded = append(encoded, encodedRow...)
			}

			prevBitPos = 0
			bitPos = 0
			for j := i + 1; j < (i+e.K) && j < len(pixels); j++ {
				encodedRow, bitPos = AddCode(nil, bitPos, EOL0)

				a0 := -1

				var a1, b1, b2 int

				eol := false

				for !eol {
					a1 = seekChangingElem(pixels[j], a0)
					b1 = seekB1(pixels[j], pixels[j-1], a0)
					b2 = seekChangingElem(pixels[j], b1)

					if b2 < a1 {
						// do pass mode
						encodedRow, bitPos = AddCode(encodedRow, bitPos, P)

						a0 = b2
					} else {
						if (b1 - a1) > 3 {
							a2 := seekChangingElem(pixels[j], a1)

							// do horizontal mode
							isWhite := pixels[j][a0] == white

							encodedRow, bitPos = AddCode(encodedRow, bitPos, H)
							a0a1RunLen, _ := calcRun(pixels[j], isWhite, a0)

							a0a1Code, a0a1Rem, a0a1IsTerminal := getRunCode(a0a1RunLen, isWhite)
							encodedRow, bitPos = AddCode(encodedRow, bitPos, a0a1Code)

							if !a0a1IsTerminal {
								if a0a1Rem == 0 {
									// obviously terminal
									a0a1Code, a0a1Rem, _ = getRunCode(a0a1Rem, isWhite)
									encodedRow, bitPos = AddCode(encodedRow, bitPos, a0a1Code)
								} else {
									for {
										a0a1Code, a0a1Rem, a0a1IsTerminal = getRunCode(a0a1Rem, isWhite)
										encodedRow, bitPos = AddCode(encodedRow, bitPos, a0a1Code)

										if a0a1IsTerminal {
											break
										}

										if a0a1Rem == 0 && !a0a1IsTerminal {
											// obviously terminal
											a0a1Code, a0a1Rem, _ = getRunCode(a0a1Rem, isWhite)
											encodedRow, bitPos = AddCode(encodedRow, bitPos, a0a1Code)

											break
										}
									}
								}
							}

							isWhite = !isWhite

							a1a2RunLen, _ := calcRun(pixels[j], isWhite, a1)

							a1a2Code, a1a2Rem, a1a2IsTerminal := getRunCode(a1a2RunLen, isWhite)
							encodedRow, bitPos = AddCode(encodedRow, bitPos, a1a2Code)

							if !a1a2IsTerminal {
								if a1a2Rem == 0 {
									// obviously terminal
									a1a2Code, a1a2Rem, _ = getRunCode(a1a2Rem, isWhite)
									encodedRow, bitPos = AddCode(encodedRow, bitPos, a1a2Code)
								} else {
									for {
										a1a2Code, a1a2Rem, a1a2IsTerminal = getRunCode(a1a2Rem, isWhite)
										encodedRow, bitPos = AddCode(encodedRow, bitPos, a1a2Code)

										if a1a2IsTerminal {
											break
										}

										if a1a2Rem == 0 && !a1a2IsTerminal {
											// obviously terminal
											a1a2Code, a1a2Rem, _ = getRunCode(a1a2Rem, isWhite)
											encodedRow, bitPos = AddCode(encodedRow, bitPos, a1a2Code)

											break
										}
									}
								}
							}

							a0 = a2

							// set eol
							if a0 >= len(pixels[j]) {
								eol = true
							}
						} else {
							// do vertical mode

							var vCode Code

							switch b1 - a1 {
							case -1:
								vCode = V1R
							case -2:
								vCode = V2R
							case -3:
								vCode = V3R
							case 0:
								vCode = V0
							case 1:
								vCode = V1L
							case 2:
								vCode = V2L
							case 3:
								vCode = V3L
							}

							encodedRow, bitPos = AddCode(encodedRow, bitPos, vCode)

							a0 = a1

							// set eol
							if a0 >= len(pixels[j]) {
								eol = true
							}
						}
					}
				}

				if len(encoded) > 0 && prevBitPos != 0 {
					// no byte aligning
					// TODO: change to be dependent on the ByteAlign option
					encoded[len(encoded)-1] = encoded[len(encoded)-1] | encodedRow[0]

					encoded = append(encoded, encodedRow[1:]...)
				} else {
					encoded = append(encoded, encodedRow...)
				}
			}
		}

		// put rtc
		var encodedRTC []byte

		encodedRTC, bitPos = encodeRTC(bitPos)

		if len(encoded) > 0 && prevBitPos != 0 {
			encoded[len(encoded)-1] = encoded[len(encoded)-1] | encodedRTC[0]

			encoded = append(encoded, encodedRTC[1:]...)
		} else {
			encoded = append(encoded, encodedRTC...)
		}

		return encoded
	}

	// TODO: add G32D and G4
	return nil
}

func seekChangingElem(row []byte, currElem int) int {
	i := 0
	for i < len(row) {
		if row[i] != row[currElem] {
			break
		}

		i++
	}

	return i
}

func seekB1(codingLine, refLine []byte, a0 int) int {
	changingElem := seekChangingElem(refLine, a0)
	if codingLine[a0] == refLine[changingElem] {
		changingElem = seekChangingElem(refLine, changingElem)
	}

	return changingElem
}

func (e *Encoder) encodeG31D(pixels [][]byte) []byte {
	var encoded []byte

	bitPos := 0
	prevBitPos := 0
	for i := range pixels {
		var encodedRow []byte

		encodedRow, bitPos = encodeRow(pixels[i], bitPos, EOL)

		if len(encoded) > 0 && prevBitPos != 0 {
			// no byte aligning
			// TODO: change to be dependent on the ByteAlign option
			encoded[len(encoded)-1] = encoded[len(encoded)-1] | encodedRow[0]

			encoded = append(encoded, encodedRow[1:]...)
		} else {
			encoded = append(encoded, encodedRow...)
		}

		prevBitPos = bitPos
	}

	var encodedRTC []byte

	encodedRTC, bitPos = encodeRTC(bitPos)

	if len(encoded) > 0 && prevBitPos != 0 {
		encoded[len(encoded)-1] = encoded[len(encoded)-1] | encodedRTC[0]

		encoded = append(encoded, encodedRTC[1:]...)
	} else {
		encoded = append(encoded, encodedRTC...)
	}

	return encoded
}

// encodeRTC encodes the RTC code (6 EOL in a row). bitPos is equal to the one in
// encodeRow
func encodeRTC(bitPos int) ([]byte, int) {
	var encoded []byte

	for i := 0; i < 6; i++ {
		encoded, bitPos = AddCode(encoded, bitPos, EOL)
	}

	return encoded, bitPos % 8
}

func encodeRTC2D(bitPos int) ([]byte, int) {
	var encoded []byte

	for i := 0; i < 6; i++ {
		encoded, bitPos = AddCode(encoded, bitPos, EOL1)
	}

	return encoded, bitPos % 8
}

// encodeRow encodes single raw of the image. bitPos is the bit position
// global for the row array. bitPos is used to indicate where to start the
// encoded sequences. It is used for the EncodedByteAlign option implementation.
// Returns the encoded data and the number of the bits taken from the last byte
func encodeRow(row []byte, bitPos int, eolCode Code) ([]byte, int) {
	// always start with whites
	isWhite := true
	var encoded []byte

	// always add EOL before the scan line
	encoded, bitPos = AddCode(nil, bitPos, eolCode)

	bytePos := 0
	var runLen int
	for bytePos < len(row) {
		runLen, bytePos = calcRun(row, isWhite, bytePos)

		code, rem, isTerminal := getRunCode(runLen, isWhite)
		encoded, bitPos = AddCode(encoded, bitPos, code)

		if !isTerminal {
			if rem == 0 {
				// obviously terminal
				code, rem, _ = getRunCode(rem, isWhite)
				encoded, bitPos = AddCode(encoded, bitPos, code)
			} else {
				for {
					code, rem, isTerminal = getRunCode(rem, isWhite)
					encoded, bitPos = AddCode(encoded, bitPos, code)

					if isTerminal {
						break
					}

					if rem == 0 && !isTerminal {
						// obviously terminal
						code, rem, _ = getRunCode(rem, isWhite)
						encoded, bitPos = AddCode(encoded, bitPos, code)

						break
					}
				}
			}
		}

		// switch color
		isWhite = !isWhite
	}

	return encoded, bitPos % 8
}

// getRunCode gets the code for the specified run. If the code is not
// terminal, returns the remainder to be determined later. Otherwise
// returns 0 remainder. Also returns the bool flag to indicate if
// the code is terminal
func getRunCode(runLen int, isWhite bool) (Code, int, bool) {
	if runLen < 64 {
		if isWhite {
			return WTerms[runLen], 0, true
		} else {
			return BTerms[runLen], 0, true
		}
	} else {
		multiplier := runLen / 64

		// stands for lens more than 2560 which are not
		// covered by the Huffman codes
		if multiplier > 40 {
			return CommonMakeups[2560], runLen - 2560, false
		}

		// stands for lens more than 1728. These should be common
		// for both colors
		if multiplier > 27 {
			return CommonMakeups[multiplier*64], runLen - multiplier*64, false
		}

		// for lens < 27 we use the specific makeups for each color
		if isWhite {
			return WMakeups[multiplier*64], runLen - multiplier*64, false
		} else {
			return BMakeups[multiplier*64], runLen - multiplier*64, false
		}
	}
}

// calcRun calculates the nex pixel run. Returns the number of the
// pixels and the new position in the original array
func calcRun(row []byte, isWhite bool, pos int) (int, int) {
	count := 0
	for pos < len(row) {
		if isWhite {
			if row[pos] != white {
				break
			}
		} else {
			if row[pos] != black {
				break
			}
		}

		count++
		pos++
	}

	return count, pos
}

func AddCode(encoded []byte, pos int, code Code) ([]byte, int) {
	i := 0
	for i < code.BitsWritten {
		bytePos := pos / 8
		bitPos := pos % 8

		if bytePos >= len(encoded) {
			encoded = append(encoded, 0)
		}

		toWrite := 8 - bitPos
		leftToWrite := code.BitsWritten - i
		if toWrite > leftToWrite {
			toWrite = leftToWrite
		}

		if i < 8 {
			encoded[bytePos] = encoded[bytePos] | byte(code.Code>>uint(8+bitPos-i))&masks[8-toWrite-bitPos]
		} else {
			encoded[bytePos] = encoded[bytePos] | (byte(code.Code<<uint(i-8))&masks[8-toWrite])>>uint(bitPos)
		}

		pos += toWrite

		i += toWrite
	}

	return encoded, pos
}
