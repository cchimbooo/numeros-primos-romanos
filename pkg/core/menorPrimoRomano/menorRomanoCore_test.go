package menorPrimoRomano

import (
	"reflect"
	"testing"
)

func TestCriarProcessarStringInterface(t *testing.T) {
	type args struct {
		primos []int
	}
	tests := []struct {
		name string
		args args
		want ProcessarStringInterface
	}{
		{"teste 1", args{primos: []int{1, 2, 3, 4, 5}}, coreMenorRomano{mapa: map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5}}},
		{"teste 1", args{primos: []int{2}}, coreMenorRomano{mapa: map[string]int{"II": 2}}},
		{"teste 1", args{primos: []int{25, 3}}, coreMenorRomano{mapa: map[string]int{"XXV": 25, "III": 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CriarProcessarStringInterface(tt.args.primos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CriarProcessarStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coreMenorRomano_ProcessarString(t *testing.T) {

	m := map[string]int{
		"II":   2,
		"III":  3,
		"V":    5,
		"VII":  7,
		"XI":   11,
		"XIII": 13,
		"XVII": 17,
		"IXX":  19,
	}

	type fields struct {
		mapa map[string]int
	}
	type args struct {
		texto string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MenorNumero
	}{
		{"AXIBIV", fields{mapa: m}, args{texto: "AXIBIV"}, MenorNumero{Romano: "XI", Decimal: 11}},
		// Quer 0 no DIDATICA pois I não é número e DID não é válido.
		{"DIDATICA", fields{mapa: m}, args{texto: "DIDATICA"}, MenorNumero{Romano: "", Decimal: 0}},
		{"ASLKDJOIQWKEMOOIPUUV", fields{mapa: m}, args{texto: "ASLKDJOIQWKEMOOIPUUV"}, MenorNumero{Romano: "V", Decimal: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coreMenorRomano{
				mapa: tt.fields.mapa,
			}
			if got := c.ProcessarString(tt.args.texto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessarString() = %v, want %v", got, tt.want)
			}
		})
	}
}
