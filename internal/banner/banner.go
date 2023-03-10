package banner

import (
	"fmt"

	"github.com/pkg/errors"

	_banner "moul.io/banner"
)

func checkFormat(format string) error {
	if format != "plain" {
		return fmt.Errorf("unsupported format: %s", format)
	}

	return nil
}

func Render(input, format string) (string, error) {
	err := checkFormat(format)
	if err != nil {
		return "", errors.Wrap(err, "unable to render")
	}

	return _banner.Inline(input), nil
}
