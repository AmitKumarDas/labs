# What:
# A reproducible nix based shell environment
#
# Usage:
# nix-shell
# - tr ':' '\n' <<< "$PATH"
# - which go
# - which jq
let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.05";
  pkgs = import nixpkgs { config = {}; overlays = []; };
in
pkgs.mkShellNoCC {
  packages = with pkgs; [
    go
    jq
  ];
}
