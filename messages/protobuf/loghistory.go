// Copyright (c) 2019 NEC Laboratories Europe GmbH.
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

package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type loghistory struct {
	Loghistory
}

func newLogHistory() *loghistory {
	return &loghistory{}
}

func (m *loghistory) init(r, peerID uint32, seq uint64, logs []byte, hash []byte) {
	m.Loghistory = Loghistory{Msg: &Loghistory_M{
		ReplicaId: r,
		PeerId: peerID,
		Sequence: seq,
		Logs: logs,
		Hash: hash,
	}}
}

func (m *loghistory) set(pbMsg *Loghistory) {
	m.Loghistory = *pbMsg
}

func (m *loghistory) MarshalBinary() ([]byte, error) {
	return proto.Marshal(&Message{Type: &Message_Loghistory{Loghistory: &m.Loghistory}})
}

func (m *loghistory) ReplicaID() uint32 {
	return m.Msg.GetReplicaId()
}

func (m *loghistory) PeerID() uint32 {
	return m.Msg.GetPeerId()
}

func (m *loghistory) PrevHash() []byte {
	return []byte{}
}

func (m *loghistory) Sequence() uint64 {
	return m.Msg.GetSequence()
}

func (m *loghistory) Authenticator() []byte {
	return []byte{}
}

func (m *loghistory) ExtractMessage() []byte {
	return []byte{}
}

func (m *loghistory) Logs() []byte {
	return m.Msg.GetLogs()
}

func (m *loghistory) Hash() []byte {
	return m.Msg.GetHash()
}

func (loghistory) ImplementsReplicaMessage() {}
func (loghistory) ImplementsLogHistory() {}