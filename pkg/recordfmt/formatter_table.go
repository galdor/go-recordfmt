package recordfmt

import (
	"fmt"
)

type TableColumnAlignment string

const (
	TableColumnAlignmentLeft  TableColumnAlignment = "left"
	TableColumnAlignmentRight TableColumnAlignment = "right"
)

type TableFormatter struct {
	Header  []string
	Columns []TableColumn
}

type TableColumn struct {
	Alignment TableColumnAlignment
}

func (f *TableFormatter) FormatRecordRows(
	records []Record, opts *FormattingOptions,
) error {
	rows := make([][]string, len(records))
	for i, record := range records {
		rows[i] = record.RecordFields()
	}

	widths := f.columnWidths(rows)

	if opts.PrintHeader {
		for i, column := range f.Columns {
			if i > 0 {
				if _, err := fmt.Fprintf(opts.HeaderOutput, "  "); err != nil {
					return err
				}
			}

			fmtString := "%-*s"
			if column.Alignment == TableColumnAlignmentRight {
				fmtString = "%*s"
			}

			label := fmt.Sprintf(fmtString, widths[i], f.Header[i])
			if _, err := fmt.Fprint(opts.HeaderOutput, label); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintln(opts.HeaderOutput); err != nil {
			return err
		}
	}

	for _, row := range rows {
		for j, s := range row {
			column := f.Columns[j]

			if j > 0 {
				fmt.Printf("  ")
			}

			fmtString := "%-*s"
			if column.Alignment == TableColumnAlignmentRight {
				fmtString = "%*s"
			}

			if _, err := fmt.Printf(fmtString, widths[j], s); err != nil {
				return err
			}
		}

		fmt.Println("")
	}

	return nil
}

func (f *TableFormatter) columnWidths(rows [][]string) []int {
	widths := make([]int, len(f.Header))

	for i, label := range f.Header {
		widths[i] = len(label)
	}

	for _, row := range rows {
		for j, s := range row {
			if len(s) > widths[j] {
				widths[j] = len(s)
			}
		}
	}

	return widths
}
