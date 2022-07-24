package pytest

import (
	"github.com/test-instructor/cheetah/server/hrp/internal/builtin"
	"github.com/test-instructor/cheetah/server/hrp/internal/sdk"
)

func RunPytest(args []string) error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp pytest",
	})

	args = append([]string{"run"}, args...)
	return builtin.ExecPython3Command("httprunner", args...)
}
