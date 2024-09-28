{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
, mkGoEnv ? pkgs.mkGoEnv
, gomod2nix ? pkgs.gomod2nix
, gopls ? pkgs.gopls
, git ? pkgs.gopls
}:

let
  goEnv = mkGoEnv { pwd = ./.; };
in
pkgs.mkShell {
  packages = [
    goEnv
    gomod2nix
    gopls # Language Server
    git 
  ];
  shellHook = ''
    # Get pkgs path to zsh
    export SHELL=${pkgs.zsh}/bin/zsh
    export INFLAKE=1
    # export GOPATH=$PWD
    export PATH=$PATH:$GOPATH/bin
    # Now actually enter zsh
    exec zsh
    '';
}
