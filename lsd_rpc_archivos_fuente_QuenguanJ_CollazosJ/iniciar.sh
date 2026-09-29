#!/bin/bash
RAIZ="$(cd "$(dirname "$0")" && pwd)"

# RabbitMQ debe estar activo (localhost:5672)
if ! (echo > /dev/tcp/127.0.0.1/5672) 2>/dev/null; then
  echo "RabbitMQ no responde en localhost:5672. Inícialo primero:"
  echo "  sudo systemctl start rabbitmq-server"
  echo "  o: docker run -d --name rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management"
  exit 1
fi

abrir() {  # abrir "Título" "carpeta" "comando"
  [ -d "$RAIZ/$2" ] || { echo "Falta la carpeta $2"; exit 1; }
  local cmd="cd '$RAIZ/$2' && $3; echo; echo '[terminó: $1]'; exec bash"
  if   command -v gnome-terminal >/dev/null; then gnome-terminal --title="$1" -- bash -c "$cmd"
  elif command -v konsole        >/dev/null; then konsole -p "tabtitle=$1" -e bash -c "$cmd" &
  elif command -v xfce4-terminal >/dev/null; then xfce4-terminal --title="$1" -x bash -c "$cmd" &
  elif command -v xterm          >/dev/null; then xterm -T "$1" -e bash -c "$cmd" &
  else echo "No encontré ninguna terminal (gnome-terminal, konsole, xfce4-terminal, xterm)"; exit 1
  fi
}

# Servidores (primero estadísticas, para que la cola exista antes de publicar)
abrir "Estadisticas" "ServidorDeEstadisticas"   "go run ."
sleep 2
abrir "Streaming"    "ServidorDeStreaming"      "go run ./main"
abrir "Metadata"     "ServidorMetadataDeAudios" "go run ."
abrir "ServidorAudios" "ServidorDeAudios"       "go run ./main"
sleep 4

# Programas interactivos
abrir "Cliente"       "Cliente"        "go run ."
abrir "Administrador" "Administrador"  "go run ."
