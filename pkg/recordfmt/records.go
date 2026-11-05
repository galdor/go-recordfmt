package recordfmt

import (
	"os"
)

type Record interface {
	RecordFields() []string
}

func FormatRecords(
	records []Record, formatter Formatter, opts *FormattingOptions,
) error {
	if opts == nil {
		opts = &FormattingOptions{}
	}

	if opts.DataOutput == nil {
		opts.DataOutput = os.Stdout
	}

	if opts.HeaderOutput == nil {
		opts.HeaderOutput = os.Stderr
	}

	if opts.PrintHeader {
		if !IsTerminal(opts.HeaderOutput) && !opts.ForcePrintHeader {
			opts.PrintHeader = false
		}
	}

	return formatter.FormatRecordRows(records, opts)
}
