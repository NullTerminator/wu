package adapters

type (
	Logger interface {
		Debugf(string, ...interface{})
	}
)
