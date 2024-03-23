#!/bin/sh

echo "Downloading latest version of Goo"
mkdir /tmp/.goo
curl -L -o /tmp/.goo/goo https://github.com/TheWisePigeon/goo/releases/download/1.0/goo
curl -L -o /tmp/.goo/goo.service https://github.com/TheWisePigeon/goo/raw/main/goo.service
echo "Finished :)"
