package circumstances

import (
	"context"
	"errors"
	"testing"
)

var NotInCtx error = errors.New("not in context")

var ErrTestEnded = errors.New("test ended")

func ContextFromTest(t *testing.T) (context.Context, context.CancelCauseFunc) {
	ctx := context.Background()

	if deadline, ok := t.Deadline(); ok {
		// with cancel cause will cancel this intermediate context
		ctx, _ = context.WithDeadline(ctx, deadline)
	}

	return context.WithCancelCause(ctx)
}
