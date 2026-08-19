{
  description = "OPSX - Linux-first DevOps and DevSecOps CLI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            golangci-lint
          ];

          shellHook = ''
            echo ""
            echo "OPSX development environment"
            echo ""

            echo "Go:"
            go version

            echo ""
            echo "System:"
            uname -a

            echo ""
          '';
        };
      }
    );
}
