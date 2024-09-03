package workers

import "context"

type Outbox interface {
	Start(ctx context.Context)
}
