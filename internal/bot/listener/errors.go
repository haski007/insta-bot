package listener

const (
	ErrInternalServerError = "Internal server error, please contact the bot creator in the bot description"
	ErrNoCSGOPlayers       = "There are no cs go players registered in this chat!" +
		"\nTo register enter /reg_csgo_players {username_1} {username_2} ..."
	ErrNoPlayers = "There are no players registered in this chat!" +
		"\nTo register enter /reg_csgo_players {username_1} {username_2} ..."
	ErrNoArguments  = "This command requires more arguments, use /help"
	ErrWrongFormat  = "This command requires another format of arguments, use /help"
	ErrAccessDenied = "Access denied, please contact the bot creator in the bot description"

	ErrNoPUBGPlayers = "There are no PUBG players registered in this chat!" +
		"\nTo register enter /reg_pubg_players {username_1} {username_2} ..."
	ErrNoFinalsPlayers = "There are no The Finals players registered in this chat!" +
		"\nTo register enter /reg_finals_players {username_1} {username_2} ..."
)
