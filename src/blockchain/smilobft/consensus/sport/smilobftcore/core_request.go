// Copyright 2019 The go-smilo Authors
// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package smilobftcore

import (
	"go-didux/src/blockchain/smilobft/consensus/sport"
)

func (c *core) handleRequest(request *sport.Request) error {
	logger := c.logger.New("state", c.state, "seq", c.current.sequence)

	if err := c.checkRequestMsg(request); err != nil {
		if err == errInvalidMessage {
			logger.Warn("invalid request")
			return err
		}
		logger.Warn("unexpected request", "err", err, "number", request.BlockProposal.Number(), "hash", request.BlockProposal.Hash())
		return err
	}

	logger.Trace("handleRequest", "number", request.BlockProposal.Number(), "hash", request.BlockProposal.Hash())

	c.current.pendingRequest = request
	if c.state == StateAcceptRequest {
		logger.Debug("handleRequest, StateAcceptRequest, sendPreprepare", "number", request.BlockProposal.Number(), "hash", request.BlockProposal.Hash())
		c.sendPreprepare(request)
	}
	return nil
}

// check request state
// return errInvalidMessage if the message is invalid
// return errFutureMessage if the sequence of proposal is larger than current sequence
// return errOldMessage if the sequence of proposal is smaller than current sequence
func (c *core) checkRequestMsg(request *sport.Request) error {
	if request == nil || request.BlockProposal == nil {
		return errInvalidMessage
	}

	if c := c.current.sequence.Cmp(request.BlockProposal.Number()); c > 0 {
		return errOldMessage
	} else if c < 0 {
		return errFutureMessage
	} else {
		return nil
	}
}

func (c *core) storeRequestMsg(request *sport.Request) {
	logger := c.logger.New("state", c.state)

	logger.Trace("Store future request", "number", request.BlockProposal.Number(), "hash", request.BlockProposal.Hash())

	c.pendingRequestsMu.Lock()
	defer c.pendingRequestsMu.Unlock()

	c.pendingRequests.Push(request, float32(-request.BlockProposal.Number().Int64()))
}

func (c *core) processPendingRequests() {
	c.pendingRequestsMu.Lock()
	defer c.pendingRequestsMu.Unlock()

	for !(c.pendingRequests.Empty()) {
		m, prio := c.pendingRequests.Pop()
		r, ok := m.(*sport.Request)
		if !ok {
			c.logger.Warn("Malformed request, skip", "msg", m)
			continue
		}
		// Push back if it's a future message
		err := c.checkRequestMsg(r)
		if err != nil {
			if err == errFutureMessage {
				c.logger.Trace("Stop processing request", "number", r.BlockProposal.Number(), "hash", r.BlockProposal.Hash())
				c.pendingRequests.Push(m, prio)
				break
			}
			c.logger.Trace("Skip the pending request", "number", r.BlockProposal.Number(), "hash", r.BlockProposal.Hash(), "err", err)
			continue
		}
		c.logger.Trace("Post pending request", "number", r.BlockProposal.Number(), "hash", r.BlockProposal.Hash())

		go c.sendEvent(sport.RequestEvent{
			BlockProposal: r.BlockProposal,
		})
	}
}
