#!/bin/sh

mkdir -p /tmp/.goo
echo "Downloading latest version of Goo"
curl -L -o /tmp/.goo/goo https://github.com/TheWisePigeon/goo/releases/download/1.0/goo
curl -L -o /tmp/.goo/goo.service https://github.com/TheWisePigeon/goo/raw/main/goo.service
echo "Making binary executable"
chmod u+x /tmp/.goo/goo
if command -v sudo &> /dev/null; then
    sudo cp /tmp/.goo/goo.service /lib/systemd/system/
    sudo systemctl enable goo.service
    sudo systemctl start goo.service 
else
    cp /tmp/.goo/goo.service /lib/systemd/system/
    systemctl enable goo.service
    systemctl start goo.service 
fi
echo "Cleaning up"
rm -rf /tmp/.goo
echo "All clear. Visit https://github.com/TheWisePigeon/goo to learn more"
