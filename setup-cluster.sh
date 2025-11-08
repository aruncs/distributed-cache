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
trap "echo '🛑 Stopping...'; kill $PID1 $PID2 $PID3 $PID4 $PID5" EXIT

wait
