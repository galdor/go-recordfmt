package recordfmt

import (
	"encoding/json"
	"fmt"
)

type JSONFormatter struct {
}

func (f *JSONFormatter) FormatRecordRows(
	records []Record, opts *FormattingOptions,
) error {
	encoder := json.NewEncoder(opts.DataOutput)

	for _, record := range records {
		if err := encoder.Encode(record); err != nil {
			return fmt.Errorf("cannot encode record: %w", err)
		}
	}

	return nil
}
