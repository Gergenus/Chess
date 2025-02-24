package main

func main() {
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x0, x1 int) int {
	if x1 < x0 {
		return x1
	} else {
		return x0
	}
}

func max(x0, x1 int) int {
	if x1 < x0 {
		return x0
	} else {
		return x1
	}
}

type Figure interface {
	ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool
	ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool
}

type King struct {
	Name string
	Code int
}

func newKing(name string, code int) *King {
	return &King{
		Name: name,
		Code: code,
	}
}

func (k King) ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool {
	return abs(x0-x1) <= 1 && abs(y0-y1) <= 1
}

func (k King) ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool {
	return abs(x0-x1) <= 1 && abs(y0-y1) <= 1
}

type Queen struct {
	Name string
	Code int
}

func newQueen(name string, code int) *Queen {
	return &Queen{
		Name: name,
		Code: code,
	}
}

func (q Queen) ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool {
	if !((x0-x1)*(y0-y1) == 0 || abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != 0 {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != 0 {
						return false
					}
				}
			}
		}
	} else {
		if x1 == x0 {
			for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
				if matrix[x0][i] != 0 {
					return false
				}
			}
		} else if y0 == y1 {
			for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
				if matrix[i][y0] != 0 {
					return false
				}
			}
		}
	}
	return true
}
func (q Queen) ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool {
	if !((x0-x1)*(y0-y1) == 0 || abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != 0 {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != 0 {
						return false
					}
				}
			}
		}
	} else {
		if x1 == x0 {
			for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
				if matrix[x0][i] != 0 {
					return false
				}
			}
		} else if y0 == y1 {
			for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
				if matrix[i][y0] != 0 {
					return false
				}
			}
		}
	}
	return true
}

type Rook struct {
	Name string
	Code int
}

func newRook(name string, code int) *Rook {
	return &Rook{
		Name: name,
		Code: code,
	}
}

func (r Rook) ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool {
	if !((x1-x0)*(y1-y0) == 0) {
		return false
	}
	if x1 == x0 {
		for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
			if matrix[x0][i] != 0 {
				return false
			}
		}
	} else if y0 == y1 {
		for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
			if matrix[i][y0] != 0 {
				return false
			}
		}
	}
	return true
}

func (r Rook) ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool {
	if !((x1-x0)*(y1-y0) == 0) {
		return false
	}
	if x1 == x0 {
		for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
			if matrix[x0][i] != 0 {
				return false
			}
		}
	} else if y0 == y1 {
		for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
			if matrix[i][y0] != 0 {
				return false
			}
		}
	}
	return true
}

type Bishop struct {
	Name string
	Code int
}

func newBishop(name string, code int) *Bishop {
	return &Bishop{
		Name: name,
		Code: code,
	}
}

func (b Bishop) ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool {
	if !(abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != 0 {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != 0 {
						return false
					}
				}
			}
		}
	}
	return true
}

func (b Bishop) ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool {
	if !(abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != 0 {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != 0 {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != 0 {
						return false
					}
				}
			}
		}
	}
	return true
}

type Knight struct {
	Name string
	Code int
}

func newKnight(name string, code int) *Knight {
	return &Knight{
		Name: name,
		Code: code,
	}
}

func (k Knight) ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool {
	return abs((x0-x1)*(y0-y1)) == 2
}

func (k Knight) ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool {
	return abs((x0-x1)*(y0-y1)) == 2
}

type Pawn struct {
	Name string
	Code int
}

func (p Pawn) ValidateTurn(x0, y0, x1, y1 int, matrix [][]int) bool {
	if x0 == 1 || x0 == 6 {
		if !(y0 == y1) || abs(x0-x1) > 2 {
			return false
		} else {
			if !(y0 == y1) || abs(x0-x1) > 1 {
				return false
			}
		}
	}
	return true
}

func (p Pawn) ValidateAttack(x0, y0, x1, y1 int, matrix [][]int) bool {
	return abs(x1-x0) == 1 && abs(y1-y0) == 1
}

type ChessBoard struct {
}
