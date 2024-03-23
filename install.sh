#!/bin/sh

echo "Downloading latest version of Goo"
mkdir /tmp/.goo
curl -o /tmp/.goo/goo https://github.com/TheWisePigeon/goo/releases/download/1.0/goo
curl -o /tmp/.goo/goo.service https://github.com/TheWisePigeon/goo/raw/1.0/goo.service
echo "Finished :)"
