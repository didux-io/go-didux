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
// TODO: config bootnodes
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
// TODO: config bootnodes
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// Sport are the enode URLs of the P2P bootstrap nodes running on the SPORT consensus mainnet.
var SportBootnodes = []string{
	// Didux Go Bootnodes
	"enode://95f1106daed2292d350a08815ddfa51f8896cd908e2b924ad626c6221133b209a1b44c8eb7559f2b9d11244a5e23d42e9d375b88be9b86b0074013db62840c63@212.32.245.80:21000",
	"enode://171de506cdae1dfb4a47b3c89093017f14a3df96fb2c33aa1d022ef677c7e4c48a9910b022322ff423d0d15a639233c5a54119748252bc49dad6be69c3b3ed2c@212.32.245.83:21000",
	"enode://f7fd400dd961bfa872fcc8f9c440f4efb6eeaa8a452a3fd4ace253beff8435b7ad4f0e8ad376b4a6830814511a67691525fd5a1f672b4a450140570d6682f476@212.32.245.88:21000",
	"enode://457090d01ba0043491cccf648bf8a5190c991ed315a0efe8fc258f52c4c9fabec1fd90a3fe974ef948ef7304f7bca449356f539bd91ab2167788e843dc63fe07@212.32.245.91:21000",
	"enode://754eb9fadc62f9aab15ce8f0a3ebc950607f659af2505ec4e7d4ea316b9301a4892d72370ff852a264c0ed63dea2c174d7b3cd461f3d0317bb58c7bef9d0b19f@51.89.103.232:21000",
	"enode://cf8dba408484d11f7942a9876761cfebe17d63f8ed28302e7a3c5d3d0513287ec4bd2dda14cb0b05c493c0ba9a827f825949020847cec1ed6c87ef70dded13b0@51.89.103.235:21000",
	"enode://85cfb66d4b7f7680dc125a49a6d5449ab41136e9ff0237bace3a2530609f673ca152339d1c289f5a757d84f572af6249305a6d3e8dd00cd59f665e65b1b94e5e@37.59.131.16:21000",
	"enode://4ecdf1f9f1a01e4ac8af4bb5831154ecf2b4127a825abcbcefa3e2bcce5a5705dc295d4f883555b46194f6b2cc42fed4dcb27d61af4ab520470513b2a2bce906@37.59.131.19:21000",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
// TODO: config bootnodes
var DiscoveryV5Bootnodes = []string{}
