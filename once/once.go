package once

import (
	"context"
)

type Once interface {
	Do(ctx context.Context) error
}
