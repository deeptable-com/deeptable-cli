// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/deeptable-cli/internal/mocktest"
)

func TestFilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "files", "retrieve",
			"--api-key", "string",
			"--file-id", "file_01kfxgjd94fn9stqm414vjb0s8",
		)
	})
}

func TestFilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "files", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "file_01kfxgjd94fn9stqm414vjb0s8",
			"--limit", "20",
		)
	})
}

func TestFilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "files", "delete",
			"--api-key", "string",
			"--file-id", "file_01kfxgjd94fn9stqm414vjb0s8",
		)
	})
}

func TestFilesDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "files", "download",
			"--api-key", "string",
			"--file-id", "file_01kfxgjd94fn9stqm414vjb0s8",
			"--output", "/dev/null",
		)
	})
}

func TestFilesUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "files", "upload",
			"--api-key", "string",
			"--file", "Example data",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("file: Example data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "files", "upload",
			"--api-key", "string",
		)
	})
}
