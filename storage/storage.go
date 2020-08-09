package storage

import (
	"github.com/olivetree123/htable"
)

const HashType int = 1
const QueueType int = 2
const SetType int = 3

var t *htable.HashTable

func init() {
	t = htable.NewHashTable()
}
