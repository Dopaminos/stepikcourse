package functions

type Day int

const (
	_ Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func isWeekend(day Day) bool {
	return day >= 6 && day <= 7
}
