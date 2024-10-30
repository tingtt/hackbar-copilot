package toml

import (
	"bytes"
	"testing"

	"github.com/tingtt/options"
)

func TestEncode(t *testing.T) {
	type args struct {
		i interface{}
		o []options.Applier[Option]
	}

	type child struct {
		Name string
	}
	type parent struct {
		Child child
	}

	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				i: parent{Child: child{Name: "test"}},
				o: nil,
			},
			wantW: `[Child]
  Name = "test"
`,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				i: parent{Child: child{Name: "test"}},
				o: []Applier{WithIndent("")},
			},
			wantW: `[Child]
Name = "test"
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Encode(w, tt.args.i, tt.args.o...); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Encode() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
