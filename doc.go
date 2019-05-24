// Package multicast offers a Chan type that can multicast and replay the
// messages you send to it to multiple receivers.
//
// The standard Go channel cannot multicast the same message to multiple receivers
// and it cannot play back messages previously sent to it. The Chan type offered
// here allows multicasting and playback of messages. You can even limit playback
// to messages younger than a certain age because `Chan` also stores a timestamp\
// with each message send.
//
// This multicast channel is different from other multicast implementations in
// that it uses only fast synchronization primitives like atomic operations to
// implement its features. Furthermore, it also doesn't use goroutines internally.
// This allows it to exhibit a very high performance.
//
// If you are looking for a situation where you need to record and replay a stream
// of data or need to split a stream of data into multiple identical streams,
// then this package offers a fast and simple implementation.
//
//	HOW THIS PACKAGE IS GENERATED
//
//	The package consists of three files; doc.go, example_test.go and channel.go.
//	The actual package code is in channel.go, but it is generated from a template
//	library using information from doc.go and example_test.go.
//
// 	The [doc.go] file imports the generics library we want to use:
//
//		import _ "github.com/reactivego/channel/generic"
//
// 	The [example_test.go] file contains an example program of all the
// 	functionality that is needed from the generics library. See the Example
//	below this text for the actual source of example_test.go.
//
// 	The package is generated by running the jig command. The code needed to
//	compile the example is generated into the file [channel.go]. That file then
//	contains the code for the actual channel package.
//
//	$ go get github.com/reactivego/multicast/generic
//	$ cd $GOPATH/src/github.com/reactivego/multicast
//	$ go run -v github.com/reactivego/jig -rv
//	removing file "multicast.go"
//	found 16 templates in package "channel" (github.com/reactivego/multicast/generic)
//	generating "NewChan"
//	  ChanPadding
//	  ChanState
//	  Chan
//	  ErrOutOfEndpoints
//	  endpoints
//	  NewChan
//	generating "Endpoint"
//	  Endpoint
//	generating "Chan FastSend"
//	  Chan slideBuffer
//	  Chan FastSend
//	generating "Chan Send"
//	  Chan Send
//	generating "Chan Close"
//	  Chan Close
//	generating "Chan Closed"
//	  Chan Closed
//	generating "Chan NewEndpoint"
//	  Chan NewEndpoint
//	generating "Endpoint Range"
//	  Endpoint Range
//	generating "Endpoint Cancel"
//	  Endpoint Cancel
//	generating "Endpoint commitData"
//	  Chan commitData
//	  missing "Endpoint commitData"
//	writing file "multicast.go"
package multicast

import _ "github.com/reactivego/multicast/generic"
