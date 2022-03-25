/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cl

import (
	"math/big"

	"fmt"

	"github.com/go-redis/redis"
)

// ReceiverRecordManager manages receiver records
// tied to particular nyms.
type ReceiverRecordManager interface {
	// Store stores the nym and the corresponding ReceiverRecord,
	// returning error in case the data was not successfully stored.
	Store(*big.Int, *ReceiverRecord) error

	// Load loads the ReceiverRecord associated with the given
	// nym, returning an error in case no record was found, or
	// in case of error in the interaction with the
	// storage backend.
	Load(*big.Int) (*ReceiverRecord, error)
}

// RedisClient wraps a redis client in order to interact with the
// redis database for management of receiver records.
type RedisClient struct {
	*redis.Client
}

// NewRedisClient accepts an instance of redis.Client and returns
// an instance of RedisClient.
func NewRedisClient(c *redis.Client) *RedisClient {
	return &RedisClient{
		Client: c,
	}
}

func (m *RedisClient) Store(nym *big.Int, r *ReceiverRecord) error {
	return m.Set(nym.String(), r, 0).Err()
}

func (m *RedisClient) Load(nym *big.Int) (*ReceiverRecord, error) {
	r, err := m.Get(nym.String()).Result()
	if err != nil {
		return nil, err
	}
	var rec ReceiverRecord
	rec.UnmarshalBinary([]byte(r))

	return &rec, nil
}

// MockRecordManager is a mock implementation of the ReceiverRecordManager
// interface. It stores key-value pairs of nyms and corresponding
// receiver records in a map.
type MockRecordManager struct {
	data map[string]ReceiverRecord
}

// NewMockRecordManager initializes the map that will hold the data.
func NewMockRecordManager() *MockRecordManager {
	return &MockRecordManager{
		data: make(map[string]ReceiverRecord),
	}
}

func (rm *MockRecordManager) Load(nym *big.Int) (*ReceiverRecord, error) {
	r, present := rm.data[nym.String()]
	if !present {
		return nil, fmt.Errorf("record does not exist")
	}

	return &r, nil
}

func (rm *MockRecordManager) Store(nym *big.Int, r *ReceiverRecord) error {
	rm.data[nym.String()] = *r
	return nil
}
