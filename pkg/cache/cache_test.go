package cache_test

import (
	"testing"
	"time"

	"github.com/di-dao/sonr/pkg/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache_SetGet(t *testing.T) {
	type MyString string
	type Data struct {
		Value int
	}

	cache := cache.New[MyString, Data](5*time.Minute, 10*time.Minute)

	key := MyString("testKey")
	expectedValue := Data{Value: 42}

	cache.Set(key, expectedValue)
	retrievedValue, ok := cache.Get(key)

	require.True(t, ok, "Expected to retrieve a value from the cache")
	require.NotNil(t, retrievedValue, "Expected a non-nil value")
	assert.Equal(t, expectedValue, *retrievedValue, "Retrieved value does not match the expected value")
}

func TestCache_Get_NonExistentKey(t *testing.T) {
	type MyString string
	type Data struct {
		Value int
	}

	cache := cache.New[MyString, Data](5*time.Minute, 10*time.Minute)

	key := MyString("nonExistentKey")
	retrievedValue, ok := cache.Get(key)

	require.False(t, ok, "Expected not to retrieve a value from the cache for a non-existent key")
	require.Nil(t, retrievedValue, "Expected a nil value for a non-existent key")
}

func TestCache_Set_Overwrite(t *testing.T) {
	type MyString string
	type Data struct {
		Value int
	}

	cache := cache.New[MyString, Data](5*time.Minute, 10*time.Minute)

	key := MyString("testKey")
	initialValue := Data{Value: 42}
	updatedValue := Data{Value: 100}

	cache.Set(key, initialValue)
	cache.Set(key, updatedValue)

	retrievedValue, ok := cache.Get(key)

	require.True(t, ok, "Expected to retrieve a value from the cache")
	require.NotNil(t, retrievedValue, "Expected a non-nil value")
	assert.Equal(t, updatedValue, *retrievedValue, "Retrieved value does not match the updated value")
}

func TestCache_Set_Get_OtherType(t *testing.T) {
	type MyString string
	type Data struct {
		Value int
	}

	cache := cache.New[MyString, Data](5*time.Minute, 10*time.Minute)

	key := MyString("testKey")
	intVal := 42

	cache.Set(key, intVal)

	retrievedValue, ok := cache.Get(key)

	require.False(t, ok, "Expected not to retrieve a value from the cache for a key with a wrong type")
	require.Nil(t, retrievedValue, "Expected a nil value for a key with a wrong type")
}

func TestCache_Set_WithExpiration(t *testing.T) {
	type MyString string
	type Data struct {
		Value int
	}

	expiration := 1 * time.Second
	cache := cache.New[MyString, Data](expiration, 10*time.Minute)

	key := MyString("testKey")
	expectedValue := Data{Value: 42}

	cache.Set(key, expectedValue)
	time.Sleep(2 * time.Second)

	retrievedValue, ok := cache.Get(key)

	require.False(t, ok, "Expected not to retrieve a value from the cache after expiration")
	require.Nil(t, retrievedValue, "Expected a nil value after expiration")
}
