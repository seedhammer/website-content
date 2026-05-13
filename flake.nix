{
  description = "SeedHammer website";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
    utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      utils,
    }:
    utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        formatter = pkgs.nixpkgs-fmt;
        packages = {
          default = pkgs.writeShellScriptBin "run-viewer" ''
            sigint_handler()
            {
              kill $PID
              exit
            }

            trap sigint_handler SIGINT

            function watch() {
            		${pkgs.fswatch}/bin/fswatch -1 -r -e '.git/*' -e '.jj/*' -e notify.socket .
            }

            TMPDIR=$(mktemp -d)
            while true; do
            	if ! ${pkgs.go}/bin/go build -o "$TMPDIR/viewer" ./cmd/viewer; then
            		watch
            	  continue
            	fi
            	"$TMPDIR/viewer" $@ &
            	PID=$!
            	watch
            	kill $PID
            done
          '';
        };
        devShells.default =
          with pkgs;
          mkShell ({
            packages = [
              go
            ];
          });
      }
    );
}
