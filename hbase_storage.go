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
// HBase 虽支持 checkAndPut 原子操作，但本实现尚未接入真实 HBase 服务端，
// 仅作为非 CAS 存储的占位实现，不可用于分布式锁。真正接入时需用 checkAndPut 实现 CAS。
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
		storage.CapabilityReliableTime,
		storage.CapabilityWatch,
	}
}

func (x *HBaseStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}

func (x *HBaseStorage) Init(ctx context.Context) error {
	return nil
}

func (x *HBaseStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *HBaseStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *HBaseStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
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
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}
