package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool
type Chessboard map[string]File

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if _, file_exist := cb[file]; !file_exist {
		return 0
	}
	var cnt int
	for _, occupied := range cb[file] {
		if occupied {
			cnt += 1
		}
	}
	return cnt
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
		return 0
	}
	var ret int
	for file := range cb {
		if cb[file][rank-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var ret int
	for file := range cb {
		for range cb[file] {
			ret++
		}
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var ret int
	for file := range cb {
		for _, is_occupied := range cb[file] {
			if is_occupied {
				ret++
			}
		}
	}
	return ret
}
