/*
 * Copyright (c) 2017-2020 The qitmeer developers
 */

package consensus

import (
	"github.com/Qitmeer/qng-core/common/hash"
)

type ChainVM interface {
	VM

	GetBlock(*hash.Hash) (Block, error)

	BuildBlock([]Tx) (Block, error)

	ParseBlock([]byte) (Block, error)

	LastAccepted() (*hash.Hash, error)

	GetBalance(string) (int64, error)

	VerifyTx(tx Tx) (int64, error)

	ConnectBlock(block Block) error

	DisconnectBlock(block Block) error
}
