package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test 0",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "base test 1",
			str:  "my name is ted",
			want: "my name is ted",
		},
		{
			name: "base test 2",
			str:  "Hello World!",
			want: "!hello !world!",
		},
		{
			name: "base test 3",
			str:  "OK! I will do that!",
			want: "!o!k! !i will do that!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test 0",
			str:  "!ted",
			want: "001000100110100101",
		},
		{
			name: "base test 1",
			str:  "!hello !world!",
			want: "00100000111010010010010011000111001" +
				"0000000011100010100000100100101001000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		encoder := New()

		t.Run(tt.name, func(t *testing.T) {
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedText []byte
		want        string
	}{
		{
			name:        "base test",
			encodedText: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		decoder := New()

		t.Run(tt.name, func(t *testing.T) {
			if got := decoder.Decode(tt.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
