# Talktocow local service install

`install-local.sh` installs the current checkout as systemd services, similar to a live local workspace deployment:

- `talktocow-backend.service` runs `go run .` from the repo root.
- `talktocow-frontend.service` runs the Vite dev server from `frontend/` on port `3080`.
- Backend configuration is read from `.env` in the repo root.

```bash
sudo ./deploy/install-local.sh
```

Common overrides:

```bash
sudo RUN_USER=puppe REPO_DIR=/home/puppe/.openclaw/workspace/talktocow ./deploy/install-local.sh
sudo SERVICE_PREFIX=talktocow-dev FRONTEND_PORT=3080 ./deploy/install-local.sh
```

The database still needs to be available separately, for example with:

```bash
docker compose up -d database
```

Useful service commands:

```bash
sudo systemctl status talktocow-backend talktocow-frontend
sudo journalctl -u talktocow-backend -f
sudo journalctl -u talktocow-frontend -f
```

## nginx reverse proxy

If Talktocow is served through nginx, raise nginx's default request body limit so browser image uploads can reach the backend. The app still enforces its own per-file upload limit.

Use `deploy/nginx-talktocow.conf` as a starting point, or add this inside the HTTPS `server` block:

```nginx
client_max_body_size 16m;
```

Then validate and reload nginx:

```bash
sudo nginx -t
sudo systemctl reload nginx
```
