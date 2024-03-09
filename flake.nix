{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

    # Framework For Defining "Flake Modules" with imports
    flake-parts.url = "github:hercules-ci/flake-parts";

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = inputs@{ self, flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" ];

      perSystem = { config, self', inputs', pkgs, system, ... }: {
        imports = [
          ./cmd
        ];


        _module.args.pkgs = import inputs.nixpkgs {
          inherit system;
          config.allowUnfree = true;
          overlays =
            let
              graft = pkgs: pkg:
                pkg.override { buildGoModule = pkgs.buildGo121Module; };
            in
            [
              inputs.gomod2nix.overlays.default
              (final: prev: {
                go = prev.go_1_21;
                go-tools = graft prev prev.go-tools;
                gotools = graft prev prev.gotools;
                gopls = graft prev prev.gopls;
              })
            ];
        };

        # Development shell for the backend side of the project.
        devShells.backend = pkgs.mkShell {
          buildInputs = with pkgs; [
            dive
            inputs'.gomod2nix.packages.default
            go
            go-tools
            gotools
            gopls
          ];
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            just
          ];
        };
      };
    };
}
