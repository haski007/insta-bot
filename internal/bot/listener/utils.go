package listener

func (rcv *InstaBotService) IsCreator(uid int64) bool {
	return rcv.creatorID == uid
}
