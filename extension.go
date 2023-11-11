package prompt

import (
	k6modules "go.k6.io/k6/js/modules"
)

func init() {
	k6modules.Register("k6/x/prompt", new(Prompt))
}
