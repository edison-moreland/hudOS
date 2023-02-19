#!/bin/bash

systemctl enable hud_{app}.service

# shellcheck disable=SC2050
if [[ "{setcap}" != "" ]]
then
  setcap '{setcap}' {prefix}/bin/{app}_arm
fi