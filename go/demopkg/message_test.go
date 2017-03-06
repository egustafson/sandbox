package demopkg

import "testing"

var msgValue = "Hola Mundo!\n"

func TestMessage(t *testing.T) {
	var msg = GetMessage()
	if msg != msgValue {
		t.Errorf("message value does not match.")
	}
}
