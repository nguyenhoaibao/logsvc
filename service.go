package logsvc

import (
	"github.com/nguyenhoaibao/logsvc/pb"
	context "golang.org/x/net/context"
)

// NewService returns a log service.
func NewService(repo Repository) pb.LogServiceServer {
	return &service{repo}
}

type service struct {
	repo Repository
}

func (s *service) Write(ctx context.Context, req *pb.Log) (*pb.WriteResponse, error) {
	l := logFromProtoMsg(req)
	if err := s.repo.Save(ctx, l); err != nil {
		return nil, err
	}
	return &pb.WriteResponse{
		Success: true,
	}, nil
}

func (s *service) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	logs, err := s.repo.Get(ctx, req.ClientIp, req.ServerIp, req.Tags.GetTags())
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return &pb.GetResponse{
			Logs: []*pb.Log{},
		}, nil
	}

	resp := make([]*pb.Log, 0, len(logs))
	for _, log := range logs {
		resp = append(resp, logToProtoMsg(&log))
	}
	return &pb.GetResponse{resp}, nil
}

func logFromProtoMsg(req *pb.Log) *Log {
	return &Log{
		ClientIP: req.ClientIp,
		ServerIP: req.ServerIp,
		Msg:      req.Msg,
		Tags:     req.Tags.GetTags(),
	}
}

func logToProtoMsg(l *Log) *pb.Log {
	return &pb.Log{
		ClientIp: l.ClientIP,
		ServerIp: l.ServerIP,
		Tags: &pb.Tags{
			Tags: l.Tags,
		},
		Msg: l.Msg,
	}
}
