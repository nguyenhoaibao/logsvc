package logsvc

import (
	"time"

	"github.com/nguyenhoaibao/logsvc/pb"
	"go.uber.org/zap"
	context "golang.org/x/net/context"
)

type loggingMiddleware struct {
	logger *zap.Logger
	pb.LogServiceServer
}

// NewLoggingService returns a logging service middleware.
func NewLoggingService(l *zap.Logger) ServiceMiddleware {
	return func(s pb.LogServiceServer) pb.LogServiceServer {
		return &loggingMiddleware{l, s}
	}
}

func (l *loggingMiddleware) Write(ctx context.Context, req *pb.Log) (resp *pb.WriteResponse, err error) {
	defer func(begin time.Time) {
		logger := l.logger.With(
			zap.String("method", "write"),
			zap.Duration("took", time.Since(begin)),
		)
		if err != nil {
			logger.Error("write error: %+v", zap.Error(err))
			return
		}
		logger.Info("write success")
	}(time.Now())

	return l.LogServiceServer.Write(ctx, req)
}

func (l *loggingMiddleware) Get(ctx context.Context, req *pb.GetRequest) (resp *pb.GetResponse, err error) {
	defer func(begin time.Time) {
		logger := l.logger.With(
			zap.String("method", "get"),
			zap.Duration("took", time.Since(begin)),
		)
		if err != nil {
			logger.Error("get error: %+v", zap.Error(err))
			return
		}
		logger.Info("get success")
	}(time.Now())

	return l.LogServiceServer.Get(ctx, req)

}
