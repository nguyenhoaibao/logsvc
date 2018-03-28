package logsvc_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nguyenhoaibao/logsvc"
	"github.com/nguyenhoaibao/logsvc/mocks"
	"github.com/nguyenhoaibao/logsvc/pb"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestService_Write(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	svc := logsvc.NewService(mockRepo)

	var (
		ctx    = context.Background()
		logReq = &pb.Log{
			ClientIp: "127.0.0.1",
			ServerIp: "127.0.1.1",
			Tags: &pb.Tags{
				Tags: map[string]string{
					"key": "val",
				},
			},
			Msg: "log message",
		}
	)

	mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).
		Do(func(_ context.Context, log *logsvc.Log) {
			assert.Equal(t, logReq.ClientIp, log.ClientIP)
			assert.Equal(t, logReq.ServerIp, log.ServerIP)
			assert.Equal(t, logReq.Tags.GetTags(), log.Tags)
			assert.Equal(t, logReq.Msg, log.Msg)
		})

	_, err := svc.Write(ctx, logReq)
	assert.NoError(t, err)
}

func TestService_WriteReturnsError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	svc := logsvc.NewService(mockRepo)

	dummyErr := errors.New("dummy error")
	mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(dummyErr)

	_, err := svc.Write(context.Background(), &pb.Log{})
	assert.Equal(t, dummyErr, err)
}

func TestService_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	svc := logsvc.NewService(mockRepo)

	var (
		ctx = context.Background()
		req = &pb.GetRequest{
			ClientIp: "127.0.0.1",
			ServerIp: "127.0.1.1",
			Tags: &pb.Tags{
				Tags: map[string]string{
					"key": "val",
				},
			},
		}
	)

	mockRepo.EXPECT().Get(ctx, req.ClientIp, req.ServerIp, req.Tags.GetTags())

	_, err := svc.Get(ctx, req)
	assert.NoError(t, err)
}

func TestService_GetReturnsError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	svc := logsvc.NewService(mockRepo)

	dummyErr := errors.New("dummy error")
	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, dummyErr)

	_, err := svc.Get(context.Background(), &pb.GetRequest{})
	assert.Equal(t, dummyErr, err)
}
