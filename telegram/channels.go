package telegram

// Channel represents a Telegram channel
type Channel string

// Predefined channels
const (
	ChannelSugar   Channel = "sugar"
	ChannelMBank   Channel = "mbank"
	ChannelLab     Channel = "lab"
	ChannelCommits Channel = "commits"
)

// String returns the string representation of the channel
func (c Channel) String() string {
	return string(c)
}

// AllChannels returns a slice of all predefined channels
func AllChannels() []Channel {
	return []Channel{
		ChannelSugar,
		ChannelMBank,
		ChannelLab,
		ChannelCommits,
	}
}

// IsValidChannel checks if the provided channel is one of the predefined channels
func IsValidChannel(channel Channel) bool {
	switch channel {
	case ChannelSugar, ChannelMBank, ChannelLab, ChannelCommits:
		return true
	default:
		return false
	}
}
