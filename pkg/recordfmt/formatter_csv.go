package recordfmt

import (
	"encoding/csv"
)

type CSVFormatter struct {
	Header []string
}

func (f *CSVFormatter) FormatRecordRows(
	records []Record, opts *FormattingOptions,
) error {
	if opts.PrintHeader && f.Header != nil {
		csvWriter := csv.NewWriter(opts.HeaderOutput)

		if err := csvWriter.Write(f.Header); err != nil {
			return err
		}

		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			return err
		}
	}

	csvWriter := csv.NewWriter(opts.DataOutput)

	for _, record := range records {
		fields := record.RecordFields()

		if err := csvWriter.Write(fields); err != nil {
			return err
		}
	}

	csvWriter.Flush()
	return csvWriter.Error()
}
