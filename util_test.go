package forecast_test

import (
	"fmt"
	"os"
)

func ReadFile(name string) string {
	f, _ := os.ReadFile(fmt.Sprintf("%s%s%s", "testdata", string(os.PathSeparator), name))
	return string(f)
}
