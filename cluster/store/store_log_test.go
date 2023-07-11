package store

import (
	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/lib/datastore"
	"github.com/hashicorp/raft"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	foo := func(conf config.Options) (raft.LogStore, error) {
		return &datastore.InMemStore{}, nil
	}
	conf := config.Options{}
	err := Register("memory", foo)
	assert.NoError(t, err)
	// get the datastore and check
	ds, err := getDataStore("memory", conf)
	assert.NoError(t, err)
	assert.IsType(t, &datastore.InMemStore{}, ds)
}
