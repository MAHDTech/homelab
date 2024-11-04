{
  pkgs,
  config,
  lib,
  ...
}: let
  packages = with pkgs; [
    hello
  ];

  devPackages = with pkgs; [
    figlet
    go
    go-tools
    golangci-lint
    hello
    nix
    pulumi-bin
    pulumictl
  ];
in {
  name = "homelab";

  env = {
    PROJECT = "homelab";
  };

  cachix = {
    pull = [
      "pre-commit-hooks"
      "salt-labs"
    ];
    push = "salt-labs";
  };

  devenv = {
    warnOnNewVersion = true;
  };

  dotenv = {
    enable = true;
    disableHint = false;
  };

  packages =
    packages
    ++ lib.optionals (!config.container.isBuilding) devPackages;

  enterShell = ''
    figlet -f starwars -w 120 $PROJECT

    hello --greeting="Hello ''${USER:-user}, welcome to the $PROJECT project!"
  '';

  languages = {
    nix = {
      enable = true;
    };

    shell = {
      enable = true;
    };

    go = {
      enable = true;
    };
  };

  difftastic = {
    enable = true;
  };

  git-hooks = {
    excludes = [
      ".cache"
      ".devenv"
      ".direnv"
      "vendor"
    ];
    hooks = {
      actionlint.enable = true;
      beautysh.enable = true;
      check-json.enable = true;
      check-merge-conflicts.enable = true;
      check-shebang-scripts-are-executable.enable = true;
      check-symlinks.enable = true;
      check-yaml.enable = true;
      commitizen.enable = true;
      convco.enable = true;
      deadnix.enable = true;
      dialyzer.enable = true;
      editorconfig-checker.enable = true;
      gofmt.enable = true;
      golangci-lint.enable = true;
      golines.enable = true;
      gotest.enable = true;
      govet.enable = true;
      gptcommit.enable = true;
      markdownlint = {
        enable = true;
        settings = {
          configuration = {
            MD013 = {
              line_length = 180;
            };
          };
        };
      };
      mixed-line-endings.enable = true;
      nixfmt-rfc-style.enable = true;
      pre-commit-hook-ensure-sops.enable = true;
      prettier.enable = true;
      pretty-format-json.enable = true;
      revive.enable = true;
      ripsecrets.enable = true;
      shellcheck.enable = true;
      shfmt.enable = true;
      staticcheck.enable = true;
      statix.enable = true;
      trufflehog.enable = true;
      trim-trailing-whitespace.enable = true;
      typos.enable = true;
      yamllint = {
        enable = true;
        settings = {
          configuration = ''
            extends: relaxed
            rules:
              line-length: disable
              indentation: enable
          '';
        };
      };
    };
  };

  starship = {
    enable = true;
    config = {
      enable = false;
    };
  };

  devcontainer = {
    enable = true;
    settings = {
      customizations = {
        vscode = {
          extensions = [
            "arrterian.nix-env-selector"
            "esbenp.prettier-vscode"
            "github.vscode-github-actions"
            "golang.go"
            "gruntfuggly.todo-tree"
            "johnpapa.vscode-peacock"
            "kamadorueda.alejandra"
            "mkhl.direnv"
            "ms-azuretools.vscode-docker"
            "ms-kubernetes-tools.vscode-kubernetes-tools"
            "ms-vscode.makefile-tools"
            "nhoizey.gremlins"
            "pinage404.nix-extension-pack"
            "redhat.vscode-yaml"
            "streetsidesoftware.code-spell-checker"
            "tekumura.typos-vscode"
            "timonwong.shellcheck"
            "tuxtina.json2yaml"
            "hediet.vscode-drawio"
            "vscodevim.vim"
            "wakatime.vscode-wakatime"
            "yzhang.markdown-all-in-one"
          ];
        };
      };
    };
  };

  enterTest = ''
    echo "Running devenv tests..."
  '';
}
