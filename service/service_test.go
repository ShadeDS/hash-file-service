package service

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"hash-file-service/util"
)

func TestProcessFileRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := &util.MockLogger{}
	query := `insert into files \(type, version, content, hash\) values \(\$1, \$2, \$3, \$4\)`
	mock.ExpectExec(query).WithArgs("core", "1.0.0", []byte(`{"key":"value"}`), "e43abcf3375244839c012f9633f95862d232a95b00d5bc7348b3098b9fed7f32").WillReturnResult(sqlmock.NewResult(1, 1))

	payload := `{"type":"core","version":"1.0.0"}`
	fileContent := []byte(`{"key":"value"}`)

	// Ensure the directory exists
	dirPath := "/nakama/data/core"
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		t.Fatalf("An error '%s' was not expected when creating the directory", err)
	}

	// Write the test file
	filePath := filepath.Join(dirPath, "1.0.0.json")
	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		t.Fatalf("An error '%s' was not expected when writing the test file", err)
	}

	defer os.RemoveAll("/nakama/data")

	response, err := ProcessFileRequest(context.Background(), logger, db, nil, payload)
	assert.NoError(t, err)
	assert.Contains(t, response, `"type":"core"`)
	assert.Contains(t, response, `"version":"1.0.0"`)
	assert.Contains(t, response, `"hash":"e43abcf3375244839c012f9633f95862d232a95b00d5bc7348b3098b9fed7f32"`)
	assert.Contains(t, response, `"content":"{\"key\":\"value\"}"`)

	assert.NoError(t, mock.ExpectationsWereMet())
}
