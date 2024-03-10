{ self, inputs, lib, ... }: {
  flake.nixosModules.cron = import ./module.nix self;
}
