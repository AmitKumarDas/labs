# What:
# A reproducible nix based shell environment
#
# Usage:
# - Get into the shell:
# $ nix-shell
# - Run following commands from within above shell:
# $ tr ':' '\n' <<< "$PATH"
# $ which go
# $ go test
let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.05";
  pkgs = import nixpkgs { config = {}; overlays = []; };
in
pkgs.mkShellNoCC {
  packages = with pkgs; [
    go
  ];
}
