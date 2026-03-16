// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/deeptable-com/deeptable-cli/internal/mocktest"
)

func TestStructuredSheetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets", "create",
			"--api-key", "string",
			"--file-id", "file_01h45ytscbebyvny4gc8cr8ma2",
			"--sheet-name", "[Sheet1, Financials]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file_id: file_01h45ytscbebyvny4gc8cr8ma2\n" +
			"sheet_names:\n" +
			"  - Sheet1\n" +
			"  - Financials\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "structured-sheets", "create",
			"--api-key", "string",
		)
	})
}

func TestStructuredSheetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets", "retrieve",
			"--api-key", "string",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
		)
	})
}

func TestStructuredSheetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--after", "ss_01kfxgjd94fn9stqm42nejb627",
			"--limit", "20",
		)
	})
}

func TestStructuredSheetsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets", "delete",
			"--api-key", "string",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
		)
	})
}

func TestStructuredSheetsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets", "cancel",
			"--api-key", "string",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
		)
	})
}

func TestStructuredSheetsDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets", "download",
			"--api-key", "string",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
			"--format", "sqlite",
			"--output", "/dev/null",
		)
	})
}
