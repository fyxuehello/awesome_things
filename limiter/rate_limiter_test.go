package limiter

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRateLimiter(t *testing.T) {
	tokenBucket := RateLimiter()
	available := TakeAvailable(false, tokenBucket)
	assert.Equal(t, available, true)
}
