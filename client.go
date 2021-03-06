package reflex

import (
	"context"
	"strconv"

	"github.com/luno/reflex/reflexpb"
)

// StreamClientPB defines a common interface for reflex stream gRPC
// generated implementations.
type StreamClientPB interface {
	Recv() (*reflexpb.Event, error)
}

// WrapStreamPB wraps a gRPC client's stream method and returns a StreamFunc.
func WrapStreamPB(wrap func(context.Context, *reflexpb.StreamRequest) (
	StreamClientPB, error)) StreamFunc {
	return func(ctx context.Context, after string, opts ...StreamOption) (StreamClient, error) {
		optionspb, err := optsToProto(opts)
		if err != nil {
			return nil, err
		}

		afterInt, _ := strconv.ParseInt(after, 10, 64) // Remove after migration

		cspb, err := wrap(ctx, &reflexpb.StreamRequest{
			After:    after,
			AfterInt: afterInt,
			Options:  optionspb,
		})
		if err != nil {
			return nil, err
		}

		return streamClientFromProto(cspb), nil
	}
}
