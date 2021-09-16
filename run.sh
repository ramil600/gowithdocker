#!/bin/sh

echo "######################################"
echo "Waiting for Eureka"
echo "######################################"

while ! `nc -z discovery 8761`; do sleep 3; done;

echo "######################################"
echo "Waiting for Proxy"
echo "######################################"

while ! `nc -z proxy 8000`; do sleep 3; done;

echo "***   STARTING   DISCOVERY   SERVICE   ***"
./dispatcher