package logsvc

import (
	"encoding/json"

	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	context "golang.org/x/net/context"
)

// Repository contains methods to interact with the datastore.
type Repository interface {
	Save(context.Context, *Log) error
	Get(ctx context.Context, clientIP, serverIP string, tags map[string]string) ([]Log, error)
}

// NewRepository returns a new repository.
func NewRepository(db *pg.DB) Repository {
	return &repo{db}
}

// Log is a entity model represents a record in the DB.
type Log struct {
	ID       int64
	ClientIP string
	ServerIP string
	Tags     map[string]string
	Msg      string
}

type repo struct {
	db *pg.DB
}

func (r *repo) Save(ctx context.Context, l *Log) error {
	r.db.WithContext(ctx)

	if err := r.db.Insert(l); err != nil {
		return errors.Wrap(err, "repo: could not save log")
	}
	return nil
}

func (r *repo) Get(ctx context.Context, clientIP, serverIP string, tags map[string]string) ([]Log, error) {
	r.db.WithContext(ctx)

	var logs []Log

	q := r.db.Model(&logs)
	if clientIP != "" {
		q = q.Where("client_ip = ?", clientIP)
	}
	if serverIP != "" {
		q = q.Where("server_ip = ?", serverIP)
	}
	if tags != nil {
		json, err := json.Marshal(tags)
		if err != nil {
			return nil, errors.Wrapf(err, "repo: could not marshal json from tags: %v", string(json))
		}
		q = q.Where("tags @> ?", string(json))
	}
	if err := q.Select(); err != nil {
		return nil, errors.Wrap(err, "repo: could not get logs")
	}
	return logs, nil
}
