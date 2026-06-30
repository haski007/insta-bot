package listener

import "strings"

func (rcv *InstaBotService) IsCreator(uid int64) bool {
	return rcv.creatorID == uid
}

func (rcv *InstaBotService) parseCommandArgs(text string) []string {
	parts := strings.Fields(text)
	if len(parts) <= 1 {
		return []string{}
	}
	return parts[1:]
}