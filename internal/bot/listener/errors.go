package listener

const (
	ErrInternalServerError = "Internal server error, please contact the bot creator in the bot description"
	ErrNoCSGOPlayers       = "There are no cs go players registered in this chat!" +
		"\nTo register enter /reg_csgo_players {username_1} {username_2} ..."
	ErrNoPlayers = "There are no players registered in this chat!" +
		"\nTo register enter /reg_csgo_players {username_1} {username_2} ..."
	ErrNoArguments = "This command requires one or more arguments"

	ErrNoPUBGPlayers = "There are no PUBG players registered in this chat!" +
		"\nTo register enter /reg_pubg_players {username_1} {username_2} ..."
)
