{ config, self', inputs', pkgs, system, ... }:
{
  packages.cron = pkgs.buildGoApplication {
    name = "cron";
    src = ../.;
    modules = ../gomod2nix.toml;
    subpackages = [ "cmd/cron" ];
  };
}
