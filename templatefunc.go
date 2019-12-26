package main

func monthBefore(queryMonth int) int {
	switch queryMonth {
	case 1:
		return 12
	default:
		return queryMonth - 1
	}
}

func monthAfter(queryMonth int) int {
	switch queryMonth {
	case 12:
		return 1
	default:
		return queryMonth + 1
	}
}

func yearBefore(queryYear, queryMonth int) int {
	switch queryMonth {
	case 1:
		return queryYear - 1
	default:
		return queryYear
	}
}

func yearAfter(queryYear, queryMonth int) int {
	switch queryMonth {
	case 12:
		return queryYear + 1
	default:
		return queryYear
	}
}
