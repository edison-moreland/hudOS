#!/bin/bash

# Post install for weston

systemctl -M hud@.host --user enable weston-session.target weston.service weston.socket
systemctl -M hud@.host --user start weston-session.target weston.service weston.socket
