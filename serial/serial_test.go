package serial_test

import (
	"testing"

	"github.com/jurgen-kluft/go-esp32-control/serial"
	"github.com/jurgen-kluft/go-utils/assert"
)

func TestSerialBegin(t *testing.T) {
	baud := serial.BaudRate115200
	serial.Baud = 0
	serial.Begin(baud)
	assert.That("baud rate is 115200", t, serial.Baud, serial.BaudRate115200)
}
