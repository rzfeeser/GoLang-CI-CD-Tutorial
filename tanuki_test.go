package main

import (
  "strings"
  "testing"
)

func TestGetTanuki(t *testing.T) {
  //Simple test, check the included color codes and a typical string.
  tables := []string{ "\033[38;5;208m", "-ooooooooo-++++++++++++++-ooooooooo-" }

  res := GetTanuki(true)
  for _, table := range(tables) {
    if strings.Contains(res, table) != true {
      t.Errorf("GetTanuki failed at checking for '%s'.", table)
    }
  }
}
