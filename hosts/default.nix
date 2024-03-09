{ self, inputs, lib, ... }:
let
  mkSystem = extraModules: inputs.nixpkgs.lib.nixosSystem {
    system = "x86_64-linux";
    modules = with inputs; [
      # Agenix Secrets
      agenix.nixosModules.default
    ] ++ extraModules;
  };

  mkSystemDeployNode = configPath: hostname: sshUser: {
    inherit hostname;
    fastConnection = true;
    profiles = {
      system = {
        inherit sshUser;
        path = inputs.deploy-rs.lib.x86_64-linux.activate.nixos configPath;
        user = "root";
      };
    };
  };
in
{
  flake.nixosConfigurations = {
    cron = mkSystem [ ./cron ];
  };

  # Just Dump Here For Now, Until I Get Around To Making A Function.
  flake.deploy.nodes = {
    cron = mkSystemDeployNode self.nixosConfigurations.cron "cron" "root";
  };
}
