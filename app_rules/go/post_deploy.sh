#!/bin/bash

systemctl -M hud@.host --user enable hud_{app}.service
systemctl -M hud@.host --user start hud_{app}.service

# shellcheck disable=SC2050
if [[ "{setcap}" != "" ]]
then
  setcap '{setcap}' {prefix}/bin/{app}_arm
fi