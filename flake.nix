{
  description = "Sonr Network Development Environment";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";
    process-compose-flake.url = "github:Platonic-Systems/process-compose-flake";
  };

  outputs = inputs@{ flake-parts, nixpkgs, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        inputs.process-compose-flake.flakeModule
      ];
      systems = import inputs.systems;
      perSystem = { config, self', inputs', pkgs, system, ... }: {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go and related tools
            go
            goreleaser
            sqlc
            templ
            
            # Database
            postgresql
            
            # IPFS
            ipfs
            
            # CLI tools
            gum
            fzf
            gh
            jq
            
            # Task runner
            go-task

            # Development tools
            git
          ];

          shellHook = ''
            echo "Welcome to Sonr Network Development Environment"
          '';
        };

        # Default package remains as a placeholder
        packages.default = pkgs.hello;
      };
    };
}
