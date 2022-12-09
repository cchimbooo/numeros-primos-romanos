package romanos

import "testing"

func TestGerar(t *testing.T) {
	type args struct {
		numero int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"999", args{numero: 999}, "CMXCIX", false},
		{"14", args{numero: 14}, "XIV", false},
		{"2000", args{numero: 2000}, "MM", false},
		{"3999", args{numero: 3999}, "MMMCMXCIX", false},
		{"803", args{numero: 803}, "DCCCIII", false},
		{"654", args{numero: 654}, "DCLIV", false},
		{"647", args{numero: 647}, "DCXLVII", false},
		{"429", args{numero: 429}, "CDXXIX", false},
		{"4000", args{numero: 4000}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Gerar(tt.args.numero)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gerar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gerar() got = %v, want %v", got, tt.want)
			}
		})
	}
}
