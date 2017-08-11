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

func TestColorBright(t *testing.T) {
	if RedBright("MyString") != "\x1b[91m"+"MyString"+"\u001b[0m" {
		t.Error("Red bright colors don't match colors are not equal")
	}
}

func TestBgColorBright(t *testing.T) {
	if GreenBright(BgBlueBright("MyString")) != "\x1b[92m"+"\x1b[104m"+"MyString"+"\u001b[0m"+"\u001b[0m" {
		t.Error("Red bright colors don't match colors are not equal")
	}
}

func TestNestedStyles(t *testing.T) {
	if Underscore(GreenBright("foo")) != "\x1b[4m"+"\x1b[92m"+"foo"+"\u001b[0m"+"\u001b[0m" {
		t.Error("Nesting of styles fails")
	}

	if BgBlack(GreenBright("foo")) != "\x1b[40m"+"\x1b[92m"+"foo"+"\u001b[0m"+"\u001b[0m" {
		t.Error("Nesting of styles fails")
	}
}
