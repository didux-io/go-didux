// Copyright 2019 The go-smilo Authors
// Copyright 2016 The go-ethereum Authors
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

package blackbox

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/patrickmn/go-cache"
)

func (b *Blackbox) PostRaw(data []byte, from string, to []string) (out []byte, err error) {
	if b == nil || b.isBlackboxNotInUse {
		log.Error("Could not start Blackbox, Post, PostData, ", "b", b, "error", ErrBlackboxIsNotStarted)
		return nil, ErrBlackboxIsNotStarted
	}
	out, err = b.node.PostDataRaw(data, from, to)
	if err != nil {
		log.Error("Could not Post to Blackbox, Post, PostDataRaw, ", "error", err)
		return nil, err
	}
	return out, nil
}

func (b *Blackbox) PostRawTransaction(data []byte, to []string) (out []byte, err error) {
	if b == nil || b.isBlackboxNotInUse {
		log.Error("Could not start Blackbox, Post, PostData, ", "b", b, "error", ErrBlackboxIsNotStarted)
		return nil, ErrBlackboxIsNotStarted
	}
	out, err = b.node.PostDataRawTransaction(data, to)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *Blackbox) Get(data []byte) ([]byte, error) {
	if b == nil || b.isBlackboxNotInUse {
		log.Error("Could not start Blackbox, Get ", "error", ErrBlackboxIsNotStarted)
		return nil, ErrBlackboxIsNotStarted
	}
	if len(data) == 0 {
		return data, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	dataStr := string(data)
	x, found := b.cache.Get(dataStr)
	if found {
		return x.([]byte), nil
	}
	pl, _ := b.node.GetData(data)
	b.cache.Set(dataStr, pl, cache.DefaultExpiration)
	return pl, nil
}

func New(path string) (*Blackbox, error) {
	info, err := os.Lstat(path)
	if err != nil {
		log.Error("Could not start Blackbox, New, os.Lstat ", "path", path, "error", err)
		return nil, err
	}
	// We accept either the socket or a configuration file that points to
	// a socket.
	isSocket := info.Mode()&os.ModeSocket != 0
	if !isSocket {
		cfg, err := LoadConfig(path)
		if err != nil {
			log.Error("Could not start Blackbox, New, LoadConfig, ", "path", path, "error", err)
			return nil, err
		}
		path = filepath.Join(cfg.WorkDir, cfg.Socket)
	}
	err = RunNode(path)
	if err != nil {
		log.Error("Could not start Blackbox, New, RunNode, ", "path", path, "error", err)
		return nil, err
	}
	n, err := CreateClient(path)
	if err != nil {
		log.Error("Could not start Blackbox, New, CreateClient, ", "path", path, "error", err)
		return nil, err
	}
	return &Blackbox{
		node:               n,
		cache:              cache.New(1*time.Minute, 1*time.Minute),
		isBlackboxNotInUse: false,
	}, nil
}

func CreateNew(path string) *Blackbox {
	log.Debug("############################## Connecting to BlackBox, CreateNew, ", "path", path)
	if strings.EqualFold(path, "ignore") {
		return &Blackbox{
			node:               nil,
			cache:              nil,
			isBlackboxNotInUse: true,
		}
	}
	b, err := New(path)
	if err != nil || b == nil {
		log.Error("############################## ERROR: Failed to connect to BlackBox, CreateNew, ", "path", path, "error", err)
	}
	return b
}
