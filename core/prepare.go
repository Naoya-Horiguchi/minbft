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
	"time"

	"github.com/spf13/viper"
	"github.com/hyperledger-labs/minbft/messages"
)

// prepareValidator validates a Prepare message.
//
// It authenticates and checks the supplied message for internal
// consistency. It does not use replica's current state and has no
// side-effect. It is safe to invoke concurrently.
type prepareValidator func(prepare *messages.Prepare) error

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

// makePrepareApplier constructs an instance of prepareApplier using
// id as the current replica ID, and the supplied abstract interfaces.
func makePrepareApplier(id uint32, prepareSeq requestSeqPreparer, collectCommitment commitmentCollector, handleGeneratedUIMessage generatedUIMessageHandler, stopPrepTimer prepareTimerStopper, startReqTimer requestTimerStarter) prepareApplier {
	return func(prepare *messages.Prepare) error {
		request := prepare.Msg.Request
		if new := prepareSeq(request); !new {
			return fmt.Errorf("Request already prepared")
		}

		scenario := viper.GetInt("debug.scenario")
		if (scenario == 3) {
			time.Sleep(500*time.Millisecond)
			fmt.Printf("Prepare Delayed for 0.5 sec\n")
		} else if (scenario == 4) {
			fmt.Printf("Prepare Filtered\n")
			return nil
		}

		// TODO: need more better comment
		stopPrepTimer(prepare.Msg.ReplicaId)
		startReqTimer(request.Msg.ClientId, prepare.Msg.View)

		if (scenario == 5) {
			time.Sleep(500*time.Millisecond)
			fmt.Printf("Prepare Delayed for 0.5 sec started request timer\n")
		} else if (scenario == 6) {
			fmt.Printf("Prepare Filtered after started request timer\n")
			return nil
		}

		primaryID := prepare.ReplicaID()

		if err := collectCommitment(primaryID, prepare); err != nil {
			return fmt.Errorf("Prepare cannot be taken into account: %s", err)
		}

		if id == primaryID {
			return nil // primary does not generate Commit
		}

		commit := &messages.Commit{
			Msg: &messages.Commit_M{
				View:      prepare.Msg.View,
				ReplicaId: id,
				PrimaryId: primaryID,
				Request:   request,
				PrimaryUi: prepare.UIBytes(),
			},
		}

		handleGeneratedUIMessage(commit)

		return nil
	}
}
