package enums

type EventStatusEnum int

const (
	Draft     EventStatusEnum = 1
	Published EventStatusEnum = 2
	Cancelled EventStatusEnum = 3
	Completed EventStatusEnum = 4
)

func (e EventStatusEnum) GetLabel() string {
	switch e {
	case Draft:
		return "Draft"
	case Published:
		return "Published"
	case Cancelled:
		return "Cancelled"
	case Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}
