package solvers

var dayMap = make(map[int]func() string)

func init() {
	//dayMap[1] = day1.Solve
}

func Solve(day int) string {
	return dayMap[day]()
}
