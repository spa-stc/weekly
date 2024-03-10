let
  devKeys = [ "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICz6aGtthhHrVYHrvk3BzCJTRFb3ppRB2MjHGI+eFteG foehammer127@gmail.com" ];

  machineKeys = [ "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICOvhkoMWwKybtPlrFPPBThhHnjhYoCqLKwKe3X+z6Fy" ];

  keys = devKeys ++ machineKeys;
in
rec {
  "secrets/cronenv.age".publicKeys = keys;
}
