package indexnode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamTable_Init(t *testing.T) {
	Params.Init()
}

func TestParamTable_Address(t *testing.T) {
	address := Params.Address
	fmt.Println(address)
}

func TestParamTable_MinIOAddress(t *testing.T) {
	address := Params.MinIOAddress
	fmt.Println(address)
}

func TestParamTable_MinIOAccessKeyID(t *testing.T) {
	accessKeyID := Params.MinIOAccessKeyID
	assert.Equal(t, accessKeyID, "minioadmin")
}

func TestParamTable_MinIOSecretAccessKey(t *testing.T) {
	secretAccessKey := Params.MinIOSecretAccessKey
	assert.Equal(t, secretAccessKey, "minioadmin")
}

func TestParamTable_MinIOUseSSL(t *testing.T) {
	useSSL := Params.MinIOUseSSL
	assert.Equal(t, useSSL, false)
}
