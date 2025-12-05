package day5

type rangePair struct {
	lower int64
	upper int64
}

func (rp *rangePair) isBetween(num int64) bool {
	return num >= rp.lower && num <= rp.upper
}
