default: 

deploy-all:
  nix develop -c deploy

deploy NODE:
  nix develop -c deploy .#{{NODE}}

build-cron-docker:
  nix build .#cron-docker
  docker load < result
