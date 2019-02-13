package whitespace

import "testing"

func TestRemoveWhitespace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", ""},
		{"Space", " ", ""},
		{"Tab", "\t", ""},
		{"Newline", "\n", ""},
		{"Return", "\r", ""},
		{"Mixed", " \t\n\r", ""},
		{"Non-blank with space", "one two", "onetwo"},
		{"Non-blank with tab", "one \t two", "onetwo"},
		{"Non-blank with newline", "one\ntwo", "onetwo"},
		{"Non-blank with return", "one\rtwo", "onetwo"},
		{"Non-blank with mixed", "one \t\n\r two", "onetwo"},
		{"Non-blank with surrounding space", " onetwo ", "onetwo"},
		{"Non-blank with surrounding tab", "\t one\ttwo \t", "onetwo"},
		{"Non-blank with surrounding newline", "\none\ntwo\n", "onetwo"},
		{"Non-blank with surrounding return", "\rone\rtwo\r", "onetwo"},
		{"Non-blank with surrounding mixed", "\t\n\r onetwo \t\n\r", "onetwo"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := remove(test.input)

			if s != test.want {
				t.Errorf("got: <%v>, want: <%v>", s, test.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"Empty", "", true},
		{"Space", " ", true},
		{"Tab", "\t", true},
		{"Newline", "\n", true},
		{"Return", "\r", true},
		{"Mixed", " \t\n\r", true},
		{"Non-blank with space", "one two", false},
		{"Non-blank with tab", "one \t two", false},
		{"Non-blank with newline", "one\ntwo", false},
		{"Non-blank with return", "one\rtwo", false},
		{"Non-blank with mixed", "one \t\n\r two", false},
		{"Non-blank with surrounding space", " onetwo ", false},
		{"Non-blank with surrounding tab", "\t one\ttwo \t", false},
		{"Non-blank with surrounding newline", "\none\ntwo\n", false},
		{"Non-blank with surrounding return", "\rone\rtwo\r", false},
		{"Non-blank with surrounding mixed", "\t\n\r onetwo \t\n\r", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := isBlank(test.input)

			if b != test.want {
				t.Errorf("got: <%v>, want: <%v>", b, test.want)
			}
		})
	}
}
