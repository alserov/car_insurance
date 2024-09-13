package limiter

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestLeakyBucketSuite(t *testing.T) {
	suite.Run(t, new(LeakyBucketSuite))
}

type LeakyBucketSuite struct {
	suite.Suite

	lim int
}

func (s *LeakyBucketSuite) SetupTest() {
	s.lim = 5000
}

func (s *LeakyBucketSuite) TestTooManyRequests() {
	limiter := newLeakyBucket(s.lim)

	reqAmount := 5500
	timeout := 10

	expectedCanceledReqCount := (reqAmount - s.lim) / int(time.Duration(timeout)*time.Millisecond/refresh)
	canceledReqCount := 0

	for i := 0; i < reqAmount; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))

		if limiter.Limit(ctx) {
			canceledReqCount++
		}

		cancel()
	}

	s.Require().Equal(expectedCanceledReqCount, canceledReqCount)
}

func (s *LeakyBucketSuite) TestDefaultThroughput() {
	limiter := newLeakyBucket(s.lim)

	timeout := 10

	for i := 0; i < s.lim; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))

		s.Require().False(limiter.Limit(ctx))

		cancel()
	}
}
