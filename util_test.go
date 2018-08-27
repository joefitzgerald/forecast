package forecast_test

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(name string) string {
	f, _ := ioutil.ReadFile(fmt.Sprintf("%s%s%s", "testdata", string(os.PathSeparator), name))
	return string(f)
}
