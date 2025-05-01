package digital_test

import (
	"testing"

	"github.com/jurgen-kluft/go-esp32-control/digital"
	"github.com/jurgen-kluft/go-utils/assert"
)

func TestDigitalWrite(t *testing.T) {
	pin := 1
	digital.GPIOValues[pin] = digital.Low
	digital.Write(pin, digital.High)
	val := digital.GPIOValues[pin]
	assert.That("val is set to high", t, val, digital.High)
}

func TestPinMode(t *testing.T) {
	pin := 1
	digital.GPIOModes[pin] = digital.ModeInput
	digital.PinMode(pin, digital.ModeOutput)
	mode := digital.GPIOModes[pin]
	assert.That("mode is set to output", t, mode, digital.ModeOutput)
}
