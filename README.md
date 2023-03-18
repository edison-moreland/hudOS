# hudOS
Experimental PinePhone Pro distro to provide a heads-up display using the Nreal Air.
This is very much a work in progress and not ready to be used.

In theory PinePhone support should be possible, but I haven't been able to get the Nreal glasses to work.

# Phone setup
NOTE: Setup was built on x86_64 Manjaro. This should work on any modern linux distro, YMMV.

Dependencies
  - [jq](https://stedolan.github.io/jq/)
  - [Buildroot Dependencies](https://buildroot.org/downloads/manual/manual.html#requirement)

1. Setup tools
   1. Download dependencies/tools
      - `./hb vendor`
   2. Create a new device
      - `./hb devices new`
2. Install Tow-Boot (optional)
   - All PinePhone Pros sold after July 2022 should come with Tow-Boot preinstalled to the SPI flash.
   1. Flash `.build/vendor/towboot/spi.installer.img` to an SD card.
   2. Put the SD card in the PinePhone, and boot.
   3. Install Tow-Boot to the SPI.
   4. Remove the SD card.
3. Building hudOS
   1. Configure OS 
      - `./hb buildroot build nconfig`
      - Under `External options`
         - Select `hudOS deploy user`
         - Fill in WiFi settings
   2. Build OS 
      - `./hb buildroot build`
      - This step will take 30 minutes to an hour and consume significant amounts of system resources.
4. Flash the PinePhone 
   1. Plug the PinePhone into your computer, it should start booting.
   2. As soon as the PinePhone vibrates start holding the volume up button.
   3. Stop holding button once the blue LED turns on.
   4. The PinePhone's eMMC should now appear as a block device, `/dev/sdX`.
   5. Flash `./output/hudOS.img` to the eMMC.
   6. Reboot the PinePhone.
5. Deploying the HUD
   - If the PinePhone was able to successfully connect to the WiFi, you should see it's ip address printed to the screen on boot.
   1. Add ip to device configuration
      - `./hb devices host <device-name> <ip-address>`
   2. Deploy apps to the phone
      - `./hb deploy`

# Notes
## Development
Deploy a new buildroot image to the phone:
```
./hb buildroot build # Buildroot produces a new rootfs
./hb buildroot deploy # Rootfs is written to an unused partition on the phone
# The phone should now reboot into the new image
# Note: This does not update the kernel or boot script
```

Deploy a single app the the phone:
```
./hb deploy --only <app_name>
```

Incorporate hudctl changes into the base image:
```
./hb buildroot build hudctl-rebuild # Clean the output of the previous build
./hb buildroot build # rebuild the image
```

## Documentation
- Tow-boot
    - https://tow-boot.org/devices/pine64-pinephonePro.html
- Weston
    - https://wayland.pages.freedesktop.org/weston/toc/running-weston.html
    - https://wiki.archlinux.org/title/Weston
