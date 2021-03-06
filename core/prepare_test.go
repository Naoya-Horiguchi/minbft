// Copyright (c) 2018 NEC Laboratories Europe GmbH.
//
// Authors: Sergey Fedorov <sergey.fedorov@neclab.eu>
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
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"

	"github.com/hyperledger-labs/minbft/messages"
	"github.com/hyperledger-labs/minbft/usig"
)

func TestMakePrepareValidator(t *testing.T) {
	mock := new(testifymock.Mock)
	defer mock.AssertExpectations(t)

	n := randN()
	view := randView()
	primary := primaryID(n, view)
	backup := randOtherReplicaID(primary, n)

	request := &messages.Request{
		Msg: &messages.Request_M{
			ClientId: rand.Uint32(),
		},
	}
	ui := &usig.UI{Counter: rand.Uint64()}
	makePrepareMsg := func(id uint32) *messages.Prepare {
		return &messages.Prepare{
			Msg: &messages.Prepare_M{
				View:      view,
				ReplicaId: id,
				Request:   request,
			},
		}
	}

	verifyUI := func(msg messages.MessageWithUI) (*usig.UI, error) {
		args := mock.MethodCalled("uiVerifier", msg)
		return args.Get(0).(*usig.UI), args.Error(1)
	}
	validateRequest := func(request *messages.Request) error {
		args := mock.MethodCalled("requestValidator", request)
		return args.Error(0)
	}
	validate := makePrepareValidator(n, verifyUI, validateRequest)

	prepare := makePrepareMsg(backup)
	err := validate(prepare)
	assert.Error(t, err)

	prepare = makePrepareMsg(primary)

	mock.On("requestValidator", request).Return(fmt.Errorf("Invalid signature")).Once()
	err = validate(prepare)
	assert.Error(t, err)

	mock.On("requestValidator", request).Return(nil).Once()
	mock.On("uiVerifier", prepare).Return((*usig.UI)(nil), fmt.Errorf("UI not valid")).Once()
	err = validate(prepare)
	assert.Error(t, err)

	mock.On("requestValidator", request).Return(nil).Once()
	mock.On("uiVerifier", prepare).Return(ui, nil).Once()
	err = validate(prepare)
	assert.NoError(t, err)
}

func TestMakePrepareApplier(t *testing.T) {
	mock := new(testifymock.Mock)
	defer mock.AssertExpectations(t)

	n := randN()
	view := randView()
	primary := primaryID(n, view)
	id := randOtherReplicaID(primary, n)
	prepareRequestSeq := func(request *messages.Request) (new bool) {
		args := mock.MethodCalled("requestSeqPreparer", request)
		return args.Bool(0)
	}
	collectCommitment := func(id uint32, prepare *messages.Prepare) error {
		args := mock.MethodCalled("commitmentCollector", id, prepare)
		return args.Error(0)
	}
	handleGeneratedUIMessage := func(msg messages.MessageWithUI) {
		mock.MethodCalled("generatedUIMessageHandler", msg)
	}
	apply := makePrepareApplier(id, prepareRequestSeq, collectCommitment, handleGeneratedUIMessage)

	request := &messages.Request{
		Msg: &messages.Request_M{
			Seq: rand.Uint64(),
		},
	}
	ownPrepare := &messages.Prepare{
		Msg: &messages.Prepare_M{
			View:      viewForPrimary(n, id),
			ReplicaId: id,
			Request:   request,
		},
	}
	prepare := &messages.Prepare{
		Msg: &messages.Prepare_M{
			View:      view,
			ReplicaId: primary,
			Request:   request,
		},
	}
	commit := &messages.Commit{
		Msg: &messages.Commit_M{
			View:      view,
			ReplicaId: id,
			PrimaryId: primary,
			Request:   request,
		},
	}

	mock.On("requestSeqPreparer", request).Return(false).Once()
	err := apply(prepare)
	assert.Error(t, err, "Request ID already prepared")

	mock.On("requestSeqPreparer", request).Return(false).Once()
	err = apply(ownPrepare)
	assert.Error(t, err, "Request ID already prepared")

	mock.On("requestSeqPreparer", request).Return(true).Once()
	mock.On("commitmentCollector", id, ownPrepare).Return(fmt.Errorf("Error")).Once()
	err = apply(ownPrepare)
	assert.Error(t, err, "Failed to collect commitment")

	mock.On("requestSeqPreparer", request).Return(true).Once()
	mock.On("commitmentCollector", id, ownPrepare).Return(nil).Once()
	err = apply(ownPrepare)
	assert.NoError(t, err)

	mock.On("requestSeqPreparer", request).Return(true).Once()
	mock.On("commitmentCollector", primary, prepare).Return(fmt.Errorf("Error")).Once()
	err = apply(prepare)
	assert.Error(t, err, "Failed to collect commitment")

	mock.On("requestSeqPreparer", request).Return(true).Once()
	mock.On("commitmentCollector", primary, prepare).Return(nil).Once()
	mock.On("generatedUIMessageHandler", commit).Once()
	err = apply(prepare)
	assert.NoError(t, err)
}
