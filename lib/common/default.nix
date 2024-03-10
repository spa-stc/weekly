{ config, lib, pkgs, ... }:
let
  cfg = config.stc;
in
with lib; {
  options.stc = {
    sshRootLogin = mkEnableOption "Allow Root Login Over SSH";
  };

  config = {
    users.groups.stc = { };

    nix = {
      extraOptions = ''
        experimental-features = nix-command flakes
      '';

      gc = {
        automatic = true;
        dates = "weekly";
        options = "--delete-older-than 7d";
      };
    };

    nixpkgs.config = {
      allowUnfree = true;
    };

    users.users.root = mkIf cfg.sshRootLogin {
      openssh.authorizedKeys.keys = [
        # Foehammer
        "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICz6aGtthhHrVYHrvk3BzCJTRFb3ppRB2MjHGI+eFteG foehammer127@gmail.com"

        # Github actions runners.
        "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICn76lECi6dMUiz5nLnZLDLbFPSnDRoRr7gQfkH83xEQ stc@spa.edu"
      ];
    };
  };
}
