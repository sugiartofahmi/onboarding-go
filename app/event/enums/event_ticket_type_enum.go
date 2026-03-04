package enums

type EventTicketTypeEnum int

const (
	Regular EventTicketTypeEnum = 1
	VIP     EventTicketTypeEnum = 2
	VIPPlus EventTicketTypeEnum = 3
)

func (e EventTicketTypeEnum) GetLabel() string {
	switch e {
	case Regular:
		return "Regular"
	case VIP:
		return "VIP"
	case VIPPlus:
		return "VIP Plus"
	default:
		return "Unknown"
	}
}
