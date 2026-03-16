// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/deeptable-com/deeptable-cli/internal/mocktest"
)

func TestStructuredSheetsTablesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets:tables", "retrieve",
			"--api-key", "string",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
			"--table-id", "tbl_01kfxgjd94fn9stqm45rqr2pnz",
		)
	})
}

func TestStructuredSheetsTablesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets:tables", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
			"--after", "tbl_01kfxgjd94fn9stqm45rqr2pnz",
			"--limit", "20",
		)
	})
}

func TestStructuredSheetsTablesDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "structured-sheets:tables", "download",
			"--api-key", "string",
			"--structured-sheet-id", "ss_01kfxgjd94fn9stqm42nejb627",
			"--table-id", "tbl_01kfxgjd94fn9stqm45rqr2pnz",
			"--format", "parquet",
			"--output", "/dev/null",
		)
	})
}
