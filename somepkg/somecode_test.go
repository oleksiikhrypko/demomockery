package somepkg

import (
	"errors"
	"testing"

	"demomockery/somepkg/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SomeLogic_simple(t *testing.T) {

	p := new(mocks.DataProvider)

	p.On("GetData", 1).Return(100, nil)
	p.On("GetData", 100).Return(200, nil)

	res, err := SomeLogic(p)
	require.NoError(t, err)
	assert.Equal(t, 300, res)
}

func Test_SomeLogic(t *testing.T) {

	p := new(mocks.DataProvider)
	p.On("GetData", 1).Return(2, nil)
	p.On("GetData", 2).Return(3, nil)

	p2 := new(mocks.DataProvider)
	p2.On("GetData", 1).Return(0, errors.New("some error 1"))

	p3 := new(mocks.DataProvider)
	p3.On("GetData", 1).Return(2, nil)
	p3.On("GetData", 2).Return(0, errors.New("some error 2"))

	type args struct {
		provider DataProvider
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{
			name: "Happy path",
			args: args{
				provider: p,
			},
			want:    5,
			wantErr: "",
		},
		{
			name: "Error path 1",
			args: args{
				provider: p2,
			},
			want:    0,
			wantErr: "some error 1",
		},
		{
			name: "Error path 2",
			args: args{
				provider: p3,
			},
			want:    0,
			wantErr: "some error 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SomeLogic(tt.args.provider)
			if (err != nil) && err.Error() != tt.wantErr {
				t.Errorf("SomeLogic error = %v, wantErr %v", err.Error(), tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
