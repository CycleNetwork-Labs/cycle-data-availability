package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/0xPolygon/cdk-data-availability/client"
	"github.com/0xPolygon/cdk-data-availability/log"
	jTypes "github.com/0xPolygon/cdk-data-availability/rpc"
	"github.com/0xPolygon/cdk-data-availability/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"strings"
)

const url = "http://localhost:8454"

const (
	keystorePath            = "sequencer.keystore"
	password                = "45RricuRBV87@6534qxiWBNO"
	trustedSequencerAddress = "0xc4b1c6748ad745eafae93e3e41d30d610c15f3a5"

	newTrustedSequencerAddress    = "0xd2D1A0aF207675D571b37c8989c26Fedb08aEF59"
	newTrustedSequencerPrivateKey = "762e5f69ea54dc60d1a3585c4ecff591d31d8d56525470e7ee2cc1d7f74ec691"
)

var daClient *client.Client
var key *ecdsa.PrivateKey

func main() {
	var err error

	// sign from trusted sequencer
	err = testSignSequence(5, 1, key)
	if err != nil {
		log.Fatal("Failed to sign sequence ", err)
	} else {
		log.Infof("Successfully signed sequence")
	}

	// sign from non-trusted sequencer
	pkStr := "b1571a030bbe12187308e2672e0819a5146bc62e307501d71bddb08bb15514ea"
	pk, err := getPrivateKeyFromString(pkStr)
	if err != nil {
		log.Fatal(err)
	}

	err = testSignSequence(5, 1, pk)
	if err == nil {
		log.Fatal("expect sign sequence failed but returns ok")
	} else {
		log.Infof("sign sequence failed and returns %v", err)
	}

	// test GetOffChainData
	err = testGetOffChainData(5, 1, key)
	if err == nil {
		log.Infof("GetOffChainData succeeded")
	} else {
		log.Fatalf("testGetOffChainData failed %v", err)
	}

	// test storing with overlapping
	err = testGetOffChainDataOverlap(key)
	if err == nil {
		log.Infof("GetOffChainDataOverlap succeeded")
	} else {
		log.Fatal("testGetOffChainDataOverlap failed ", err)
	}

	// test trusted sequencer address event
	newKey, err := getPrivateKeyFromString(newTrustedSequencerPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = testTrustedSequenceChange(key, common.HexToAddress(newTrustedSequencerAddress), newKey)
	if err == nil {
		log.Infof("TestTrustedSequenceChange succeeded")
	} else {
		log.Fatal("testTrustedSequenceChange failed ", err)
	}
}

func init() {
	var err error
	key, err = NewKeyFromKeystore(keystorePath, password)
	if err != nil {
		log.Fatal(err)
	}

	daClient = client.New(url)

	err = initContract(l1Url, zkevmContractAddress, chainId, zkevmAdminPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
}

func testSignSequence(batchNum int, startBatch int, priKey *ecdsa.PrivateKey) error {
	sequence := generateSequences(batchNum, startBatch, trustedSequencerAddress)

	signedSequence, err := sequence.Sign(priKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = daClient.SignSequence(*signedSequence)
	if err != nil {
		return err
	}
	return nil
}

func testGetOffChainData(batchNum int, startBatch int, priKey *ecdsa.PrivateKey) error {
	sequence := generateSequences(batchNum, startBatch, trustedSequencerAddress)

	signedSequence, err := sequence.Sign(priKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = daClient.SignSequence(*signedSequence)
	if err != nil {
		log.Fatal(err)
	}

	return checkOffChainDataWithSequence(*sequence)
}

// test storing data with overlapping batches
func testGetOffChainDataOverlap(priKey *ecdsa.PrivateKey) error {
	batchNum := 5
	startBatch := 1

	sequence := generateSequences(batchNum, startBatch, trustedSequencerAddress)
	signedSequence, err := sequence.Sign(priKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = daClient.SignSequence(*signedSequence)
	if err != nil {
		log.Fatal(err)
	}

	// generate another sequences and append to previous sequence
	sequence2 := generateSequences(batchNum, startBatch+batchNum, trustedSequencerAddress)

	for _, batch := range sequence2.Batches {
		sequence.Batches = append(sequence.Batches, batch)
	}

	signedSequence2, err := sequence.Sign(priKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = daClient.SignSequence(*signedSequence2)
	if err != nil {
		log.Fatal(err)
	}

	// check all data are stored
	return checkOffChainDataWithSequence(*sequence)
}

func checkOffChainDataWithSequence(sequence types.Sequence) error {
	for _, batch := range sequence.Batches {
		hash := crypto.Keccak256Hash(batch.L2Data)

		data, err := daClient.GetOffChainData(context.Background(), hash)
		if err != nil {
			return err
		}

		if !bytes.Equal(batch.L2Data, data) {
			return fmt.Errorf("l2 data not match, data %v, expected %v", data, batch.L2Data)
		}
	}
	log.Infof("checkOffChainDataWithSequence success, number of batches %d", len(sequence.Batches))
	return nil
}

func generateSequences(batchNum int, startBatch int, addrStr string) *types.Sequence {
	var sequence types.Sequence

	coinbase := common.HexToAddress(addrStr)
	for i := 0; i < batchNum; i++ {
		batch := types.Batch{
			Number:         jTypes.ArgUint64(startBatch + i),
			GlobalExitRoot: common.Hash{},
			Timestamp:      0,
			Coinbase:       coinbase,
			L2Data:         randData(1000),
		}

		sequence.Batches = append(sequence.Batches, batch)
	}
	sequence.OldAccInputHash = randHash()
	return &sequence
}

func randData(dataLen int) []byte {
	data := make([]byte, dataLen)
	rand.Read(data)
	return data
}

func randHash() common.Hash {
	var hash common.Hash
	rand.Read(hash[:])
	return hash
}

func NewKeyFromKeystore(path string, password string) (*ecdsa.PrivateKey, error) {
	keystoreEncrypted, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keystoreEncrypted, password)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}

func getPrivateKeyFromString(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyStr, "0x"))
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
