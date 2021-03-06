package longpoll // import "github.com/derad6709org/vksdk/v2/longpoll-bot"

import (
	"context"

	"github.com/derad6709org/vksdk/v2/internal"
)

// TsFromContext returns the ts from context.
func TsFromContext(ctx context.Context) int {
	return ctx.Value(internal.LongPollTsKey).(int)
}
