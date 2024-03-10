{ self', inputs', pkgs, system, ... }:
rec {
  packages.cron = pkgs.buildGoApplication {
    name = "cron";
    src = ../.;
    modules = ../gomod2nix.toml;
    subpackages = [ "cmd/cron" ];
  };

  packages.cron-docker = pkgs.dockerTools.buildLayeredImage {
    name = "newsletter/cron";
    tag = "latest";
    config = {
      Env = [ "PRODUCTION=true" ];
      Cmd = [ "${packages.cron}/bin/cron" ];
    };
  };
}
