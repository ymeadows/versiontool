{
  description = "Versiontool does simple semver math";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      rec {
        packages = rec {
          versiontool = pkgs.buildGoModule {
            pname = "versiontool";
            version = "1.0.0";

            src = builtins.path {
              path = ./.;
            };
            vendorHash = "sha256-4psCGkRw9GZIZlPMDpm70QukPX7tCtz2RHCxhb65B1c=";
          };

          default = versiontool;
        };

        defaultPackage = packages.default;

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # These tools are used by the Github Actions
            # We default to putting them in the devShell to ensure they can be used during development
            gzip
            skopeo
            trivy
            nix-update
            # If you need more tooling for this project,
            # add package names here
            go_1_25
            google-cloud-sdk
          ];
        };
      }
    );
}
