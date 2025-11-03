{
  description = "Devshell for Go + MySQL.";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};
  in {
    devShells.${system}.default = pkgs.mkShell {
      # Add packages here.
      buildInputs = with pkgs; [
        mysql80
        just
        go
        gopls
        gotools
        nodejs
      ];

      # Shell hooks.
      shellHook = ''
        echo "Entering the development environment!"
        mysql --version

        echo 'Just Operations:'
        echo '  "just init" -- initialize MySQL.'
        echo '  "just server" -- start MySQL server.'
        echo '  "just start" -- enter MySQL shell.'
        echo '  "just stop" -- stop MySQL server.'
      '';
    };
  };
}
