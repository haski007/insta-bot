package listener

import "strings"

func (rcv *InstaBotService) IsCreator(uid int64) bool {
	return rcv.creatorID == uid
}

func (rcv *InstaBotService) parseCommandArgs(text string) []string {
	// Split the message text into parts
	parts := strings.Fields(text)

	// If we have no parts or just the command, return empty slice
	if len(parts) <= 1 {
		return []string{}
	}

	// Return all parts after the command (skip the first element)
	return parts[1:]
}
