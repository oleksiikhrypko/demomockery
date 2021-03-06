package somepkg

import (
	"testing"

	"demomockery/somepkg/mocks"
	"demomockery/somepkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SomeLogic_simple(t *testing.T) {
	// build mock
	p := new(mocks.DataProvider)
	res := []models.DataRec{
		{ID: 10, Val: "value A"},
		{ID: 20, Val: "value B"},
	}
	// describe expected behaviour
	p.On("GetData", 1).Return(res, nil)

	// call tested function
	c, err := SomeLogic(p)

	// check results
	require.NoError(t, err)
	require.NotNil(t, c)
	require.Equal(t, 2, len(c))
	// check rec1
	assert.Equal(t, 10, c[0].ID)
	assert.Equal(t, "value A", c[0].Val)
	// check rec2
	assert.Equal(t, 20, c[1].ID)
	assert.Equal(t, "value B", c[1].Val)
}
