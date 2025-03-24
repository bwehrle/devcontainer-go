package hello_test

import(
	"github.com/josys-src/dev-container-go/module"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		test string // description of this test case
		name string
		expected string
	}{
		{"test1", "foo", "Hello, foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hello.Hello(tt.name)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.expected {
				t.Errorf("Hello() = %v, want %v", got, tt.expected)
			}
		})
	}
}
