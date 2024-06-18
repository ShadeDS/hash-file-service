package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/stretchr/testify/assert"
	"hash-file-service/util"
)

func TestInsertData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := &util.MockLogger{}
	query := `insert into files \(type, version, content, hash\) values \(\$1, \$2, \$3, \$4\)`
	mock.ExpectExec(query).WithArgs("core", "1.0.0", []byte("content"), "hash").WillReturnResult(sqlmock.NewResult(1, 1))

	err = InsertData(db, logger, "core", "1.0.0", "hash", []byte("content"))
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
