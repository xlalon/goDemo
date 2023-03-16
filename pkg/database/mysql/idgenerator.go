package mysql

import (
	"sync"
	"time"
)

const (
	epoch             = int64(1577808000) // 2020-01-01 00:00:00
	timestampBits     = uint(31)          // 41 -> 31
	datacenteridBits  = uint(2)           // 2 -> 2
	workeridBits      = uint(3)           // 7 -> 3
	sequenceBits      = uint(11)          // 12 -> 11
	timestampMax      = int64(-1 ^ (-1 << timestampBits))
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))
	workeridMax       = int64(-1 ^ (-1 << workeridBits))
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))
	workeridShift     = sequenceBits
	datacenteridShift = sequenceBits + workeridBits
	timestampShift    = sequenceBits + workeridBits + datacenteridBits
)

var defaultIdGenerator = IdGenerator{}

// IdGenerator Snowflake
type IdGenerator struct {
	mu           sync.Mutex
	timestamp    int64
	workerid     int64
	datacenterid int64
	sequence     int64
}

func (ig *IdGenerator) nextID() int64 {
	ig.mu.Lock()
	now := time.Now().Unix()
	if ig.timestamp == now {
		ig.sequence = (ig.sequence + 1) & sequenceMask
		if ig.sequence == 0 {
			for now <= ig.timestamp {
				now = time.Now().Unix()
			}
		}
	} else {
		ig.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		ig.mu.Unlock()
		return 0
	}
	ig.timestamp = now
	r := (t)<<timestampShift | (ig.datacenterid << datacenteridShift) | (ig.workerid << workeridShift) | (ig.sequence)
	ig.mu.Unlock()
	return r
}

func NextID() int64 {
	return defaultIdGenerator.nextID()
}
