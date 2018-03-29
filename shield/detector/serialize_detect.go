package detector

import (
	"bytes"
	"fmt"

	"github.com/yaptide/converter/format"
)

func Serialize(detectors []Detector) string {
	w := &bytes.Buffer{}

	const collumnsPerRow = 6
	for i, detector := range detectors {

		if i > 0 {
			fmt.Fprint(w, "\n")
		}

		fmt.Fprintf(w, "%-10s", detector.ScoringType)

		for j, arg := range detector.Arguments {
			if j > 0 && j%collumnsPerRow == 0 {
				fmt.Fprintf(w, "\n%10s", "")
			}

			switch entry := arg.(type) {
			case float64:
				fmt.Fprint(w, format.FloatToFixedWidthString(entry, 10))
			case int64:
				fmt.Fprintf(w, "%10d", entry)
			case string:
				fmt.Fprintf(w, "%-10s", entry)
			}
		}

	}

	return w.String()
}
