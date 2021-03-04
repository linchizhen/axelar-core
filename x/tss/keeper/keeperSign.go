package keeper

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"

	broadcast "github.com/axelarnetwork/axelar-core/x/broadcast/exported"
	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	"github.com/axelarnetwork/axelar-core/x/tss/exported"
	"github.com/axelarnetwork/axelar-core/x/tss/types"
	voting "github.com/axelarnetwork/axelar-core/x/vote/exported"
)

// StartSign starts a tss signing protocol using the specified key for the given chain.
func (k Keeper) StartSign(ctx sdk.Context, keyID string, sigID string, msg []byte) error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeSign,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueStart)))

	if _, ok := k.signStreams[sigID]; ok {
		return fmt.Errorf("signing protocol for ID %s already in progress", sigID)
	}

	counter, ok := k.GetSnapshotCounterForKeyID(ctx, keyID)
	if !ok {
		return fmt.Errorf("no snapshot counter for key ID %s registered", keyID)
	}
	snapshot, ok := k.snapshotter.GetSnapshotActiveValidators(ctx, counter)
	if !ok {
		return fmt.Errorf("no snapshot found for counter num %d", counter)
	}
	allValidatorsWithShares, ok := k.snapshotter.GetSnapshot(ctx, counter)
	// for now we recalculate the threshold
	// might make sense to store it with the snapshot after keygen is done.
	threshold := k.ComputeCorruptionThreshold(ctx, len(allValidatorsWithShares.Validators))
	activeValidators := snapshot.Validators

	k.Logger(ctx).Info(fmt.Sprintf("starting sign with threshold [%d] (need [%d]), online validators count [%d]",
		threshold, threshold+1, len(activeValidators)))

	if len(activeValidators) <= threshold {
		return fmt.Errorf(fmt.Sprintf("not enough active validators are online: threshold [%d], online [%d]",
			threshold, len(activeValidators)))
	}

	poll := voting.PollMeta{Module: types.ModuleName, Type: "sign", ID: sigID}
	if err := k.voter.InitPoll(ctx, poll); err != nil {
		return err
	}

	k.Logger(ctx).Info(fmt.Sprintf("new Sign: sig_id [%s] key_id [%s] message [%s]", sigID, keyID, string(msg)))

	// BEGIN: validity check

	// sign cannot proceed unless all validators have registered broadcast proxies
	if err := k.checkProxies(ctx, activeValidators); err != nil {
		return err
	}

	/*
		END: validity check -- any error below this point is local to the specific validator,
		so do not return an error but simply close the result channel
	*/

	if _, ok := k.getKeyIDForSig(ctx, sigID); ok {
		return fmt.Errorf("sigID %s has been used before", sigID)
	}
	k.setKeyIDForSig(ctx, sigID, keyID)

	voteChan := make(chan voting.MsgVote)

	stream, signInit := k.prepareSign(ctx, keyID, sigID, msg, activeValidators)
	if stream == nil || signInit == nil {
		close(voteChan)
		return nil // don't propagate nondeterministic errors
	}
	k.signStreams[sigID] = stream

	go func() {
		if err := stream.Send(&tofnd.MessageIn{Data: signInit}); err != nil {
			k.Logger(ctx).Error(sdkerrors.Wrap(err, "failed tofnd gRPC sign send sign init data").Error())
		}
	}()

	broadcastChan, resChan := k.handleStream(ctx, stream)

	// handle intermediate messages
	go func() {
		for msg := range broadcastChan {
			k.Logger(ctx).Debug(fmt.Sprintf(
				"handler goroutine: outgoing msg: session id [%.20s] broadcast? [%t] to [%.20s]",
				sigID, msg.IsBroadcast, msg.ToPartyUid))
			// sender is set by broadcaster
			tssMsg := &types.MsgSignTraffic{SessionID: sigID, Payload: msg}
			if err := k.broadcaster.Broadcast(ctx, []broadcast.MsgWithSenderSetter{tssMsg}); err != nil {
				k.Logger(ctx).Error(sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing sign msg").Error())
				return
			}
		}
	}()

	// handle result
	go func() {
		defer close(voteChan)
		sig, ok := <-resChan
		k.Logger(ctx).Info("handler goroutine: received sig from server!")
		if ok {
			k.voter.RecordVote(&types.MsgVoteSig{PollMeta: poll, SigBytes: sig})
		}
	}()
	return nil
}

// SignMsg takes a types.MsgSignTraffic from the chain and relays it to the keygen protocol
func (k Keeper) SignMsg(ctx sdk.Context, msg types.MsgSignTraffic) error {
	msgIn, err := k.prepareTrafficIn(ctx, msg.Sender, msg.SessionID, msg.Payload)
	if err != nil {
		return err
	}
	if msgIn == nil {
		return nil // don't propagate nondeterministic errors
	}

	stream, ok := k.signStreams[msg.SessionID]
	if !ok {
		k.Logger(ctx).Error(fmt.Sprintf("no signature session with id %s", msg.SessionID))
		return nil // don't propagate nondeterministic errors
	}

	if err := stream.Send(msgIn); err != nil {
		k.Logger(ctx).Error(sdkerrors.Wrap(err, "failure to send incoming msg to gRPC server").Error())
		return nil // don't propagate nondeterministic errors
	}
	return nil
}

// GetSig returns the signature associated with sigID
// or nil, nil if no such signature exists
func (k Keeper) GetSig(ctx sdk.Context, sigID string) (exported.Signature, bool) {
	bz := ctx.KVStore(k.storeKey).Get([]byte(sigPrefix + sigID))
	if bz == nil {
		return exported.Signature{}, false
	}
	btcecSig, err := btcec.ParseDERSignature(bz, btcec.S256())
	if err != nil {
		// the setter is controlled by the keeper alone, so an error here should be a catastrophic failure
		panic(err)
	}

	return exported.Signature{R: btcecSig.R, S: btcecSig.S}, true
}

// SetSig stores the given signature by its ID
func (k Keeper) SetSig(ctx sdk.Context, sigID string, signature []byte) {
	// Delete the reference to the signing stream with sigID because entering this function means the tss protocol has completed
	delete(k.signStreams, sigID)

	ctx.KVStore(k.storeKey).Set([]byte(sigPrefix+sigID), signature)
}

func (k Keeper) prepareSign(ctx sdk.Context, keyID, sigID string, msg []byte, validators []snapshot.Validator) (types.Stream, *tofnd.MessageIn_SignInit) {
	// TODO call GetLocalPrincipal only once at launch? need to wait until someone pushes a RegisterProxy message on chain...
	myAddress := k.broadcaster.GetLocalPrincipal(ctx)
	if myAddress.Empty() {
		k.Logger(ctx).Info("ignore Sign: my validator address is empty so I must not be a validator")
		return nil, nil
	}

	partyUids, _, err := addrToUids(validators, myAddress)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, nil
	}

	grpcCtx, _ := k.newGrpcContext()
	stream, err := k.client.Sign(grpcCtx)
	if err != nil {
		k.Logger(ctx).Error(sdkerrors.Wrap(err, "failed tofnd gRPC call Sign").Error())
		return nil, nil
	}
	k.signStreams[sigID] = stream
	// TODO refactor
	signInit := &tofnd.MessageIn_SignInit{
		SignInit: &tofnd.SignInit{
			NewSigUid:     sigID,
			KeyUid:        keyID,
			PartyUids:     partyUids,
			MessageToSign: msg,
		},
	}

	return stream, signInit
}

// GetKeyForSigID returns the key that produced the signature corresponding to the given ID
func (k Keeper) GetKeyForSigID(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool) {
	keyID, ok := k.getKeyIDForSig(ctx, sigID)
	if !ok {
		return ecdsa.PublicKey{}, false
	}
	return k.GetKey(ctx, keyID)
}

func (k Keeper) setKeyIDForSig(ctx sdk.Context, sigID string, keyID string) {
	ctx.KVStore(k.storeKey).Set([]byte(keyIDForSigPrefix+sigID), []byte(keyID))
}

func (k Keeper) getKeyIDForSig(ctx sdk.Context, sigID string) (string, bool) {
	bz := ctx.KVStore(k.storeKey).Get([]byte(keyIDForSigPrefix + sigID))
	if bz == nil {
		return "", false
	}
	return string(bz), true
}
