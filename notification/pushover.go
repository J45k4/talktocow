package notification

type PushoverNotifier struct {
	UserKey string
	Token   string
}

func (n *PushoverNotifier) Notify(message string) error {
	return nil
}
