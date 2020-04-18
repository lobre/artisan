package artisans

import (
    "os/exec"
)

const NIX_ENV string = "nix-env"

type Nix struct {
    Channel string
    Attr bool

    pkgs []string
}

func NewNix(pkgs ...string) *Nix {
    nix := Nix{
        Channel: "nixpkgs",
        Attr: false,
        pkgs: pkgs,
    }
    return &nix
}

// Add will populate more packages to install.
func (nix *Nix) Add(pkgs ...string) {
    nix.pkgs = append(nix.pkgs, pkgs...)
}

// Craft implements Artisan.
//
// It makes sure the list of packages is installed.
//
// By now, there is no way to know the original attribute path of an installed
// package. Because of this, we have no way to know if the package is already
// installed if we are in Attr mode. For this reason, we always try to install
// without checking if already installed.
//
// See https://github.com/NixOS/nix/pull/3380.
func (n *Nix) Craft() ([]byte, error) {
    return n.install()
}

// Install will execute the nix-env command to install the 
// list of packages.
func (nix *Nix) install() ([]byte, error) {
    attr := []string{"--install"}
    if nix.Attr {
        attr = append(attr, "--attr")
    }

    for _, pkg := range nix.pkgs {
        p := pkg
        if nix.Attr {
            p = nix.Channel + "." + p
        }
        attr = append(attr, p)
    }

    cmd := exec.Command(NIX_ENV, attr...)

    return cmd.CombinedOutput()
}
