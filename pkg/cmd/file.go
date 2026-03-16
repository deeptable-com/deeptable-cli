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

var filesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get metadata for a specific file.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "The unique identifier of the file.",
			Required: true,
		},
	},
	Action:          handleFilesRetrieve,
	HideHelpCommand: true,
}

var filesList = cli.Command{
	Name:    "list",
	Usage:   "List all files uploaded by the current user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "after",
			Usage:     "A cursor for pagination. Use the `last_id` from a previous response to fetch the next page.",
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of files to return.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleFilesList,
	HideHelpCommand: true,
}

var filesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a file. This cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "The unique identifier of the file.",
			Required: true,
		},
	},
	Action:          handleFilesDelete,
	HideHelpCommand: true,
}

var filesDownload = cli.Command{
	Name:    "download",
	Usage:   "Download the original uploaded file content.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "The unique identifier of the file.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleFilesDownload,
	HideHelpCommand: true,
}

var filesUpload = cli.Command{
	Name:    "upload",
	Usage:   "Upload an Excel spreadsheet file for later processing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file",
			Usage:    "The spreadsheet file to upload",
			Required: true,
			BodyPath: "file",
		},
	},
	Action:          handleFilesUpload,
	HideHelpCommand: true,
}

func handleFilesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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
	_, err = client.Files.Get(ctx, cmd.Value("file-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "files retrieve", obj, format, transform)
}

func handleFilesList(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.FileListParams{}

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
		_, err = client.Files.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "files list", obj, format, transform)
	} else {
		iter := client.Files.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "files list", iter, format, transform, maxItems)
	}
}

func handleFilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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
	_, err = client.Files.Delete(ctx, cmd.Value("file-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "files delete", obj, format, transform)
}

func handleFilesDownload(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	response, err := client.Files.Download(ctx, cmd.Value("file-id").(string), options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handleFilesUpload(ctx context.Context, cmd *cli.Command) error {
	client := deeptable.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := deeptable.FileUploadParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Files.Upload(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "files upload", obj, format, transform)
}
