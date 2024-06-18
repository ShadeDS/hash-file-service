package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/heroiclabs/nakama-common/runtime"
	"hash-file-service/database"
	"hash-file-service/util"
)

type RequestPayload struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

type ResponsePayload struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

// ReadFileFromDisk reads the file from the given path and returns its content.
func ReadFileFromDisk(filePath string) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

// ValidateJSON validates the structure of the loaded JSON.
func ValidateJSON(data []byte) error {
	var jsonTest interface{}
	if err := json.Unmarshal(data, &jsonTest); err != nil {
		return err
	}
	return nil
}

// HandleFileHashingAndInsert handles the hashing of the file content and insertion into the database.
func HandleFileHashingAndInsert(ctx context.Context, logger runtime.Logger, db *sql.DB, fileType, version, providedHash string, fileContent []byte) (ResponsePayload, error) {
	// Calculate the hash of the content
	calculatedHash := util.CalculateHash(fileContent, logger)

	// Insert data into the database
	if err := database.InsertData(db, logger, fileType, version, calculatedHash, fileContent); err != nil {
		return ResponsePayload{}, runtime.NewError("unable to insert data", 13)
	}

	// Prepare the response
	response := ResponsePayload{
		Type:    fileType,
		Version: version,
		Hash:    calculatedHash,
	}

	// Verify the provided hash
	if providedHash != "" && providedHash != calculatedHash {
		response.Content = ""
	} else {
		response.Content = string(fileContent)
	}

	return response, nil
}

// ProcessFileRequest processes the file request by reading the file, validating the content, and handling hashing and insertion.
func ProcessFileRequest(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// Use defaults if not present in the payload
	var requestPayload RequestPayload
	if err := json.Unmarshal([]byte(payload), &requestPayload); err != nil {
		return "", runtime.NewError("unable to unmarshal payload", 13)
	}
	if requestPayload.Type == "" {
		requestPayload.Type = "core"
	}
	if requestPayload.Version == "" {
		requestPayload.Version = "1.0.0"
	}

	// Read the file from disk
	filePath := filepath.Join("/nakama/data", requestPayload.Type, requestPayload.Version+".json")
	fileContent, err := ReadFileFromDisk(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", runtime.NewError("file does not exist", 14)
		}
		return "", runtime.NewError("unable to read file", 13)
	}

	// Validate JSON content
	if err := ValidateJSON(fileContent); err != nil {
		return "", runtime.NewError("invalid JSON content", 13)
	}

	// Handle file hashing and database insertion
	response, err := HandleFileHashingAndInsert(ctx, logger, db, requestPayload.Type, requestPayload.Version, requestPayload.Hash, fileContent)
	if err != nil {
		return "", err
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return "", runtime.NewError("unable to marshal response", 13)
	}

	return string(responseData), nil
}

// RegisterFileProcessingRpc registers the RPC function for file processing.
func RegisterFileProcessingRpc(logger runtime.Logger, initializer runtime.Initializer) error {
	logger.Info("Registering RPC function for file processing...")
	if err := initializer.RegisterRpc("file_processing_rpc", ProcessFileRequest); err != nil {
		logger.Error("Unable to register file processing RPC", "error", err)
		return err
	}
	logger.Info("File processing RPC registered successfully")
	return nil
}
