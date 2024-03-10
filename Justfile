default: 

deploy-all:
  nix develop -c deploy

deploy NODE:
  nix develop -c deploy .#{{NODE}}
