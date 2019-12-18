// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
//TODO: config bootnodes
// Ethereum Foundation Go Bootnodes
var MainnetBootnodes = []string{}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	//TODO: config bootnodes
	"enode://07f7317ac02bbdf04024ce8b832395bb372e7869b34648d4d2adb2f72c09b18b6e710039a6f85d12fefbb25467020cc5e0f8a6bfe1013e2c06dce169cad74c72@51.89.103.232:21000",
	"enode://2765f23ee14b18fb7af06117edc5d73e35cb937d062724e7313ca10004bb0ab2f078335933aa907fb160d81f3f85ffdeba502488a6936f6d60da0b8a1937bedc@51.89.103.233:21000",
	"enode://394f1c72f7049a419d432dab3755a804c7858cf99a5f8c63d9cb4b13c7d0c012450c6c8e3f241b2d2926179b74c24e3f67afeb6c52bbd1d586a9838a69062635@51.89.103.234:21000",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
//TODO: config bootnodes
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// Sport are the enode URLs of the P2P bootstrap nodes running on the SPORT consensus mainnet.
var SportBootnodes = []string{
	// Smilo Foundation Go Bootnodes
	"enode://95f1106daed2292d350a08815ddfa51f8896cd908e2b924ad626c6221133b209a1b44c8eb7559f2b9d11244a5e23d42e9d375b88be9b86b0074013db62840c63@212.32.245.80:21000",
	"enode://f7fd400dd961bfa872fcc8f9c440f4efb6eeaa8a452a3fd4ace253beff8435b7ad4f0e8ad376b4a6830814511a67691525fd5a1f672b4a450140570d6682f476@212.32.245.88:21000",
	"enode://754eb9fadc62f9aab15ce8f0a3ebc950607f659af2505ec4e7d4ea316b9301a4892d72370ff852a264c0ed63dea2c174d7b3cd461f3d0317bb58c7bef9d0b19f@51.89.103.232:21000",
	"enode://85cfb66d4b7f7680dc125a49a6d5449ab41136e9ff0237bace3a2530609f673ca152339d1c289f5a757d84f572af6249305a6d3e8dd00cd59f665e65b1b94e5e@37.59.131.16:21000",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
//TODO: config bootnodes
var DiscoveryV5Bootnodes = []string{}
