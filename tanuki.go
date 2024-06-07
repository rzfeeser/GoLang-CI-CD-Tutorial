package main

import (
  "fmt"
)

func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

var (
  purple = Color("\033[38;5;129m%s\033[0m");
  orange = Color("\033[38;5;208m%s\033[0m");
  yellow = Color("\033[38;5;220m%s\033[0m");
  red = Color("\033[38;5;196m%s\033[0m");
  green = Color("\033[38;5;76m%s\033[0m");
  blue = Color("\033[38;5;27m%s\033[0m");
)

func GetTanuki(mode bool) string {
  var out string;

  //Borrowed from https://gitlab.com/gitlab-org/gitlab-foss/-/raw/4086e41fbbcc9d3fe21adf2fb404d778b699d9d7/app/assets/javascripts/console_swag.js
  //Note: Avoid multi line strings in Golang, escaping backticks is a pain.
  out += purple("           +                        +") + "\n";
  out += purple("          :s:                      :s:") + "\n";
  out += purple("         .oso'                    'oso.") + "\n";
  out += purple("         +sss+                    +sss+") + "\n";
  out += purple("        :sssss-                  -sssss:") + "\n";
  out += purple("       'ossssso.                .ossssso'") + "\n";
  out += purple("       +sssssss+                +sssssss+") + "\n";
  out += purple("      -ooooooooo-++++++++++++++-ooooooooo-") + "\n";
  out += red("     ':/") + orange("+++++++++") + yellow("osssssssssssso") + green("+++++++++") + blue("/:'") + "\n";
  out += red("     -///") + orange("+++++++++") + yellow("cssssssssssss") + green("+++++++++") + blue("///-") + "\n";
  out += red("    .//////") + orange("+++++++") + yellow("cosssssssssso") + green("+++++++") + blue("//////.") + "\n";
  out += red("    :///////") + orange("+++++++") + yellow("osssssssso") + green("+++++++") + blue("///////:") + "\n";
  out += red("     .:///////") + orange("++++++") + yellow("ssssssss") + green("++++++") + blue("///////:.'") + "\n";
  out += red("       '-://///") + orange("+++++") + yellow("osssssso") + green("+++++") + blue("/////:-'") + "\n";
  out += red("          '-:////") + orange("++++") + yellow("osssso") + green("++++") + blue("////:-'") + "\n";
  out += red("             .-:///") + orange("++") + yellow("osssso") + green("++") + blue("///:-.") + "\n";
  out += red("               '.://") + orange("++") + yellow("cosso") + green("++") + blue("//:.'") + "\n";
  out += red("                  '-:/") + orange("+") + yellow("coo") + green("+") + blue("/:-'") + "\n";
  out += red("                     '-") + yellow("++") + blue("-'") + "\n";
  out += "\n\n" + red("Tanuki has been summoned...") + "\n";

  return out;
}
