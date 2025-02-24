package caps

/*
	WARNING: this file is autogenerated by `make capdefs`
	DO NOT EDIT MANUALLY.
*/

const (
	// number of recognized capabilities:
	numCapabs = 32
	// length of the uint32 array that represents the bitset:
	bitsetLen = 1
)

const (
	// AccountNotify is the IRCv3 capability named "account-notify":
	// https://ircv3.net/specs/extensions/account-notify-3.1.html
	AccountNotify Capability = iota

	// AccountTag is the IRCv3 capability named "account-tag":
	// https://ircv3.net/specs/extensions/account-tag-3.2.html
	AccountTag Capability = iota

	// AwayNotify is the IRCv3 capability named "away-notify":
	// https://ircv3.net/specs/extensions/away-notify-3.1.html
	AwayNotify Capability = iota

	// Batch is the IRCv3 capability named "batch":
	// https://ircv3.net/specs/extensions/batch-3.2.html
	Batch Capability = iota

	// CapNotify is the IRCv3 capability named "cap-notify":
	// https://ircv3.net/specs/extensions/cap-notify-3.2.html
	CapNotify Capability = iota

	// ChgHost is the IRCv3 capability named "chghost":
	// https://ircv3.net/specs/extensions/chghost-3.2.html
	ChgHost Capability = iota

	// AccountRegistration is the draft IRCv3 capability named "draft/account-registration":
	// https://github.com/ircv3/ircv3-specifications/pull/435
	AccountRegistration Capability = iota

	// ChannelRename is the draft IRCv3 capability named "draft/channel-rename":
	// https://ircv3.net/specs/extensions/channel-rename
	ChannelRename Capability = iota

	// Chathistory is the proposed IRCv3 capability named "draft/chathistory":
	// https://github.com/ircv3/ircv3-specifications/pull/393
	Chathistory Capability = iota

	// EventPlayback is the proposed IRCv3 capability named "draft/event-playback":
	// https://github.com/ircv3/ircv3-specifications/pull/362
	EventPlayback Capability = iota

	// Languages is the proposed IRCv3 capability named "draft/languages":
	// https://gist.github.com/DanielOaks/8126122f74b26012a3de37db80e4e0c6
	Languages Capability = iota

	// Multiline is the proposed IRCv3 capability named "draft/multiline":
	// https://github.com/ircv3/ircv3-specifications/pull/398
	Multiline Capability = iota

	// Persistence is the proposed IRCv3 capability named "draft/persistence":
	// https://github.com/ircv3/ircv3-specifications/pull/503
	Persistence Capability = iota

	// Preaway is the proposed IRCv3 capability named "draft/pre-away":
	// https://github.com/ircv3/ircv3-specifications/pull/514
	Preaway Capability = iota

	// ReadMarker is the draft IRCv3 capability named "draft/read-marker":
	// https://github.com/ircv3/ircv3-specifications/pull/489
	ReadMarker Capability = iota

	// Relaymsg is the proposed IRCv3 capability named "draft/relaymsg":
	// https://github.com/ircv3/ircv3-specifications/pull/417
	Relaymsg Capability = iota

	// EchoMessage is the IRCv3 capability named "echo-message":
	// https://ircv3.net/specs/extensions/echo-message-3.2.html
	EchoMessage Capability = iota

	// Nope is the Ergo vendor capability named "ergo.chat/nope":
	// https://ergo.chat/nope
	Nope Capability = iota

	// ExtendedJoin is the IRCv3 capability named "extended-join":
	// https://ircv3.net/specs/extensions/extended-join-3.1.html
	ExtendedJoin Capability = iota

	// ExtendedMonitor is the IRCv3 capability named "extended-monitor":
	// https://ircv3.net/specs/extensions/extended-monitor.html
	ExtendedMonitor Capability = iota

	// InviteNotify is the IRCv3 capability named "invite-notify":
	// https://ircv3.net/specs/extensions/invite-notify-3.2.html
	InviteNotify Capability = iota

	// LabeledResponse is the IRCv3 capability named "labeled-response":
	// https://ircv3.net/specs/extensions/labeled-response.html
	LabeledResponse Capability = iota

	// MessageTags is the IRCv3 capability named "message-tags":
	// https://ircv3.net/specs/extensions/message-tags.html
	MessageTags Capability = iota

	// MultiPrefix is the IRCv3 capability named "multi-prefix":
	// https://ircv3.net/specs/extensions/multi-prefix-3.1.html
	MultiPrefix Capability = iota

	// SASL is the IRCv3 capability named "sasl":
	// https://ircv3.net/specs/extensions/sasl-3.2.html
	SASL Capability = iota

	// ServerTime is the IRCv3 capability named "server-time":
	// https://ircv3.net/specs/extensions/server-time-3.2.html
	ServerTime Capability = iota

	// SetName is the IRCv3 capability named "setname":
	// https://ircv3.net/specs/extensions/setname.html
	SetName Capability = iota

	// StandardReplies is the IRCv3 capability named "standard-replies":
	// https://github.com/ircv3/ircv3-specifications/pull/506
	StandardReplies Capability = iota

	// STS is the IRCv3 capability named "sts":
	// https://ircv3.net/specs/extensions/sts.html
	STS Capability = iota

	// UserhostInNames is the IRCv3 capability named "userhost-in-names":
	// https://ircv3.net/specs/extensions/userhost-in-names-3.2.html
	UserhostInNames Capability = iota

	// ZNCPlayback is the ZNC vendor capability named "znc.in/playback":
	// https://wiki.znc.in/Playback
	ZNCPlayback Capability = iota

	// ZNCSelfMessage is the ZNC vendor capability named "znc.in/self-message":
	// https://wiki.znc.in/Query_buffers
	ZNCSelfMessage Capability = iota
)

// `capabilityNames[capab]` is the string name of the capability `capab`
var (
	capabilityNames = [numCapabs]string{
		"account-notify",
		"account-tag",
		"away-notify",
		"batch",
		"cap-notify",
		"chghost",
		"draft/account-registration",
		"draft/channel-rename",
		"draft/chathistory",
		"draft/event-playback",
		"draft/languages",
		"draft/multiline",
		"draft/persistence",
		"draft/pre-away",
		"draft/read-marker",
		"draft/relaymsg",
		"echo-message",
		"ergo.chat/nope",
		"extended-join",
		"extended-monitor",
		"invite-notify",
		"labeled-response",
		"message-tags",
		"multi-prefix",
		"sasl",
		"server-time",
		"setname",
		"standard-replies",
		"sts",
		"userhost-in-names",
		"znc.in/playback",
		"znc.in/self-message",
	}
)
