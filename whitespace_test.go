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
		{"Nonbreaking space", "\u00a0", ""},
		{"Tab", "\t", ""},
		{"Vertical tab", "\v", ""},
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
			s := Remove(test.input)

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
		{"Nonbreaking space", "\u00a0", true},
		{"Tab", "\t", true},
		{"Vertical tab", "\v", true},
		{"Newline", "\n", true},
		{"Return", "\r", true},
		{"Mixed", " \t\n\r\v\u00a0", true},
		{"Non-blank with space", "one two", false},
		{"Non-blank with nonbreaking space", "one\u00a0two", false},
		{"Non-blank with tab", "one \t two", false},
		{"Non-blank with vertical tab", "one \v two", false},
		{"Non-blank with newline", "one\ntwo", false},
		{"Non-blank with return", "one\rtwo", false},
		{"Non-blank with mixed", "one \t\n\r two", false},
		{"Non-blank with surrounding space", " onetwo ", false},
		{"Non-blank with surrounding nonbreaking space", "\u00a0onet\u00a0wo\u00a0", false},
		{"Non-blank with surrounding tab", "\tone\ttwo\t", false},
		{"Non-blank with surrounding vertical tab", "\vone\vtwo\v", false},
		{"Non-blank with surrounding newline", "\none\ntwo\n", false},
		{"Non-blank with surrounding return", "\rone\rtwo\r", false},
		{"Non-blank with surrounding mixed", "\t\n\r\v\u00a0 onetwo \t\n\r\v\u00a0", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := IsBlank(test.input)

			if b != test.want {
				t.Errorf("got: <%v>, want: <%v>", b, test.want)
			}
		})
	}
}

func TestHasBlank(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  bool
	}{
		{"Empty slice", nil, true},
		{"Single empty string", []string{""}, true},
		{"Multiple empty strings", []string{"", "", ""}, true},
		{"Single space string", []string{" "}, true},
		{"Multiple space strings", []string{" ", " ", " "}, true},
		{"Single nonbreaking space string", []string{"\u00a0"}, true},
		{"Multiple nonbreaking space strings", []string{"\u00a0", "\u00a0", "\u00a0"}, true},
		{"Single tab string", []string{"\t"}, true},
		{"Multiple tab strings", []string{"\t", "\t", "\t"}, true},
		{"Single vertical tab string", []string{"\v"}, true},
		{"Multiple vertical tab strings", []string{"\v", "\v", "\v"}, true},
		{"Single newline string", []string{"\n"}, true},
		{"Multiple newline strings", []string{"\n", "\n", "\n"}, true},
		{"Single return string", []string{"\r"}, true},
		{"Multiple return strings", []string{"\r", "\r", "\r"}, true},
		{"Single mixed string", []string{" \t\n\r"}, true},
		{"Multiple mixed strings", []string{" \t\n", "\r\n", "\t\n ", "\t\n\r "}, true},
		{"Non-blank with single empty string", []string{"123", ""}, true},
		{"Non-blank with multiple empty strings", []string{"", "123", ""}, true},
		{"Non-blank with single space string", []string{" ", "123 "}, true},
		{"Non-blank with multiple space strings", []string{" ", " ", " 123"}, true},
		{"Non-blank with single nonbreaking space string", []string{"\u00a0", "123\u00a0"}, true},
		{"Non-blank with multiple nonbreaking space strings", []string{"\u00a0", "\u00a0", "\u00a0123"}, true},
		{"Non-blank with single tab string", []string{"\t123", "\t"}, true},
		{"Non-blank with multiple tab strings", []string{"\t", "\t123", "\t"}, true},
		{"Non-blank with single vertical tab string", []string{"\v123", "\v"}, true},
		{"Non-blank with multiple vertical tab strings", []string{"\v", "\v123", "\v"}, true},
		{"Non-blank with single newline string", []string{"123\n", "\n"}, true},
		{"Non-blank with multiple newline strings", []string{"\n", "\n", "\n123"}, true},
		{"Non-blank with single return string", []string{"\r123", "\r"}, true},
		{"Non-blank with multiple return strings", []string{"123\r", "\r", "\r"}, true},
		{"Non-blank with single mixed string", []string{"\t\n\r", " \t\n 123 \r"}, true},
		{"Non-blank with multiple mixed strings", []string{" 123\t\n\v", "\r\n\u00a0", "\t\n\u00a0 ", "123\t\n\r\v\u00a0 "}, true},
		{"Single non-blank string", []string{"123"}, false},
		{"Multiple non-blank strings", []string{"123", "456", "789"}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := HasBlank(test.input)

			if b != test.want {
				t.Errorf("got: <%v>, want: <%v>", b, test.want)
			}
		})
	}
}
