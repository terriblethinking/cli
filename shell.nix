{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go_latest
    gopls
    gotools
    go-tools
    gomodifytags
    impl
  ];

  shellHook = ''
    export OLLAMA_API_KEY=""
  '';
}
