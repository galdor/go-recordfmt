package recordfmt

import "io"

type FormattingOptions struct {
	DataOutput io.Writer

	PrintHeader      bool
	ForcePrintHeader bool
	HeaderOutput     io.Writer
}

type Formatter interface {
	FormatRecordRows([]Record, *FormattingOptions) error
}
