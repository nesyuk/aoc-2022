package lib

func Min[num int](n1, n2 num) num {
	if n1 < n2 {
		return n1
	}
	return n2
}

func Max[num int](n1, n2 num) num {
	if n1 > n2 {
		return n1
	}
	return n2
}

func Abs[num int](n num) num {
	if n > 0 {
		return n
	}
	return n * -1
}
