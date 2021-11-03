package somepkg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_GetData(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()
	// describe expected behaviour
	mock.ExpectQuery(queryGetData).
		WithArgs(1).
		WillReturnRows(
			mock.
				NewRows([]string{"id", "val"}).
				AddRow(
					10, "value A").
				AddRow(
					20, "value B"),
		)

	// make storage for test
	// make sqlx wrapper, because the storage uses sqlx instead of sql
	dbx := sqlx.NewDb(db, "sqlmock")
	storage := NewStorage(dbx)

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
