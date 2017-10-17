package forecast_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestForecast(t *testing.T) {
	log.SetFlags(log.Llongfile)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forecast Suite")
}

func ReadFile(name string) string {
	f, _ := ioutil.ReadFile(fmt.Sprintf("%s%s%s", "testdata", string(os.PathSeparator), name))
	return string(f)
}
