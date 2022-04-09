package notification

type Notifier interface {
	Notify(message string) error
}
