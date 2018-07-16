package subject

type Subject interface {
	GetSubject() string
}

type subject struct {
	name string
}

// GetSubject func
func GetSubject() string {
	return "Subject"
}

func getSubject() string {
	return "Subject"
}
