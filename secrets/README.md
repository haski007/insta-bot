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
