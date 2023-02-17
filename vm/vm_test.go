// Copyright (C) 2022-2023, Sama , Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"reflect"
	"testing"

	"github.com/SamaNetwork/SamaVM/chain"
	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
)

func TestBlockCache(t *testing.T) {
	// create a block with "Unknown" status
	blk := &chain.StatelessBlock{
		StatefulBlock: &chain.StatefulBlock{
			Prnt:  ids.GenerateTestID(),
			Hght:  10000,
			Price: 1000,
			Cost:  100,
		},
	}
	blkID := blk.ID()

	vm := VM{
		blocks:         &cache.LRU[ids.ID, *chain.StatelessBlock]{Size: 3},
		verifiedBlocks: make(map[ids.ID]*chain.StatelessBlock),
	}

	// put the block into the cache "vm.blocks"
	// and delete from "vm.verifiedBlocks"
	vm.Accepted(blk)

	// we have not set up any persistent db
	// so this must succeed from using cache
	blk2, err := vm.GetStatelessBlock(blkID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(blk, blk2) {
		t.Fatalf("block expected %+v, got %+v", blk, blk2)
	}
}
