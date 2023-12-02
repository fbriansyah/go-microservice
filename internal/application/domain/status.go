package domain

type Status int

const (
	Deleted  Status = Status(-2)
	Inactive Status = Status(0)
	Active   Status = Status(1)
)

func (s Status) String() string {
	switch s {
	case Deleted:
		return "Deleted"
	case Inactive:
		return "Inactive"
	case Active:
		return "Active"
	default:
		return "Unknown"
	}
}
