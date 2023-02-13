# hudOS
Experimental PinePhone distro to provide a heads-up display using the Nreal Air.

This is very much a work in progress and not ready to be used.

# Syncing go deps with bazel 
```
go mod tidy
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
bazel run //:gazelle
```

# Debugging notes
Restart HUD
```bash
sudo loginctl terminate-user hud
```

Service logs
```bash
sudo journalctl _SYSTEMD_USER_UNIT=hud_bluetooth.service
```

# Documentation
- Tow-boot
    - https://tow-boot.org/devices/pine64-pinephoneA64.html
- Arm Arch
    - https://github.com/dreemurrs-embedded/Pine64-Arch
    - NetworkManager
        - https://wiki.archlinux.org/title/NetworkManager
- Weston 
    - https://wayland.pages.freedesktop.org/weston/toc/running-weston.html
    - https://wiki.archlinux.org/title/Weston


# Phone setup
TODO: Add bits about updating ANX/Modem firmware?
NOTE: Setup was built on x86_64 Manjaro, running anywhere else might cause problems. 

1. Download images
    - `./images/download.sh`
2. Install Tow-Boot
    1. Flash `images/mmcboot.installer.img` to an SD card.
    2. Put the SD card in the pinephone, and boot.
    3. Install Tow-Boot to the eMMC.
    4. Remove the SD card.
3. Install Pine64-Arch
    1. Plug the pinephone into your computer, it should start booting.
    2. As soon as the pinephone vibrates start holding the volume up button.
    3. Stop holding button once the blue LED turns on.
    4. The pinephone's eMMC should now appear as a block device, `/dev/sdX`.
    5. Flash `images/archlinux-pinephone-barebone-20230203.img` to the eMMC.
    6. Reboot the pinephone.
4. Connecting to Wi-Fi
    1. Using a serial console, log in with `root:root`.
    2. Connect to Wi-Fi using NetworkManager.
        - `nmcli device wifi connect <SSID> password <PASSWORD>`
    3. Get ip address.
        - `ip addr`
5. Provisioning OS
    1. Test SSH connection using the default user: `alarm:123456`.
    2. Run the provisioning script.
        - You'll be asked to input a password several times, use `123456`.
        - Encrypting the SSH key is optional.
        - `./provision.sh <pinephone_ip>`
5. Deploying the HUD
    1. Run the deploy script.
        - `./deploy.sh <pinephone_ip>`