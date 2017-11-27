package framework

import (
	"github.com/cenkalti/backoff"
)

func DefaultBackOffFactory() func() backoff.BackOff {
	return func() backoff.BackOff {
		b := backoff.NewExponentialBackOff()
		b.MaxElapsedTime = 0
		return backoff.WithMaxTries(b, 7)
	}
}
