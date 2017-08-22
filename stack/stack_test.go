package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadStackOK(t *testing.T) {
	_, err := Read("test", "./fixtures/ok.yml")
	require.NoError(t, err)
}

func TestReadStackReadFails(t *testing.T) {
	_, err := Read("test", "./fixtures/invalid")
	require.Error(t, err)
}

func TestReadStackCreateFails(t *testing.T) {
	_, err := Read("test", "./fixtures/bad-link")
	require.Error(t, err)
}

func TestReadStackValidateFails(t *testing.T) {
	_, err := Read("test", "./fixtures/bad-port-container")
	require.Error(t, err)
}
