package banner

import (
	_banner "moul.io/banner"
)

func Render(input string) string {
	return _banner.Inline(input)
}
