package kvdb

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/sashabaranov/go-openai"
	"go.etcd.io/bbolt"
)

var defaultBucket = "default"

type ChatContext struct {
	Time     time.Time
	Messages []openai.ChatCompletionMessage
}

type QueryParams struct {
	Bucket string
	From   time.Time
	To     time.Time
	Limit  int
}
type SaveParmas struct {
	Bucket string

	ChatContext
}

type Interface interface {
	Keys() []string
	Save(p SaveParmas) error
	Query(p QueryParams) []ChatContext
	Close()
}

var _ Interface = &kvImpl{}

type kvImpl struct {
	db *bbolt.DB
}

func (c *kvImpl) Keys() (result []string) {
	c.db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			result = append(result, string(name))
			return nil
		})
	})
	return
}

func (c *kvImpl) Save(p SaveParmas) error {
	if p.Bucket == "" {
		p.Bucket = defaultBucket
	}

	return c.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(p.Bucket))
		if err != nil {
			return err
		}

		data, err := json.Marshal(p.Messages)
		if err != nil {
			return err
		}

		return b.Put([]byte(p.Time.Format(time.RFC3339)), data)
	})
}

func (c *kvImpl) Close() {
	if err := c.db.Close(); err != nil {
		panic(err)
	}
}

func (c *kvImpl) Query(p QueryParams) (result []ChatContext) {
	if p.Bucket == "" {
		p.Bucket = defaultBucket
	}

	if p.To.IsZero() {
		p.To = time.Now()
	}

	if p.Limit == 0 {
		p.Limit = 3
	}

	c.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(p.Bucket))
		if b == nil {
			return nil
		}

		cur := b.Cursor()

		min := []byte(p.From.Format(time.RFC3339))
		max := []byte(p.To.Format(time.RFC3339))

		for k, v := cur.Seek(min); k != nil && bytes.Compare(k, max) <= 0 && p.Limit > 0; k, v = cur.Next() {
			key, err := time.Parse(time.RFC3339, string(k))
			if err != nil {
				continue
			}

			var msg []openai.ChatCompletionMessage

			if err := json.Unmarshal(v, &msg); err != nil {
				continue
			}

			result = append(result, ChatContext{key, msg})

			p.Limit--
		}

		return nil
	})

	return
}

func New(p string) (Interface, error) {
	db, err := bbolt.Open(p, 0666, nil)
	if err != nil {
		return nil, err
	}

	return &kvImpl{db}, nil
}
