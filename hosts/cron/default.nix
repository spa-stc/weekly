{ config, pkgs, lib, ... }: {
  imports = [
    # Required for the system to boot.
    ../../lib/hardware/qemu.nix

    # Root filesystem setup
    ./filesystems.nix
  ];


  services.openssh = {
    enable = true;
    settings = {
      PasswordAuthentication = false;
    };
  };

  networking = {
    hostName = "cron";
    firewall = {
      enable = true;
      allowedTCPPorts = [ 22 ];
    };
  };

  swapDevices = [{
    device = "/var/lib/swap";
    size = 1024 * 1; # 1 GB Swap.
  }];

  system.stateVersion = "23.05";

  boot = {
    growPartition = true;
    kernelParams = [ "console=ttyS0" "panic=1" "boot.panic_on_fail" ];
    initrd.kernelModules = [ "virtio_scsi" ];
    kernelModules = [ "virtio_pci" "virtio_net" ];
    loader = {
      grub.device = "/dev/vda";
      timeout = 0;
      grub.configurationLimit = 0;
    };
  };

  environment.systemPackages = with pkgs; [ vim ];
} 
