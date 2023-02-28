const std = @import("std");

const ScanProtocolsStep = @import("vendor/zig/zig-wayland/build.zig").ScanProtocolsStep;

pub fn build(b: *std.build.Builder) void {
    // Standard target options allows the person running `zig build` to choose
    // what target to build for. Here we do not override the defaults, which
    // means any target is allowed, and the default is native. Other options
    // for restricting supported target set are available.
    const target = b.standardTargetOptions(.{});

    // Standard release options allow the person running `zig build` to select
    // between Debug, ReleaseSafe, ReleaseFast, and ReleaseSmall.
    const mode = b.standardReleaseOptions();

    const scanner = ScanProtocolsStep.create(b);
    scanner.addSystemProtocol("stable/xdg-shell/xdg-shell.xml");

    scanner.generate("wl_compositor", 1);
    scanner.generate("wl_shm", 1);
    scanner.generate("xdg_wm_base", 1);

    const wayland = std.build.Pkg{
        .name = "wayland",
        .source = .{ .generated = &scanner.result },
    };

    const exe = b.addExecutable("clock", "src/main.zig");
    exe.setTarget(target);
    exe.setBuildMode(mode);

    exe.step.dependOn(&scanner.step);
    exe.addPackage(wayland);
    exe.linkLibC();
    exe.linkSystemLibrary("wayland-client");

    scanner.addCSource(exe);
    
    exe.install();
}
