package values

type ValueWriter interface {
	WriteValue() (string, error)
}
