package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/0xPolygon/cdk-data-availability/etherman/smartcontracts/polygonzkevm"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

const (
	l1Url                = "http://127.0.0.1:8555"
	zkevmContractAddress = "0x9A676e781A523b5d0C0e43731313A708CB607508"
	chainId              = 1337
	zkevmAdminPrivateKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

var ethClient *ethclient.Client
var zkevm *polygonzkevm.Polygonzkevm
var auth *bind.TransactOpts

func initContract(url, contractAddr string, chainId int, priKey string) error {
	var err error

	ethClient, err = ethclient.Dial(url)
	if err != nil {
		return err
	}
	// Create smc client
	zkevmScAddr := common.HexToAddress(contractAddr)

	zkevm, err = polygonzkevm.NewPolygonzkevm(zkevmScAddr, ethClient)
	if err != nil {
		return err
	}

	auth, err = getAuth(priKey, uint64(chainId))
	if err != nil {
		return err
	}
	return nil
}

func getAuth(privateKeyStr string, chainID uint64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyStr, "0x"))
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(0).SetUint64(chainID))
}

func testTrustedSequenceChange(key *ecdsa.PrivateKey, newTrustedSequencer common.Address, newKey *ecdsa.PrivateKey) error {
	// get current trusted sequencer
	seqAddr, err := zkevm.TrustedSequencer(&bind.CallOpts{Pending: false})
	if err != nil {
		return err
	}

	log.Infof("trusted sequencer address %v", seqAddr.String())

	if seqAddr != common.HexToAddress(trustedSequencerAddress) {
		return fmt.Errorf("trusted sequencer address is not trusted sequencer, %s, %s", seqAddr.String(), trustedSequencerAddress)
	}

	// sign with the trusted sequencer key
	err = testSignSequence(5, 1, key)
	if err != nil {
		return err
	}

	// update trusted sequencer
	tx, err := zkevm.SetTrustedSequencer(auth, newTrustedSequencer)
	if err != nil {
		return err
	}

	err = waitTxToBeMined(context.Background(), ethClient, tx, 30*time.Second)
	if err != nil {
		return err
	}

	seqAddr, err = zkevm.TrustedSequencer(&bind.CallOpts{Pending: false})
	if err != nil {
		return err
	}

	log.Infof("trusted sequencer address %v", seqAddr.String())

	// sign with the same key failed
	err = testSignSequence(5, 1, key)
	if err == nil {
		return err
	}

	log.Infof("testSignSequence with previous trusted sequecer failed as expected, err %v", err)

	// sign with the updated trusted sequencer key ok
	err = testSignSequence(5, 1, newKey)
	if err != nil {
		return err
	}

	// restore trusted sequencer
	tx, err = zkevm.SetTrustedSequencer(auth, common.HexToAddress(trustedSequencerAddress))
	if err != nil {
		return err
	}

	err = waitTxToBeMined(context.Background(), ethClient, tx, 30*time.Second)
	if err != nil {
		return err
	}

	seqAddr, err = zkevm.TrustedSequencer(&bind.CallOpts{Pending: false})
	if err != nil {
		return err
	}

	log.Infof("trusted sequencer address %v", seqAddr.String())
	return nil
}

// WaitTxToBeMined waits until a tx has been mined or the given timeout expires.
func waitTxToBeMined(parentCtx context.Context, client *ethclient.Client, tx *types.Transaction, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if errors.Is(err, context.DeadlineExceeded) {
		return err
	} else if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("transaction has failed")
	}
	return nil
}
