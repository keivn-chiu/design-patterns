package adapter

// adapter for making one interface fit to anther interface
// In this case, I want to use an adapter to make usbA fit to usbB
// IUsbA
type IUsbA interface {
	MountA() string
}

type implUsbA struct {
}

func (*implUsbA) MountA() string {
	return "A"
}

func NewUsbA() *implUsbA {
	return &implUsbA{}
}

// IUsbB
type IUsbB interface {
	MountB() string
}

type implUsbB struct {
}

func (*implUsbB) MountB() string {
	return "B"
}

func NewUsbB() *implUsbB {
	return &implUsbB{}
}

// UsbAdapter the adapter to hit the target, A -> B
type UsbAdapter struct {
	IUsbA
}

func NewUsbAdapter(usbA IUsbA) *UsbAdapter {
	return &UsbAdapter{usbA}
}

func (usb *UsbAdapter) MountB() string {
	return usb.MountA()
}
