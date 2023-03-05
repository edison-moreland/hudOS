# hudOS
Experimental PinePhone distro to provide a heads-up display using the Nreal Air.
This is very much a work in progress and not ready to be used.

HudOS should be compatible with all PinePhone hardware revisions. 
Development is being done on a PinePhone v1.2b, with some testing done on a v1.1.
Any revision before v1.2a will need [hardware modifications](https://wiki.pine64.org/wiki/PinePhone_v1.2#USB) to allow use of the Nreal Air. 

# Phone setup
NOTE: Setup was built on x86_64 Manjaro. This should work on any modern linux distro, YMMV.

Dependencies
  - [jq](https://stedolan.github.io/jq/)
  - [Buildroot Dependencies](https://buildroot.org/downloads/manual/manual.html#requirement)

1. Download dependencies
    - `./update_vendor.sh`
2. Install Tow-Boot
    1. Flash `.build/vendor/towboot/mmcboot.installer.img` to an SD card.
    2. Put the SD card in the PinePhone, and boot.
    3. Install Tow-Boot to the eMMC.
    4. Remove the SD card.
3. Building hudOS
   1. Download buildroot
      - `./buildroot/setup.sh`
   2. Configure OS 
      - `./buildroot/build.sh nconfig`
      - Under `External options`, select `hud deploy user` and `hud network`
      - Fill in WiFi settings
   3. Build OS 
      - `./buildroot/build.sh`
      - This step will take 30 minutes to an hour and consume significant amounts of system resources.
4. Flash the PinePhone 
   1. Plug the PinePhone into your computer, it should start booting.
   2. As soon as the PinePhone vibrates start holding the volume up button.
   3. Stop holding button once the blue LED turns on.
   4. The PinePhone's eMMC should now appear as a block device, `/dev/sdX`.
   5. Flash `buildroot/hudOS.img` to the eMMC.
   6. Reboot the PinePhone.
5. Deploying the HUD
   1. If the PinePhone was able to successfully connect to the WiFi, you should see it's ip address printed to the screen on boot.
   2. Run the deploy script.
      - `./deploy.sh <pinephone_ip>`

# Notes
## Documentation
- Tow-boot
    - https://tow-boot.org/devices/pine64-pinephoneA64.html
- Weston
    - https://wayland.pages.freedesktop.org/weston/toc/running-weston.html
    - https://wiki.archlinux.org/title/Weston
