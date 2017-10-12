package adapters

type (
	Logger interface {
		Debug(...interface{})
		Debugf(string, ...interface{})
	}
)
