package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
    once sync.Once
    // 全局变量，方便其它包直接调用已初始化好的 S 实例.
    S *datastore
)

type IStore interface {
	Users() UserStore
}

type datastore struct {
	db *gorm.DB
}

var _ IStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
    once.Do(func() {
        S = &datastore{db}
    })
    return S
}

func (ds *datastore) Users() UserStore {
    return newUser(ds.db)
}
