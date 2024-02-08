package text

import "testing"

func TestTitle(t *testing.T) {
	result := Title("HELLO WORLD")
	expected := "Hello world"

	if result != expected {
		t.Errorf("Test Failed. Result should be: Hello world, now it is: %s", result)
	}
}
