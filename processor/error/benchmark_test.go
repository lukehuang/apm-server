package error

import (
	"testing"

	"github.com/elastic/apm-server/config"
	"github.com/elastic/apm-server/tests/loader"
)

func BenchmarkEventWithFileLoading(b *testing.B) {
	processor := NewProcessor(config.Config{})
	for i := 0; i < b.N; i++ {
		data, _ := loader.LoadValidData("error")
		err := processor.Validate(data)
		if err != nil {
			panic(err)
		}

		processor.Transform(data)
	}
}

func BenchmarkEventFileLoadingOnce(b *testing.B) {
	processor := NewProcessor(config.Config{})
	data, _ := loader.LoadValidData("error")
	for i := 0; i < b.N; i++ {
		err := processor.Validate(data)
		if err != nil {
			panic(err)
		}

		processor.Transform(data)
	}
}
