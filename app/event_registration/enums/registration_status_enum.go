package enums

type EventRegistrationStatusEnum int

const (
	Pending   EventRegistrationStatusEnum = 1
	Confirmed EventRegistrationStatusEnum = 2
	Cancelled EventRegistrationStatusEnum = 3
)

func (e EventRegistrationStatusEnum) GetLabel() string {
	switch e {
	case Pending:
		return "Pending"
	case Confirmed:
		return "Confirmed"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}
