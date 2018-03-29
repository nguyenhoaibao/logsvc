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
		clientIP = "127.0.0.1"
		serverIP = "127.0.1.1"
		tags     = map[string]string{"key": "val"}
		msg      = "log message"

		ctx    = context.Background()
		logReq = &pb.Log{
			ClientIp: clientIP,
			ServerIp: serverIP,
			Tags: &pb.Tags{
				Tags: tags,
			},
			Msg: msg,
		}
	)

	mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).
		Do(func(_ context.Context, log *logsvc.Log) {
			assert.Equal(t, clientIP, log.ClientIP)
			assert.Equal(t, serverIP, log.ServerIP)
			assert.Equal(t, tags, log.Tags)
			assert.Equal(t, msg, log.Msg)
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
		clientIP = "127.0.0.1"
		serverIP = "127.0.1.1"
		tags     = map[string]string{"key": "val"}

		ctx = context.Background()
		req = &pb.GetRequest{
			ClientIp: clientIP,
			ServerIp: serverIP,
			Tags: &pb.Tags{
				Tags: tags,
			},
		}
	)

	mockRepo.EXPECT().Get(ctx, clientIP, serverIP, tags)

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
