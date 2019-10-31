package main

import (
	"reflect"
	"testing"
)

func Test_rgbToHex(t *testing.T) {
	tests := []struct {
		name string
		rgb  [3]int
		want string
		err  *ErrOutOfRange
	}{
		{"Black", [3]int{0, 0, 0}, "#000000", nil},
		{"White", [3]int{255, 255, 255}, "#FFFFFF", nil},
		{"Red", [3]int{255, 0, 0}, "#FF0000", nil},
		{"Lime", [3]int{0, 255, 0}, "#00FF00", nil},
		{"Blue", [3]int{0, 0, 255}, "#0000FF", nil},
		{"Yellow", [3]int{255, 255, 0}, "#FFFF00", nil},
		{"Cyan", [3]int{0, 255, 255}, "#00FFFF", nil},
		{"Magenta", [3]int{255, 0, 255}, "#FF00FF", nil},
		{"Silver", [3]int{192, 192, 192}, "#C0C0C0", nil},
		{"Gray", [3]int{128, 128, 128}, "#808080", nil},
		{"Maroon", [3]int{128, 0, 0}, "#800000", nil},
		{"Olive", [3]int{128, 128, 0}, "#808000", nil},
		{"Green", [3]int{0, 128, 0}, "#008000", nil},
		{"Purple", [3]int{128, 0, 128}, "#800080", nil},
		{"Teal", [3]int{0, 128, 128}, "#008080", nil},
		{"Navy", [3]int{0, 0, 128}, "#000080", nil},
		{"Invalid R", [3]int{-1, 100, 100}, "", new(ErrOutOfRange)},
		{"Invalid G", [3]int{100, 300, 200}, "", new(ErrOutOfRange)},
		{"Invalid B", [3]int{100, 200, 256}, "", new(ErrOutOfRange)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := rgbToHex(tc.rgb)
			if err != nil {
				switch err.(type) {
				case *ErrOutOfRange:
					if reflect.TypeOf(tc.err) != reflect.TypeOf(err) {
						t.Errorf("got error of type %T, wanted %T", err, tc.err)
					}
				default:
					t.Fatalf("UNHANDLED ERROR TYPE")
				}
			}
			if got != tc.want {
				t.Errorf("rgbToHex() = %v, want %v", got, tc.want)
			}
		})
	}
}
