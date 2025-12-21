package abstractions

type Device struct {
	name    string
	outputs []string
}

func NewDevice(
	name string,
	outputs []string,
) *Device {
	return &Device{name, outputs}
}
