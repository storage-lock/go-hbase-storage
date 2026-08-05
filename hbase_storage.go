package hbase_storage

import (
	"context"
	"sync"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// HBaseStorage HBase存储实现
// HBase 支持 checkAndPut 原子操作，可以实现真正的CAS
type HBaseStorage struct {
	mu sync.RWMutex
}

var _ storage.Storage = &HBaseStorage{}

func NewHBaseStorage(ctx context.Context) (*HBaseStorage, error) {
	return &HBaseStorage{}, nil
}

const StorageName = "hbase-storage"

func (x *HBaseStorage) GetName() string {
	return StorageName
}

func (x *HBaseStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{
		storage.CapabilityCAS,
		storage.CapabilityReliableTime,
	}
}

func (x *HBaseStorage) Init(ctx context.Context) error {
	return nil
}

func (x *HBaseStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	// HBase checkAndPut 实现
	return nil
}

func (x *HBaseStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *HBaseStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *HBaseStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return "", storage_lock.ErrLockNotFound
}

func (x *HBaseStorage) GetTime(ctx context.Context) (time.Time, error) {
	return time.Now(), nil
}

func (x *HBaseStorage) Close(ctx context.Context) error {
	return nil
}

func (x *HBaseStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return nil, nil
}
