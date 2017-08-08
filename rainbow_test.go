package rainbow

import (
	"testing"
)

// Test params on a Color Decorator
func TestColor(t *testing.T) {
	if Blue("MyString") != "\x1b[34m"+"MyString"+"\u001b[0m" {
		t.Error("Two black colors are not equal")
	}

}
