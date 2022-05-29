/*
Package params for generating query parameters.

	b := params.NewMessageSendBuilder()
	b.PeerID(123)
	b.Random(0)
	b.DontParseLinks(false)
	b.Message("Test message")

	res, err = api.MessageSend(b.Params)
*/
package params // import "github.com/derad6709org/vksdk/v2/api/params"
