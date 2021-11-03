package somepkg

import (
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache_GetData(t *testing.T) {
	// build Cache Client mock
	cl, mock := redismock.NewClientMock()
	// describe behaviour expected
	mock.ExpectHGetAll("pfx_1").SetVal(map[string]string{"10": "value A", "20": "value B"})

	// make cache for test
	storage := NewCache(cl)

	// call tested function
	c, err := storage.GetData(1)

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
