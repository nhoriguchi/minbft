// Copyright (c) 2018 NEC Laboratories Europe GmbH.
//
// Authors: Wenting Li <wenting.li@neclab.eu>
//          Sergey Fedorov <sergey.fedorov@neclab.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package minbft

import (
	"sync/atomic"

	"github.com/hyperledger-labs/minbft/api"
	"github.com/hyperledger-labs/minbft/core/internal/clientstate"
	"github.com/hyperledger-labs/minbft/messages"
)

// requestReplier provides Reply message given Request message.
//
// It returns a channel that can be used to receive a Reply message
// corresponding to the supplied Request message. It is safe to invoke
// concurrently.
type requestReplier func(request *messages.Request) <-chan *messages.Reply

// requestValidator validates a Request message.
//
// It authenticates and checks the supplied message for internal
// consistency. It does not use replica's current state and has no
// side-effect. It is safe to invoke concurrently.
type requestValidator func(request *messages.Request) error

// requestProcessor processes a valid Request message.
//
// It fully processes the supplied message in the context of the
// current replica's state. The supplied message is assumed to be
// authentic and internally consistent. The return value new indicates
// if the message has not been processed by this replica before. It is
// safe to invoke concurrently.
type requestProcessor func(request *messages.Request) (new bool, err error)

// requestExecutor given a Request message executes the requested
// operation, produces the corresponding Reply message ready for
// delivery to the client, and hands it over for further processing.
type requestExecutor func(request *messages.Request)

// operationExecutor executes an operation on the local instance of
// the replicated state machine. The result of operation execution
// will be send to the returned channel once it is ready. It is not
// allowed to invoke concurrently.
type operationExecutor func(operation []byte) (resultChan <-chan []byte)

// requestSeqCapturer synchronizes beginning of processing of request
// identifier in Request message.
//
// Processing of Request messages generated by the same client is
// synchronizes by waiting to ensure each request identifier is
// processed at most once and in the order of increasing identifier
// value. The return value new indicates if the message needs to be
// processed. The processing has to be completed by corresponding
// requestSeqReleaser invoked on the same message. It is safe to
// invoke concurrently.
type requestSeqCapturer func(request *messages.Request) (new bool)

// requestSeqReleaser synchronizes completion of processing of request
// identifier in Request message.
//
// It allows to continue processing of Request messages, synchronized
// by requestSeqCapturer. It is safe to invoke concurrently.
type requestSeqReleaser func(request *messages.Request)

// requestSeqPreparer records request identifier as prepared.
//
// It records the request identifier from the supplied message as
// prepared. It returns true if the request identifier from the client
// could not have been prepared before. The identifier must be
// previously captured. It is safe to invoke concurrently.
type requestSeqPreparer func(request *messages.Request) (new bool)

// requestSeqRetirer records request identifier as retired.
//
// It records the request identifier from the supplied message as
// retired. It returns true if the request identifier from the client
// could not have been retired before. The identifier must be
// previously prepared. It is safe to invoke concurrently.
type requestSeqRetirer func(request *messages.Request) (new bool)

// makeRequestValidator constructs an instance of requestValidator
// using the supplied abstractions.
func makeRequestValidator(verify messageSignatureVerifier) requestValidator {
	return func(request *messages.Request) error {
		return verify(request)
	}
}

// makeRequestProcessor constructs an instance of requestProcessor
// using id as the current replica ID, n as the total number of nodes,
// and the supplied abstractions.
func makeRequestProcessor(id, n uint32, view viewProvider, captureSeq requestSeqCapturer, releaseSeq requestSeqReleaser, prepareSeq requestSeqPreparer, handleGeneratedUIMessage generatedUIMessageHandler) requestProcessor {
	return func(request *messages.Request) (new bool, err error) {
		view := view()
		primary := isPrimary(view, id, n)

		if new = captureSeq(request); !new {
			return false, nil
		}
		defer releaseSeq(request)

		// TODO: A new request ID has arrived; the request
		// timer should be re-/started in backup replicas at
		// this point.

		if primary {
			prepare := &messages.Prepare{
				Msg: &messages.Prepare_M{
					View:      view,
					ReplicaId: id,
					Request:   request,
				},
			}

			handleGeneratedUIMessage(prepare)

			if new := prepareSeq(request); !new {
				panic("Duplicate Prepare generated")
			}
		}

		return true, nil
	}
}

// makeRequestReplier constructs an instance of requestReplier using
// the supplied client state provider.
func makeRequestReplier(provider clientstate.Provider) requestReplier {
	return func(request *messages.Request) <-chan *messages.Reply {
		state := provider(request.Msg.ClientId)
		return state.ReplyChannel(request.Msg.Seq)
	}
}

// makeRequestExecutor constructs an instance of requestExecutor using
// the supplied replica ID, operation executor, message signer, and
// reply consumer.
func makeRequestExecutor(id uint32, executor operationExecutor, signer replicaMessageSigner, handleGeneratedMessage generatedMessageHandler) requestExecutor {
	return func(request *messages.Request) {
		resultChan := executor(request.Msg.Payload)
		go func() {
			result := <-resultChan

			reply := &messages.Reply{
				Msg: &messages.Reply_M{
					ReplicaId: id,
					ClientId:  request.Msg.ClientId,
					Seq:       request.Msg.Seq,
					Result:    result,
				},
			}
			signer(reply)
			handleGeneratedMessage(reply)
		}()
	}
}

// makeOperationExecutor constructs an instance of operationExecutor
// using the supplied interface to external request consumer module.
func makeOperationExecutor(consumer api.RequestConsumer) operationExecutor {
	busy := uint32(0) // atomic flag to check for concurrent execution

	return func(op []byte) <-chan []byte {
		if wasBusy := atomic.SwapUint32(&busy, uint32(1)); wasBusy != uint32(0) {
			panic("Concurrent operation execution detected")
		}
		resultChan := consumer.Deliver(op)
		atomic.StoreUint32(&busy, uint32(0))

		return resultChan
	}
}

// makeRequestSeqCapturer constructs an instance of requestSeqCapturer
// using the supplied client state provider.
func makeRequestSeqCapturer(provideClientState clientstate.Provider) requestSeqCapturer {
	return func(request *messages.Request) (new bool) {
		clientID := request.Msg.ClientId
		seq := request.Msg.Seq

		return provideClientState(clientID).CaptureRequestSeq(seq)
	}
}

// makeRequestSeqReleaser constructs an instance of requestSeqReleaser
// using the supplied client state provider.
func makeRequestSeqReleaser(provideClientState clientstate.Provider) requestSeqReleaser {
	return func(request *messages.Request) {
		clientID := request.Msg.ClientId
		seq := request.Msg.Seq

		err := provideClientState(clientID).ReleaseRequestSeq(seq)
		if err != nil {
			panic(err)
		}
	}
}

// makeRequestSeqPreparer constructs an instance of requestSeqPreparer
// using the supplied interface.
func makeRequestSeqPreparer(provideClientState clientstate.Provider) requestSeqPreparer {
	return func(request *messages.Request) (new bool) {
		clientID := request.Msg.ClientId
		seq := request.Msg.Seq

		if new, err := provideClientState(clientID).PrepareRequestSeq(seq); err != nil {
			panic(err)
		} else if !new {
			return false
		}

		return true
	}
}

// makeRequestSeqRetirer constructs an instance of requestSeqRetirer
// using the supplied interface.
func makeRequestSeqRetirer(provideClientState clientstate.Provider) requestSeqRetirer {
	return func(request *messages.Request) (new bool) {
		clientID := request.Msg.ClientId
		seq := request.Msg.Seq

		if new, err := provideClientState(clientID).RetireRequestSeq(seq); err != nil {
			panic(err)
		} else if !new {
			return false
		}

		return true
	}
}
