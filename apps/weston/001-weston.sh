#!/bin/bash

mkdir -p /var/log/hud/
chown hud:hud /var/log/hud/

mkdir -m 700 -p /opt/hud/run/
chown hud:hud /opt/hud/run/

systemctl enable hud.target hud-apps.target weston.service
systemctl start hud.target hud-apps.target weston.service
