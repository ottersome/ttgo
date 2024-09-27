# We need to create a package for this go project:
{
  description = "A simple stopwatch written in Go";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        lib = pkgs.lib;
      in
      {
        # Make a package for installation as a flake
        packages.ttgo = pkgs.buildGoModule {
          name = "ttgo";
          src = ./.;
          
          vendorSha256 = lib.fakeSha256;
          # vendorHash = "3f31691070c01bfc482d2524566aebc73496023b";
          meta = with lib; {
            description = "Go Time Tracker";
            homepage    = "https://github.com/ottersome/ttgo";
            maintainers = [ "ottersome" ];
          };
        };
      }
    );
}
