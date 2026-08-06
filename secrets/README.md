# Instagram session (optional)

**Automatic (often fails on VPS):** instloader tries password login with `.env` credentials. Instagram frequently blocks this with `Unexpected null login result` — use manual import below instead.

**Recommended on VPS:** export session on your PC, then copy here:

```bash
pip3 install instaloader browser-cookie3
# log into instagram.com in Chrome first, then:
export INSTAGRAM_USERNAME=your_user
python3 scripts/export_instagram_session.py --browser chrome
scp session.json user@vps:~/projects/insta-bot/secrets/session.json
docker compose up -d instloader
```

Password-only login often fails (`no sessionid`) — use `--browser chrome` or `INSTAGRAM_SESSIONID`.

**Or** set `INSTAGRAM_SESSIONID` in `.env` (cookie value from browser DevTools → Application → Cookies → sessionid).

**Manual import:** copy a valid Instaloader session file here as `session.json` (pickle from `save_session_to_file`).

Do **not** bind-mount `./session.json` at the project root — Docker may create a **directory** instead of a file.

Backup session from the running container:

```bash
docker cp insta-bot-instloader-1:/data/session.json ./secrets/session.json
```

## Multiple accounts (rotation / failover)

You can configure additional Instagram accounts with numeric suffixes in `.env`; instloader rotates requests across every account with a valid session and automatically retries on another account when one gets rate-limited (`login_required`, `429`, `403`, ...):

```bash
INSTAGRAM_USERNAME_2="second_account"
INSTAGRAM_PASSWORD_2="..."
INSTAGRAM_PROXY_2="http://user:pass@isp.decodo.com:10002"   # recommended: a different IP per account
# INSTAGRAM_SESSIONID_2="..."           # optional, same as account 1
# INSTAGRAM_SESSION_FILE_2="/data/session_2.json"  # default shown

INSTAGRAM_USERNAME_3="third_account"
# ...up to INSTAGRAM_USERNAME_10
```

If `INSTAGRAM_PROXY_N` is not set, that account falls back to the shared `INSTAGRAM_PROXY` — fine for testing, but running multiple accounts through the same IP makes them easier to correlate/ban together, so a dedicated proxy per account is recommended for real use.

Export/import a session per extra account the same way as account 1, just point `SESSION_OUT` and the host import path at that account's suffix:

```bash
export INSTAGRAM_USERNAME=second_account
SESSION_OUT=session_2.json python3 scripts/export_instagram_session.py --browser chrome
scp session_2.json user@vps:~/projects/insta-bot/secrets/session_2.json
```

Startup logs report the validity of every configured account; the service only refuses to start if **none** of them have a usable session.
