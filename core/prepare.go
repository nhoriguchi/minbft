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
	"fmt"

	"github.com/hyperledger-labs/minbft/messages"
)

// prepareValidator validates a Prepare message.
//
// It authenticates and checks the supplied message for internal
// consistency. It does not use replica's current state and has no
// side-effect. It is safe to invoke concurrently.
type prepareValidator func(prepare *messages.Prepare) error

// prepareProcessor processes a valid Prepare message.
//
// It fully processes the supplied message. The supplied message is
// assumed to be authentic and internally consistent. The return value
// new indicates if the message has not been processed by this replica
// before. It is safe to invoke concurrently.
type prepareProcessor func(prepare *messages.Prepare) (new bool, err error)

// prepareApplier applies Prepare message to current replica state.
//
// The supplied message is applied to the current replica state by
// changing the state accordingly and producing any required messages
// or side effects. The supplied message is assumed to be authentic
// and internally consistent. It is safe to invoke concurrently.
type prepareApplier func(prepare *messages.Prepare) error

// makePrepareValidator constructs an instance of prepareValidator
// using n as the total number of nodes, and the supplied abstract
// interfaces.
func makePrepareValidator(n uint32, verifyUI uiVerifier, validateRequest requestValidator) prepareValidator {
	return func(prepare *messages.Prepare) error {
		replicaID := prepare.Msg.ReplicaId
		view := prepare.Msg.View

		if !isPrimary(view, replicaID, n) {
			return fmt.Errorf("Prepare from backup %d for view %d", replicaID, view)
		}

		if err := validateRequest(prepare.Msg.Request); err != nil {
			return fmt.Errorf("Request invalid: %s", err)
		}

		if _, err := verifyUI(prepare); err != nil {
			return fmt.Errorf("UI not valid: %s", err)
		}

		return nil
	}
}

// makePrepareProcessor constructs an instance of prepareProcessor
// using id as the current replica ID, and the supplied abstract
// interfaces.
func makePrepareProcessor(id uint32, processRequest requestProcessor, captureUI uiCapturer, view viewProvider, applyPrepare prepareApplier) prepareProcessor {
	return func(prepare *messages.Prepare) (new bool, err error) {
		if prepare.Msg.ReplicaId == id {
			return false, nil
		}

		if _, err := processRequest(prepare.Msg.Request); err != nil {
			return false, fmt.Errorf("Failed to process request: %s", err)
		}

		new, releaseUI := captureUI(prepare)
		if !new {
			return false, nil
		}
		defer releaseUI()

		currentView := view()

		if prepare.Msg.View != currentView {
			return false, fmt.Errorf("Prepare is for view %d, current view is %d",
				prepare.Msg.View, currentView)
		}

		if err := applyPrepare(prepare); err != nil {
			return false, fmt.Errorf("Failed to apply Prepare: %s", err)
		}

		return true, nil
	}
}

// makePrepareApplier constructs an instance of prepareApplier using
// id as the current replica ID, and the supplied abstract interfaces.
func makePrepareApplier(id uint32, prepareSeq requestSeqPreparer, handleGeneratedUIMessage generatedUIMessageHandler, applyCommit commitApplier) prepareApplier {
	return func(prepare *messages.Prepare) error {
		if new := prepareSeq(prepare.Msg.Request); !new {
			return fmt.Errorf("Request already prepared")
		}

		commit := &messages.Commit{
			Msg: &messages.Commit_M{
				View:      prepare.Msg.View,
				ReplicaId: id,
				PrimaryId: prepare.ReplicaID(),
				Request:   prepare.Msg.Request,
				PrimaryUi: prepare.UIBytes(),
			},
		}

		if err := applyCommit(commit); err != nil {
			panic(fmt.Errorf("Failed to apply generated Commit: %s", err))
		}

		handleGeneratedUIMessage(commit)

		return nil
	}
}
