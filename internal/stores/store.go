package stores

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Store struct {
	rdb *redis.Client
	addr string
	ctx context.Context
}

func defaultStore() *Store{
	return &Store{
		// addr: "146.190.132.169:6001",
		addr: "0.0.0.0:6001",
		ctx: context.Background(),
	}
}

type StoreOption func(*Store)

func WithAddr(addr string) StoreOption {
	return func(s *Store) {
		s.addr = addr
	}
}

func NewIcefireStore(opts ...StoreOption) (*Store, error) {
	s := defaultStore()
	for _, opt := range opts {
		opt(s)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     s.addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	s.rdb = rdb
	return s, nil
}

func (s *Store) ExistsInList(key string, value string) (bool, error) {
	result, err := s.rdb.LRange(s.ctx, key, 0, -1).Result()
	if err != nil {
		return false, err
	}
	for _, v := range result {
		if v == value {
			return true, nil
		}
	}
	return false, nil
}

func (s *Store) ExistsInMap(key string, field string) (bool, error) {
	return s.rdb.HExists(s.ctx, key, field).Result()
}

func (s *Store) ExistsInSet(key string, member string) (bool, error) {
	return s.rdb.SIsMember(s.ctx, key, member).Result()
}

func (s *Store) GetList(key string) ([]string, error) {
	return s.rdb.LRange(s.ctx, key, 0, -1).Result()
}

func (s *Store) GetMap(key string) (map[string]string, error) {
	return s.rdb.HGetAll(s.ctx, key).Result()
}

func (s *Store) GetSet(key string) ([]string, error) {
	return s.rdb.SMembers(s.ctx, key).Result()
}

func (s *Store) DelListItem(key string, value string) (int64, error) {
	return s.rdb.LRem(s.ctx, key, 0, value).Result()
}

func (s *Store) DelMapItem(key string, field string) (int64, error) {
	return s.rdb.HDel(s.ctx, key, field).Result()
}

func (s *Store) DelSetItem(key string, member string) (int64, error) {
	return s.rdb.SRem(s.ctx, key, member).Result()
}

func (s *Store) AddListItem(key string, value string) error {
	return s.rdb.RPush(s.ctx, key, value).Err()
}

func (s *Store) AddMapItem(key string, field string, value string) error {
	return s.rdb.HSet(s.ctx, key, field, value).Err()
}

func (s *Store) AddSetItem(key string, member string) error {
	return s.rdb.SAdd(s.ctx, key, member).Err()
}
