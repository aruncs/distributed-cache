#!/bin/bash

set -e  # stop on first error

echo "Building load balancer..."
cd load-balancer
go build -o ../bin/load-balancer .
echo "Build completed for load-balancer"

echo "🚀 Starting load balancer..."
# Run Go servers in background
./bin/load-balancer &
PID1=$!

echo "🚀 Started load balancer... Process id $PID1"
echo "Building nodes..."
cd ../node
go build -o ../bin/node .
echo "Build completed for nodes"

cd ../
echo "🚀 Starting nodes..."
./bin/node 8090 &
PID2=$!

./bin/node 8091 &
PID3=$!

./bin/node 8092 &
PID4=$!

echo "🚀 all nodes started"

# Start React app
cd monitoring-dashboard
npm run start &
PID5=$!

# Cleanup when user stops the script
#trap "echo '🛑 Stopping...'; kill $PID1 $PID2 $PID3 $PID4 $PID5" EXIT

cleanup() {

#!/bin/bash

# List of ports to kill
PORTS=(8080 8090 8091 8092 3000)

echo "🔍 Checking for processes on ports: ${PORTS[*]}"

for PORT in "${PORTS[@]}"; do
  # Find process ID (PID) using lsof
  PID=$(lsof -ti tcp:$PORT)

  if [ -n "$PID" ]; then
    echo "⚠️  Killing process $PID on port $PORT"
    kill -9 $PID
  else
    echo "✅ No process found on port $PORT"
  fi
done
}

trap cleanup EXIT INT TERM

wait
