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
			t,
			"--api-key", "string",
			"structured-sheets", "create",
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
			t, pipeData,
			"--api-key", "string",
			"structured-sheets", "create",
		)
	})
}

func TestStructuredSheetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"structured-sheets", "retrieve",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
		)
	})
}

func TestStructuredSheetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"structured-sheets", "list",
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
			t,
			"--api-key", "string",
			"structured-sheets", "delete",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
		)
	})
}

func TestStructuredSheetsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"structured-sheets", "cancel",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
		)
	})
}

func TestStructuredSheetsDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"structured-sheets", "download",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
			"--format", "sqlite",
			"--output", "/dev/null",
		)
	})
}
