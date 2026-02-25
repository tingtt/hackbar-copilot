package toml

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDecodeStruct(t *testing.T) {
	t.Parallel()

	type args struct {
		r io.Reader
		i any
	}

	type child struct {
		Name string
	}
	type parent struct {
		Child      child
		PointerAry *[]child
	}
	tests := []struct {
		name    string
		args    args
		wantW   parent
		wantErr bool
	}{
		{
			name: "may return decoded",
			args: args{
				r: func() io.Reader {
					r := bytes.Buffer{}
					r.Write([]byte(`[[PointerAry]]
	Name = "a"
[Child]
	Name = "test"
				`))
					return &r
				}(),
				i: &parent{},
			},
			wantW: parent{
				Child:      child{Name: "test"},
				PointerAry: &[]child{{Name: "a"}},
			},
			wantErr: false,
		},
		{
			name: "may return decoded",
			args: args{
				r: func() io.Reader {
					r := bytes.Buffer{}
					r.Write([]byte(`PointerAry = [ ]
[Child]
  Name = "test"
`))
					return &r
				}(),
				i: &parent{},
			},
			wantW: parent{
				Child:      child{Name: "test"},
				PointerAry: &[]child{},
			},
			wantErr: false,
		}, {
			name: "may return decoded",
			args: args{
				r: func() io.Reader {
					r := bytes.Buffer{}
					r.Write([]byte(`[Child]
  Name = "test"
`))
					return &r
				}(),
				i: &parent{},
			},
			wantW: parent{
				Child:      child{Name: "test"},
				PointerAry: nil,
			},
			wantErr: false,
		}, {
			name: "may return error, if failed to read file",
			args: args{
				r: func() io.Reader {
					r := new(MockIOReader)
					r.On("Read", mock.Anything).Return(0, errors.New("fail to read file"))
					return r
				}(),
				i: &parent{},
			},
			wantW:   parent{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Decode(tt.args.r, tt.args.i)
			got := tt.args.i.(*parent)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, *got, tt.wantW)
		})
	}
}

func TestDecodeMap(t *testing.T) {
	t.Parallel()

	type args struct {
		r io.Reader
		i any
	}

	tests := []struct {
		name    string
		args    args
		wantW   map[string][]string
		wantErr bool
	}{
		{
			name: "may return decoded",
			args: args{
				r: func() io.Reader {
					r := bytes.Buffer{}
					r.Write([]byte(`sample = [ ]`))
					return &r
				}(),
				i: &map[string][]string{},
			},
			wantW:   map[string][]string{"sample": {}},
			wantErr: false,
		}, {
			name: "may return error, if failed to read file",
			args: args{
				r: func() io.Reader {
					r := new(MockIOReader)
					r.On("Read", mock.Anything).Return(0, errors.New("fail to read file"))
					return r
				}(),
				i: &map[string][]string{},
			},
			wantW:   map[string][]string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Decode(tt.args.r, tt.args.i)
			got := tt.args.i.(*map[string][]string)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, *got, tt.wantW)
		})
	}
}
