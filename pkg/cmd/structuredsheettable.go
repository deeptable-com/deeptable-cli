// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/deeptable-com/deeptable-cli/internal/apiquery"
	"github.com/deeptable-com/deeptable-cli/internal/requestflag"
	"github.com/stainless-sdks/deeptable-go"
	"github.com/stainless-sdks/deeptable-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var structuredSheetsTablesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get details of a specific table extracted from the structured sheet. Only\navailable when conversion status is 'completed'.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "table-id",
			Usage:    "The unique identifier of the table.",
			Required: true,
		},
	},
	Action:          handleStructuredSheetsTablesRetrieve,
	HideHelpCommand: true,
}

var structuredSheetsTablesList = cli.Command{
	Name:    "list",
	Usage:   "List all tables extracted from the structured sheet. Only available when\nconversion status is 'completed'.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "after",
			Usage:     "A cursor for pagination. Use the `last_id` from a previous response to fetch the next page of results.",
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of tables to return per page.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleStructuredSheetsTablesList,
	HideHelpCommand: true,
}

var structuredSheetsTablesDownload = cli.Command{
	Name:    "download",
	Usage:   "Download the table data in the specified format.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "table-id",
			Usage:    "The unique identifier of the table.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "The format to download the table data in.",
			Required:  true,
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleStructuredSheetsTablesDownload,
	HideHelpCommand: true,
}

func handleStructuredSheetsTablesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("table-id") && len(unusedArgs) > 0 {
		cmd.Set("table-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.StructuredSheetTableGetParams{
		StructuredSheetID: cmd.Value("structured-sheet-id").(string),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.StructuredSheets.Tables.Get(
		ctx,
		cmd.Value("table-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "structured-sheets:tables retrieve", obj, format, transform)
}

func handleStructuredSheetsTablesList(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("structured-sheet-id") && len(unusedArgs) > 0 {
		cmd.Set("structured-sheet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.StructuredSheetTableListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.StructuredSheets.Tables.List(
			ctx,
			cmd.Value("structured-sheet-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "structured-sheets:tables list", obj, format, transform)
	} else {
		iter := client.StructuredSheets.Tables.ListAutoPaging(
			ctx,
			cmd.Value("structured-sheet-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "structured-sheets:tables list", iter, format, transform, maxItems)
	}
}

func handleStructuredSheetsTablesDownload(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("table-id") && len(unusedArgs) > 0 {
		cmd.Set("table-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.StructuredSheetTableDownloadParams{
		StructuredSheetID: cmd.Value("structured-sheet-id").(string),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	response, err := client.StructuredSheets.Tables.Download(
		ctx,
		cmd.Value("table-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
