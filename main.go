package main

import (
    "os"
    "log"
    "github.com/lobre/artisan/artisans"
)

type Artisan interface {
    Craft() ([]byte, error)
}

// Play is a list of tasks to do.
// Each task is executed by an artisan.
type Play []Artisan

func (p *Play) Add(art Artisan) {
    *p = append(*p, art)
}

func (p *Play) Perform(logger *log.Logger) error {
    for _, art := range *p {
        out, err := art.Craft()
        if err != nil {
            return err
        }
        logger.Print(string(out))
    }
    return nil
}

func main() {
    logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

    play := new(Play)

    // Define your play
    nix := artisans.NewNix("sl")
    nix.Channel = "nixos"
    nix.Attr = true

    play.Add(nix)

    if err := play.Perform(logger); err != nil {
        panic(err)
    }
}

