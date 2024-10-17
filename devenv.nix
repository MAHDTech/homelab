{
  pkgs,
  config,
  lib,
  ...
}: let
  packages = with pkgs; [
    bashInteractive
    hello
  ];

  devPackages = with pkgs; [
    age
    direnv
    figlet
    git
    glibc
    google-cloud-sdk
    jq
    kube-linter
    kubectl
    kubernetes-helm
    openssl
    openssh
    pulumi-bin
    ssh-to-age
    ssh-to-pgp
    sshuttle
    vault
    yq-go
  ];
in {
  name = "home-lab";

  env = {
    PROJECT = "home-lab";
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

    echo Hello ''${USER:-user}, welcome to the $PROJECT project
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

    rust = {
      enable = true;
    };

    python = {
      enable = true;
      venv = {
        enable = true;
      };
      poetry = {
        enable = true;
        activate = {
          enable = true;
        };
        install = {
          enable = true;
          allExtras = true;
          verbosity = "little";
        };
      };
    };
  };

  difftastic = {
    enable = true;
  };

  pre-commit = {
    default_stages = [
      "pre-commit"
    ];

    excludes = [];

    hooks = {
      actionlint = {
        enable = true;
      };

      alejandra = {
        enable = true;
        fail_fast = true;
        excludes = [
          "vendor/"
        ];
      };

      check-json = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      check-shebang-scripts-are-executable = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      check-symlinks = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      check-yaml = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      convco = {
        enable = true;
      };

      cspell = {
        enable = false;
        excludes = [
          "vendor/"
        ];
      };

      deadnix = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      dialyzer = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      editorconfig-checker = {
        enable = true;
      };

      hadolint = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      markdownlint = {
        enable = true;
        settings = {
          configuration = {
            MD013 = {
              line_length = 180;
            };
          };
        };
        excludes = [
          "vendor/"
        ];
      };

      nil = {
        enable = false;
        excludes = [
          "vendor/"
        ];
      };

      pre-commit-hook-ensure-sops = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      prettier = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      pretty-format-json = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      ripsecrets = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      shellcheck = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      shfmt = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      staticcheck = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      trim-trailing-whitespace = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

      typos = {
        enable = true;
        excludes = [
          "vendor/"
        ];
      };

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
        excludes = [
          "vendor/"
        ];
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
