package recordfmt

import (
	"io"
	"os"

	"golang.org/x/term"
)

func IsTerminal(w io.Writer) bool {
	if file, ok := w.(*os.File); ok {
		return term.IsTerminal(int(file.Fd()))
	}

	return false
}
