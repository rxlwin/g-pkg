package mc

import (
	"hash/crc32"
	"sync"
	"time"
)

const (
	BucketNum = 256
)

type BigSyncMap struct {
	// 分段存储数据,提供并发访问效率 key string, val any
	bucket [BucketNum]*sync.Map
	// 存储每个key的过期时间 key string | val int64时间戳(秒)
	expireMap *sync.Map
}

var (
	mc   *BigSyncMap //memory_cache 缩写
	once sync.Once
)

// GetMCInstance 获取内存缓存实例
func GetMCInstance() *BigSyncMap {
	once.Do(func() {
		mc = &BigSyncMap{
			expireMap: &sync.Map{},
		}
		for i := 0; i < BucketNum; i++ {
			mc.bucket[i] = &sync.Map{}
		}
		// 开启协程 定时清除过期内容
		go mc.flushExpireData()
	})
	return mc
}
func (m *BigSyncMap) Get(key string) (any, bool) {
	return m.bucket[checksum(key)].Load(key)
}
func (m *BigSyncMap) Set(key string, val any, ttl int64) {
	m.bucket[checksum(key)].Store(key, val)
	if ttl != 0 {
		m.expireMap.Store(key, time.Now().Unix()+ttl)
	}
}
func (m *BigSyncMap) Del(key string) {
	m.bucket[checksum(key)].Delete(key)
	m.expireMap.Delete(key)
}

func (m *BigSyncMap) flushExpireData() {
	var tn int64
	for {
		tn = time.Now().Unix()
		m.expireMap.Range(func(key, expireTime any) bool {
			if expireTime.(int64) <= tn {
				m.Del(key.(string))
			}
			return true
		})
		time.Sleep(time.Minute * 2)
	}
}

// checksum 根据crc32算法取key的index
func checksum(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key)) % BucketNum)
}
