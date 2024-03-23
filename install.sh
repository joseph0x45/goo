#!/bin/sh

clean_up(){
  echo "Cleaning up"
  rm -rf /tmp/.goo
  if [ -f "/lib/systemd/system/goo.service" ]; then
    echo "Service file exists. Deleting..."
    if command -v sudo &> /dev/null; then
      sudo rm /lib/systemd/system/goo.service
    else
      rm /lib/systemd/system/goo.service
    fi
  fi
  echo "Installation failed. Please open an issue on https://github.com/TheWisePigeon/goo/issues to get help"
}

check_err() {
  if [ $? -ne 0 ]; then
    echo "$1 failed"
    clean_up
    exit 1
  fi
}

echo "Creating temporary directory at /tmp/.goo"
mkdir -p /tmp/.goo
check_err "Creating temporary directory"

echo "Downloading latest version of Goo"
curl -L -o /tmp/.goo/goo https://github.com/TheWisePigeon/goo/releases/download/1.0/goo
check_err "Downloading latest version of Goo"

echo "Downloading latest service file"
curl -L -o /tmp/.goo/goo.service https://github.com/TheWisePigeon/goo/raw/main/goo.service
check_err "Downloading latest service file"

echo "Making binary executable"
chmod u+x /tmp/.goo/goo
check_err "Making binary executable"

echo "Copying service file and binary"
if command -v sudo &> /dev/null; then
  sudo cp /tmp/.goo/goo.service /lib/systemd/system/
  sudo cp /tmp/.goo/goo /usr/local/bin/
else
  cp /tmp/.goo/goo.service /lib/systemd/system/
  cp /tmp/.goo/goo /usr/local/bin/
fi
check_err "Copying service file and/or binary file"

echo "Enabling and starting service"
if command -v sudo &> /dev/null; then
  sudo systemctl enable goo.service
  sudo systemctl start goo.service
else
  systemctl enable goo.service
  systemctl start goo.service
fi
check_err "Enabling and starting service"

echo "Cleaning up"
rm -rf /tmp/.goo
echo "All clear. Visit https://github.com/TheWisePigeon/goo to learn more"
