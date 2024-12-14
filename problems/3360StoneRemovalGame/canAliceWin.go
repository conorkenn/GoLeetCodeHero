package leetcode

func canAliceWin(n int) bool {

	remove := 10
	isAliceTurn := true

	for n > 0 {
		if n < remove {
			return !isAliceTurn
		}
		n -= remove
		remove -= 1
		isAliceTurn = !isAliceTurn
	}

	return !isAliceTurn
}
