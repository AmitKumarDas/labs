# Bootstrap:
# - nix flake init
# - nix flake lock
#
# Usage:
# - nix flake show
#
# Troubleshooting:
# - zsh: command not found: nix
# -- cat /etc/bash.bashrc
# -- source /etc/bash.bashrc
# -- Alternative: . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
#
# References:
# - https://github.com/nix-systems/nix-systems
# - https://github.com/nix-systems/default
{
  description = "Display all systems of this flake";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.05";
    systems.url = "github:nix-systems/default";
  };


#  outputs = { self, nixpkgs }: {
#    packages.x86_64-linux.hello = nixpkgs.legacyPackages.x86_64-linux.hello;
#    packages.x86_64-linux.default = self.packages.x86_64-linux.hello;
#  };

  outputs = { self, systems, nixpkgs }:
    let
        eachSystem = nixpkgs.lib.genAttrs (import systems);
    in
    {
        packages = eachSystem (system: {
            hello = nixpkgs.legacyPackages.${system}.hello;
        });
    };
}
