#!/usr/bin/env bash
set -euo pipefail

SERVICE_PREFIX=${SERVICE_PREFIX:-talktocow}
BACKEND_PORT=${BACKEND_PORT:-12001}
FRONTEND_PORT=${FRONTEND_PORT:-3080}
RUN_USER=${RUN_USER:-${SUDO_USER:-$USER}}
RUN_GROUP=${RUN_GROUP:-$(id -gn "$RUN_USER")}
REPO_DIR=${REPO_DIR:-$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)}
ENV_FILE=${ENV_FILE:-$REPO_DIR/.env}
INSTALL_FRONTEND=${INSTALL_FRONTEND:-1}
INSTALL_BACKEND=${INSTALL_BACKEND:-1}
START_SERVICES=${START_SERVICES:-1}
SYSTEMD_DIR=${SYSTEMD_DIR:-/etc/systemd/system}

find_bin() {
  local name=$1
  command -v "$name" 2>/dev/null || sudo -u "$RUN_USER" bash -lc "command -v $name" 2>/dev/null || true
}

NODE_BIN=${NODE_BIN:-$(find_bin node)}
NPM_BIN=${NPM_BIN:-$(find_bin npm)}
GO_BIN=${GO_BIN:-$(find_bin go)}

usage() {
  cat <<USAGE
Install Talktocow from this local checkout as systemd services.

Environment overrides:
  SERVICE_PREFIX=talktocow       Service name prefix
  RUN_USER=$RUN_USER             User to run services as
  REPO_DIR=$REPO_DIR             Local checkout to run
  ENV_FILE=$ENV_FILE             Backend env file
  BACKEND_PORT=$BACKEND_PORT     Backend port documented for status output
  FRONTEND_PORT=$FRONTEND_PORT   Frontend dev server port
  INSTALL_BACKEND=1              Install backend service
  INSTALL_FRONTEND=1             Install frontend service
  START_SERVICES=1               Enable and start/restart services
  NODE_BIN=/path/to/node         Node binary for frontend service PATH
  NPM_BIN=/path/to/npm           npm binary
  GO_BIN=/path/to/go             Go binary

Examples:
  sudo ./deploy/install-local.sh
  sudo RUN_USER=puppe REPO_DIR=/home/puppe/talktocow ./deploy/install-local.sh
USAGE
}

if [[ ${1:-} == "-h" || ${1:-} == "--help" ]]; then
  usage
  exit 0
fi

if [[ $EUID -ne 0 ]]; then
  echo "Please run with sudo/root so systemd units can be installed." >&2
  exit 1
fi

if [[ ! -d "$REPO_DIR" ]]; then
  echo "REPO_DIR does not exist: $REPO_DIR" >&2
  exit 1
fi

if [[ ! -f "$ENV_FILE" ]]; then
  if [[ -f "$REPO_DIR/example.env" ]]; then
    cp "$REPO_DIR/example.env" "$ENV_FILE"
    chown "$RUN_USER:$RUN_GROUP" "$ENV_FILE"
    chmod 600 "$ENV_FILE"
    echo "Created $ENV_FILE from example.env. Edit it if database/key settings differ."
  else
    echo "Missing env file: $ENV_FILE" >&2
    exit 1
  fi
fi

if [[ "$INSTALL_BACKEND" == "1" && -z "$GO_BIN" ]]; then
  echo "Could not find go; set GO_BIN=/path/to/go" >&2
  exit 1
fi

if [[ "$INSTALL_FRONTEND" == "1" ]]; then
  if [[ -z "$NODE_BIN" || -z "$NPM_BIN" ]]; then
    echo "Could not find node/npm; set NODE_BIN and NPM_BIN" >&2
    exit 1
  fi

  if [[ ! -d "$REPO_DIR/frontend/node_modules" ]]; then
    echo "Installing frontend dependencies..."
    sudo -u "$RUN_USER" env HOME="$(getent passwd "$RUN_USER" | cut -d: -f6)" \
      "$NPM_BIN" --prefix "$REPO_DIR/frontend" install
  fi
fi

if [[ "$INSTALL_BACKEND" == "1" ]]; then
  echo "Downloading Go dependencies..."
  sudo -u "$RUN_USER" env HOME="$(getent passwd "$RUN_USER" | cut -d: -f6)" \
    "$GO_BIN" -C "$REPO_DIR" mod download
fi

if [[ ! -f "$REPO_DIR/jwt_private_key.pem" || ! -f "$REPO_DIR/jwt_public_key.pem" ]]; then
  if command -v openssl >/dev/null 2>&1; then
    echo "Generating JWT keypair..."
    sudo -u "$RUN_USER" openssl genrsa -out "$REPO_DIR/jwt_private_key.pem" 2048
    sudo -u "$RUN_USER" openssl rsa -in "$REPO_DIR/jwt_private_key.pem" -pubout -out "$REPO_DIR/jwt_public_key.pem"
    chmod 600 "$REPO_DIR/jwt_private_key.pem"
  else
    echo "JWT keys are missing and openssl was not found. Create jwt_private_key.pem and jwt_public_key.pem before starting." >&2
  fi
fi

backend_unit="$SYSTEMD_DIR/${SERVICE_PREFIX}-backend.service"
frontend_unit="$SYSTEMD_DIR/${SERVICE_PREFIX}-frontend.service"
user_home=$(getent passwd "$RUN_USER" | cut -d: -f6)
path_parts=()
[[ -n "$NODE_BIN" ]] && path_parts+=("$(dirname "$NODE_BIN")")
[[ -n "$NPM_BIN" ]] && path_parts+=("$(dirname "$NPM_BIN")")
[[ -n "$GO_BIN" ]] && path_parts+=("$(dirname "$GO_BIN")")
IFS=:
service_path="${path_parts[*]}:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
unset IFS

if [[ "$INSTALL_BACKEND" == "1" ]]; then
  cat > "$backend_unit" <<UNIT
[Unit]
Description=Talktocow backend (local checkout)
After=network-online.target docker.service
Wants=network-online.target

[Service]
Type=simple
User=$RUN_USER
Group=$RUN_GROUP
WorkingDirectory=$REPO_DIR
EnvironmentFile=$ENV_FILE
Environment=HOME=$user_home
Environment=PATH=$service_path
ExecStart=$GO_BIN run .
Restart=on-failure
RestartSec=3

[Install]
WantedBy=multi-user.target
UNIT
fi

if [[ "$INSTALL_FRONTEND" == "1" ]]; then
  cat > "$frontend_unit" <<UNIT
[Unit]
Description=Talktocow frontend (local checkout)
After=network-online.target ${SERVICE_PREFIX}-backend.service
Wants=network-online.target

[Service]
Type=simple
User=$RUN_USER
Group=$RUN_GROUP
WorkingDirectory=$REPO_DIR/frontend
Environment=HOME=$user_home
Environment=PATH=$service_path
ExecStart=$NPM_BIN run dev -- --host 0.0.0.0 --port $FRONTEND_PORT
Restart=on-failure
RestartSec=3

[Install]
WantedBy=multi-user.target
UNIT
fi

systemctl daemon-reload

if [[ "$START_SERVICES" == "1" ]]; then
  [[ "$INSTALL_BACKEND" == "1" ]] && systemctl enable --now "${SERVICE_PREFIX}-backend.service"
  [[ "$INSTALL_BACKEND" == "1" ]] && systemctl restart "${SERVICE_PREFIX}-backend.service"
  [[ "$INSTALL_FRONTEND" == "1" ]] && systemctl enable --now "${SERVICE_PREFIX}-frontend.service"
  [[ "$INSTALL_FRONTEND" == "1" ]] && systemctl restart "${SERVICE_PREFIX}-frontend.service"
fi

cat <<DONE
Installed Talktocow local checkout services.

Backend:  ${SERVICE_PREFIX}-backend.service  (port $BACKEND_PORT)
Frontend: ${SERVICE_PREFIX}-frontend.service (port $FRONTEND_PORT)
Repo:     $REPO_DIR
Env:      $ENV_FILE

Useful commands:
  sudo systemctl status ${SERVICE_PREFIX}-backend ${SERVICE_PREFIX}-frontend
  sudo journalctl -u ${SERVICE_PREFIX}-backend -f
  sudo journalctl -u ${SERVICE_PREFIX}-frontend -f
DONE
