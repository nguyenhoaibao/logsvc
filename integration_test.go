// +build integration

package logsvc_test

import (
	"flag"
	"testing"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/nguyenhoaibao/logsvc"
	"github.com/nguyenhoaibao/logsvc/pb"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

var (
	addr = flag.String("dbaddr", "127.0.0.1:5432", "DB addr")
	user = flag.String("dbuser", "postgres", "DB user")
	pw   = flag.String("dbpasswd", "mypostgrespw", "DB password")
	name = flag.String("dbname", "postgres", "DB name")
)

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*logsvc.Log)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func setup(tb testing.TB) (*pg.DB, func()) {
	db := pg.Connect(&pg.Options{
		Addr:     *addr,
		User:     *user,
		Password: *pw,
		Database: *name,
	})

	if err := createSchema(db); err != nil {
		tb.Fatalf("could not create schema: %v", err)
	}

	teardown := func() {
		if err := db.Close(); err != nil {
			tb.Fatalf("could not close DB: %v", err)
		}
	}
	return db, teardown
}

func TestService(t *testing.T) {
	db, teardown := setup(t)
	defer teardown()

	repo := logsvc.NewRepository(db)
	svc := logsvc.NewService(repo)

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

	_, err := svc.Write(ctx, logReq)
	assert.NoError(t, err)

	getReq := &pb.GetRequest{
		ClientIp: clientIP,
		ServerIp: serverIP,
		Tags: &pb.Tags{
			Tags: tags,
		},
	}

	results, err := svc.Get(ctx, getReq)
	assert.NoError(t, err)
	assert.Equal(t, &pb.GetResponse{
		Logs: []*pb.Log{
			{
				ClientIp: clientIP,
				ServerIp: serverIP,
				Tags: &pb.Tags{
					Tags: tags,
				},
				Msg: msg,
			},
		},
	}, results)
}

func BenchmarkServiceWrite(b *testing.B) {
	db, teardown := setup(b)
	defer teardown()

	repo := logsvc.NewRepository(db)
	svc := logsvc.NewService(repo)

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

	for n := 0; n < b.N; n++ {
		_, _ = svc.Write(ctx, logReq)
	}
}
