// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/deeptable-cli/internal/apiquery"
	"github.com/stainless-sdks/deeptable-cli/internal/requestflag"
	"github.com/stainless-sdks/deeptable-go"
	"github.com/stainless-sdks/deeptable-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var structuredSheetsCreate = cli.Command{
	Name:    "create",
	Usage:   "Start converting a spreadsheet workbook into structured data. This initiates an\nasynchronous conversion process. Poll the returned resource using the `id` to\ncheck completion status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "The unique identifier of the file to convert.",
			Required: true,
			BodyPath: "file_id",
		},
		&requestflag.Flag[any]{
			Name:     "sheet-name",
			Usage:    "List of sheet names to convert. If None, all sheets will be converted.",
			BodyPath: "sheet_names",
		},
	},
	Action:          handleStructuredSheetsCreate,
	HideHelpCommand: true,
}

var structuredSheetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the status and details of a structured sheet conversion.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
	},
	Action:          handleStructuredSheetsRetrieve,
	HideHelpCommand: true,
}

var structuredSheetsList = cli.Command{
	Name:    "list",
	Usage:   "List all structured sheets conversions for the authenticated user. Results are\npaginated using cursor-based pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "after",
			Usage:     "A cursor for pagination. Use the `last_id` from a previous response to fetch the next page of results.",
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return per page.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleStructuredSheetsList,
	HideHelpCommand: true,
}

var structuredSheetsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a structured sheet conversion and its associated exports. This action\ncannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
	},
	Action:          handleStructuredSheetsDelete,
	HideHelpCommand: true,
}

var structuredSheetsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a structured sheet conversion that is in progress. Only jobs with status\n'queued' or 'in_progress' can be cancelled.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
	},
	Action:          handleStructuredSheetsCancel,
	HideHelpCommand: true,
}

var structuredSheetsDownload = cli.Command{
	Name:    "download",
	Usage:   "Download the structured data in the specified format. Only available when\nconversion status is 'completed'.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "structured-sheet-id",
			Usage:    "The unique identifier of the structured sheet conversion.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "The export format to download.",
			Default:   "sqlite",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleStructuredSheetsDownload,
	HideHelpCommand: true,
}

func handleStructuredSheetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.StructuredSheetNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.StructuredSheets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "structured-sheets create", obj, format, transform)
}

func handleStructuredSheetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("structured-sheet-id") && len(unusedArgs) > 0 {
		cmd.Set("structured-sheet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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
	_, err = client.StructuredSheets.Get(ctx, cmd.Value("structured-sheet-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "structured-sheets retrieve", obj, format, transform)
}

func handleStructuredSheetsList(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.StructuredSheetListParams{}

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
		_, err = client.StructuredSheets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "structured-sheets list", obj, format, transform)
	} else {
		iter := client.StructuredSheets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "structured-sheets list", iter, format, transform, maxItems)
	}
}

func handleStructuredSheetsDelete(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("structured-sheet-id") && len(unusedArgs) > 0 {
		cmd.Set("structured-sheet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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
	_, err = client.StructuredSheets.Delete(ctx, cmd.Value("structured-sheet-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "structured-sheets delete", obj, format, transform)
}

func handleStructuredSheetsCancel(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("structured-sheet-id") && len(unusedArgs) > 0 {
		cmd.Set("structured-sheet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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
	_, err = client.StructuredSheets.Cancel(ctx, cmd.Value("structured-sheet-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "structured-sheets cancel", obj, format, transform)
}

func handleStructuredSheetsDownload(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("structured-sheet-id") && len(unusedArgs) > 0 {
		cmd.Set("structured-sheet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.StructuredSheetDownloadParams{}

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

	response, err := client.StructuredSheets.Download(
		ctx,
		cmd.Value("structured-sheet-id").(string),
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
