{
  pkgs,
  ...
}:
let
  pname = "nancy";
  version = "1.0.46";

  srcNancy = pkgs.fetchFromGitHub {
    owner = "sonatype-nexus-community";
    repo = "nancy";
    rev = "v${version}";
    sha256 = "sha256-116j4e8zRsuE6s0XPhbaPJLhWCa+uRNOT0BEW3IXuwM=";
  };
in

pkgs.buildGoModule rec {
  inherit pname;
  inherit version;

  src = srcNancy;

  vendorHash = "sha256-+8U38Ia6mehbcVE2k4D+0h6TyN40Ksufgs8aMpqUBXw=";

  nativeBuildInputs = with pkgs; [ musl ];

  ldflags = [
    "-s"
    "-w"
    "-X github.com/sonatype-nexus-community/nancy/buildversion.BuildVersion=${version}"
    "-extldflags '-static -L${pkgs.musl}/lib'"
  ];

  # Set build environment variables
  preBuildPhases = [
    "setupEnvironment"
  ];

  setupEnvironment = ''
    export CGO_ENABLED=0
  '';

  buildPhase = ''
    runHook preBuild
    go \
      build \
        -v \
        -ldflags="${builtins.toString ldflags}" \
        -o "$GOPATH/bin/nancy" \
        .
    runHook postBuild
  '';

  # Override the install phase to put binaries in platform-specific directories
  installPhase = ''
    runHook preInstall
    mkdir -p $out/bin/
    mv $GOPATH/bin/nancy $out/bin/nancy
    runHook postInstall
  '';

  meta = with pkgs.lib; {
    description = "A tool to check for vulnerabilities in your Golang dependencies";
    homepage = "https://github.com/sonatype-nexus-community/nancy";
    license = licenses.asl20;
    maintainers = with maintainers; [ ];
  };
}
