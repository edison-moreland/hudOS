# hudOS
Experimental PinePhone Pro distro to provide a heads-up display using the Nreal Air.
This is very much a work in progress and not ready to be used.

In theory PinePhone support should be possible, but I haven't been able to get the Nreal glasses to work.

# Phone setup
NOTE: Setup was built on x86_64 Manjaro. This should work on any modern linux distro, YMMV.

Dependencies
  - [jq](https://stedolan.github.io/jq/)
  - [Buildroot Dependencies](https://buildroot.org/downloads/manual/manual.html#requirement)

1. Download dependencies
    - `./update_vendor.sh`
2. Install Tow-Boot (optional)
    - All PinePhone Pros sold after July 2022 should come with Tow-Boot preinstalled to the SPI flash.
    1. Flash `.build/vendor/towboot/spi.installer.img` to an SD card.
    2. Put the SD card in the PinePhone, and boot.
    3. Install Tow-Boot to the SPI.
    4. Remove the SD card.
3. Building hudOS
   1. Download buildroot
      - `./buildroot/setup.sh`
   2. Configure OS 
      - `./buildroot/build.sh nconfig`
      - Under `External options`
         - Select `hudOS deploy user`
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
## Development
Deploy a new buildroot image to the phone:
```
./buildroot/build.sh # Buildroot produces a new rootfs
./buildroot/upgrade.sh <pinephone_ip> # Rootfs is written to an unused partition on the phone
# The phone should now reboot into the new image
```

Deploy a single app the the phone:
```
./deploy.sh <pinephone_ip> <app_name>
```

## Documentation
- Tow-boot
    - https://tow-boot.org/devices/pine64-pinephonePro.html
- Weston
    - https://wayland.pages.freedesktop.org/weston/toc/running-weston.html
    - https://wiki.archlinux.org/title/Weston
