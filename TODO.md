# TODO
- [ ] Investigate levinboot for faster boot
    - might require moving root partition swap to initramfs init
    - write tiny initramfs init in go?
- [ ] nReal glasses
    - [ ] reverse engineer macos app to figure out 3D support
        - maybe force glasses into a widescreen mode?
    - [ ] connection is loose when back cover is on
- [ ] hudctl
    - [ ] Multiple windows support
        - add id field to app, and another one to windows
- [ ] hud_builder
    - [ ] allow building apps outside of the project
- [ ] hudOS
    - [ ] Allow remote upgrade of kernel/boot script
    - [ ] Resize root partitions on first boot
        - How will this affect OS upgrade?
        - Will probably require adding init to initramfs
    - [ ] Add tailscale
    - [ ] Bundle kiwmi,device-info into image
    - [ ] Dim/turn off screen based on proximity sensor
    - [x] start systemd target when glasses are plugged in
        - Allows apps to run only when glasses are plugged in
- [ ] Buildroot
    - [ ] Handle branch switches better
    - [ ] Detect when certain packages need to be updated, automatically add targets to build
        - Boot script/ hudctl
    - [ ] Move wifi configuration out of buildroot and into hud_builder devices
    - [ ] Automatically build the correct key into the image when deploying to a device
        - move key into pre image buildroot hook? key would update every build
    - [ ] Experiment with top level parrallel build
- [ ] Compositor
    - [ ] Make phone screen the primary monitor
    - [ ] Is touch input working?
    - [x] Support positioning apps on both screens
    - [ ] Restart compositor when windows.lua is updated
    - [ ] Could init.lua read the window catalog directly?
    - [ ] Get rid of cursor
    - [x] restart when glasses plugged in
        - [ ] fix hotplug to remove this hack
- [ ] Apps
    - [x] device-info: display ip adress/battery/time on phone screen
        - [ ] add charging indicator
        - to read battery: https://github.com/svenwltr/i3-statusbar/blob/master/upower/upower.go