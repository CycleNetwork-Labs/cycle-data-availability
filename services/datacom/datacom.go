package datacom

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"

	"github.com/0xPolygon/cdk-data-availability/rpc"
	"github.com/0xPolygon/cdk-data-availability/sequencer"
	"github.com/0xPolygon/cdk-data-availability/types"
	"github.com/jackc/pgx/v4"
)

// APIDATACOM is the namespace of the datacom service
const APIDATACOM = "datacom"

// DataComEndpoints contains implementations for the "datacom" RPC endpoints
type DataComEndpoints struct {
	db               DBInterface
	txMan            rpc.DBTxManager
	privateKey       *ecdsa.PrivateKey
	sequencerTracker *sequencer.SequencerTracker
}

// NewDataComEndpoints returns DataComEndpoints
func NewDataComEndpoints(
	db DBInterface, privateKey *ecdsa.PrivateKey, sequencerTracker *sequencer.SequencerTracker,
) *DataComEndpoints {
	return &DataComEndpoints{
		db:               db,
		privateKey:       privateKey,
		sequencerTracker: sequencerTracker,
	}
}

// SignSequence generates the accumulated input hash aka accInputHash of the sequence and sign it.
// After storing the data that will be sent hashed to the contract, it returns the signature.
// This endpoint is only accessible to the sequencer
func (d *DataComEndpoints) SignSequence(signedSequence types.SignedSequence) (interface{}, rpc.Error) {
	// Verify that the request comes from the sequencer
	sender, err := signedSequence.Signer()
	if err != nil {
		return "0x0", rpc.NewRPCError(rpc.DefaultErrorCode, "failed to verify sender")
	}
	if sender != d.sequencerTracker.GetAddr() {
		return "0x0", rpc.NewRPCError(rpc.DefaultErrorCode, "unauthorized")
	}
	// Store off-chain data by hash (hash(L2Data): L2Data)
	_, err = d.txMan.NewDbTxScope(d.db, func(ctx context.Context, dbTx pgx.Tx) (interface{}, rpc.Error) {
		err := d.db.StoreOffChainData(ctx, signedSequence.Sequence.OffChainData(), dbTx)
		if err != nil {
			return "0x0", rpc.NewRPCError(rpc.DefaultErrorCode, "failed to store offchain data")
		}

		return nil, nil
	})
	if err != nil {
		return "0x0", rpc.NewRPCError(rpc.DefaultErrorCode, "failed to store offchain data")
	}
	// Sign
	signedSequenceByMe, err := d.signSequence(&signedSequence.Sequence)
	if err != nil {
		return "0x0", rpc.NewRPCError(rpc.DefaultErrorCode, "failed to sign")
	}
	// Return signature
	return signedSequenceByMe.Signature, nil
}

func (d *DataComEndpoints) signSequence(sequence *types.Sequence) (*types.SignedSequence, error) {
	signedSequence, err := sequence.Sign(d.privateKey)
	if err != nil {
		return nil, err
	}
	return signedSequence, nil
}

func formatSig(sig []byte) ([]byte, error) {
	if len(sig) != 65 {
		return nil, errors.New("invalid signature length")
	}

	rBytes := sig[:32]
	sBytes := sig[32:64]
	vByte := sig[64]

	if strings.ToUpper(common.Bytes2Hex(sBytes)) > "7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0" {
		magicNumber := common.Hex2Bytes("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141")
		sBig := big.NewInt(0).SetBytes(sBytes)
		magicBig := big.NewInt(0).SetBytes(magicNumber)
		s1 := magicBig.Sub(magicBig, sBig)
		sBytes = s1.Bytes()
		if vByte == 0 {
			vByte = 1
		} else {
			vByte = 0
		}
	}
	vByte += 27

	actualSignature := []byte{}
	actualSignature = append(actualSignature, rBytes...)
	actualSignature = append(actualSignature, sBytes...)
	actualSignature = append(actualSignature, vByte)

	return actualSignature, nil
}
