package toml

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tingtt/options"
)

func TestEncode(t *testing.T) {
	t.Parallel()

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
			name: "may write encoded toml with indent to io.Writer",
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
			name: "may write encoded toml to io.Writer",
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
			t.Parallel()

			w := &bytes.Buffer{}
			err := Encode(w, tt.args.i, tt.args.o...)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, w.String(), tt.wantW)
		})
	}
}
