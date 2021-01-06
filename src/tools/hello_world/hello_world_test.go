package Hello_World

import "testing"

func TestHello(t *testing.T) {
	msg := Hello("Squishy123")
	want := "Hello World, Squishy123!"
	if msg != want {
		t.Error("Invalid String")
	}
}
