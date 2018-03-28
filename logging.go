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

// NewLoggingMiddleware returns a logging service middleware.
func NewLoggingMiddleware(l *zap.Logger, s pb.LogServiceServer) pb.LogServiceServer {
	return &loggingMiddleware{l, s}
}

func (l *loggingMiddleware) Write(ctx context.Context, req *pb.Log) (resp *pb.WriteResponse, err error) {
	defer func(begin time.Time) {
		logger := l.logger.With(
			zap.String("method", "write"),
			zap.String("req_client_ip", req.ClientIp),
			zap.String("req_server_ip", req.ServerIp),
			zap.String("req_tags", req.Tags.String()),
			zap.String("req_msg", req.Msg),
			zap.Duration("took", time.Since(begin)),
		)
		if err != nil {
			logger.Error("write error: %+v", zap.Error(err))
			return
		}
		logger.With(zap.Bool("resp", resp.Success)).Info("write success")
	}(time.Now())

	resp, err = l.LogServiceServer.Write(ctx, req)
	return
}

func (l *loggingMiddleware) Get(ctx context.Context, req *pb.GetRequest) (resp *pb.GetResponse, err error) {
	defer func(begin time.Time) {
		logger := l.logger.With(
			zap.String("method", "get"),
			zap.String("req_client_ip", req.ClientIp),
			zap.String("req_server_ip", req.ServerIp),
			zap.String("req_tags", req.Tags.String()),
			zap.Duration("took", time.Since(begin)),
		)
		if err != nil {
			logger.Error("get error: %+v", zap.Error(err))
			return
		}
		logger.Info("get success")
	}(time.Now())

	resp, err = l.LogServiceServer.Get(ctx, req)
	return
}
