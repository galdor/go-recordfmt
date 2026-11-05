package main

import (
	"io/fs"
	"os"
	"strconv"

	"go.n16f.net/program"
	"go.n16f.net/recordfmt/pkg/recordfmt"
)

type FileRecord struct {
	Name        string      `json:"name"`
	IsDirectory bool        `json:"is_directory"`
	Size        int64       `json:"size"`
	Mode        fs.FileMode `json:"mode"`
}

func (r *FileRecord) RecordFields() []string {
	name := r.Name
	if r.IsDirectory {
		name += "/"
	}

	return []string{
		name,
		strconv.FormatInt(r.Size, 10),
		"0" + strconv.FormatInt(int64(r.Mode), 8),
	}
}

func main() {
	p := program.NewProgram("ls", "list files as records")

	p.AddOptionalArgument("directory", "the directory to list")
	p.AddOption("f", "format", "format", "table",
		"the output format (csv, json, table)")

	p.SetMain(CMDMain)

	p.ParseCommandLine()
	p.Run()
}

func CMDMain(p *program.Program) {
	dirPath := "."
	if dirPathOpt := p.OptionalArgumentValue("directory"); dirPathOpt != nil {
		dirPath = *dirPathOpt
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		p.Fatal("cannot read directory %q: %v", dirPath, err)
	}

	var records []recordfmt.Record

	for _, entry := range entries {
		filename := entry.Name()

		info, err := entry.Info()
		if err != nil {
			p.Fatal("cannot retrieve information about %q: %w", filename, err)
		}

		record := FileRecord{
			Name:        filename,
			IsDirectory: info.IsDir(),
			Size:        info.Size(),
			Mode:        info.Mode().Perm(),
		}

		records = append(records, &record)
	}

	formatter := NewFormatter(p)

	opts := recordfmt.FormattingOptions{
		PrintHeader: true,
	}

	recordfmt.FormatRecords(records, formatter, &opts)
}

func NewFormatter(p *program.Program) recordfmt.Formatter {
	format := p.OptionValue("format")

	var formatter recordfmt.Formatter

	switch format {
	case "csv":
		formatter = &recordfmt.CSVFormatter{
			Header: []string{"name", "size", "mode"},
		}

	case "json":
		formatter = &recordfmt.JSONFormatter{}

	case "table":
		formatter = &recordfmt.TableFormatter{
			Header: []string{"NAME", "SIZE", "MODE"},
			Columns: []recordfmt.TableColumn{
				{},
				{Alignment: recordfmt.TableColumnAlignmentRight},
				{Alignment: recordfmt.TableColumnAlignmentRight},
			},
		}

	default:
		p.Fatal("unknown format %q", format)
	}

	return formatter
}
