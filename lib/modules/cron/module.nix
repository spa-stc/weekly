self:
{ config, lib, pkgs, ... }:
let
  cfg = config.stc.cron;
in
with lib; {
  options.stc.cron = {
    enable = mkEnableOption "Enable";
  };

  config = mkIf cfg.enable {
    age.secrets = {
      cronenv = {
        file = ../../../secrets/cronenv.age;
        owner = "weeklycron";
        mode = "600";
        path = "/srv/weeklycron/.env";
        group = "stc";
      };
    };

    users.users.weeklycron = {
      createHome = true;
      group = "stc";
      description = "User for the weekly cron service";
      isSystemUser = true;
      home = "/srv/weeklycron";
    };

    systemd.services.weeklycron = {
      wantedBy = [ "multi-user.target" ];

      serviceConfig = {
        User = "weeklycron";
        Group = "stc";
        Restart = "on-failure";
        WorkingDirectory = "/srv/weeklycron";
        RestartSec = "30s";
        Type = "exec";
        StandardOutput = "/var/cronservice.log";
        StandardError = "/var/cronservice-error.log";
      };

      script =
        let
          cron = self.packages.${pkgs.system}.cron;
          filepath = config.age.secrets.cronenv.path;
        in
        ''
          [ -f ${filepath} ] && export $(cat ${filepath} | xargs)
          export PRODUCTION=true
          exec ${cron}/bin/cron
        '';
    };
  };
}
