package middleware

import "context"

type Wrapper func(ctx context.Context) context.Context

func WithWrappers(wrs ...Wrapper) {

}
