package leitor

import (
	"reflect"
	"testing"
)

func Test_verificaCaractereRomano(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// r == 'I' || r == 'V' || r == 'X' || r == 'C' || r == 'L' || r == 'M' || r == 'D'
		{"I", args{'I'}, true},
		{"V", args{'V'}, true},
		{"X", args{'X'}, true},
		{"C", args{'C'}, true},
		{"L", args{'L'}, true},
		{"M", args{'M'}, true},
		{"D", args{'D'}, true},
		{"x", args{'x'}, false},
		{"O", args{'O'}, false},
		{"Q", args{'Q'}, false},
		{"R", args{'R'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verificaCaractereRomano(tt.args.r); got != tt.want {
				t.Errorf("verificaCaractereRomano() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringParaSequenciaDeRomanos(t *testing.T) {
	type args struct {
		texto string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"XXAISPODIXOPIOLL", args{"XXAISPODIXOPIOLL"}, []string{"XX", "I", "DIX", "I", "LL"}},
		{"MAXI", args{"MAXI"}, []string{"M", "XI"}},
		{"AXIBIV", args{"AXIBIV"}, []string{"XI", "IV"}},
		{"AXXOILLODP", args{"AXXOILLODP"}, []string{"XX", "ILL", "D"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringParaSequenciaDeRomanos(tt.args.texto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringParaSequenciaDeRomanos() = %v, want %v", got, tt.want)
			}
		})
	}
}
