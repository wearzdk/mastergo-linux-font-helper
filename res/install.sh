#!/bin/bash

# Set installation path
INSTALL_PATH=$HOME/.local/bin

# Copy executable file to installation path
cp ./mastergo-font $INSTALL_PATH

# Set ownership and permissions for executable file
chown $USER:$USER $INSTALL_PATH/mastergo-font
chmod +x $INSTALL_PATH/mastergo-font

# Install systemd service file
cp ./mastergo-font.service /etc/systemd/system/

# Reload systemd daemon
systemctl daemon-reload

# Enable and start service
systemctl enable mastergo-font
systemctl start mastergo-font
