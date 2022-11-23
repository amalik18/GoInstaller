package main

import (
	"fmt"
	"math/rand"
)

var packages = []string{
	"brotli",
	"c-ares",
	"ca-certificates",
	"gcc@11",
	"gdbm",
	"gettext",
	"gmp",
	"go",
	"golangci-lint",
	"gradle",
	"icu4c",
	"isl",
	"jq",
	"libidn2",
	"libmpc",
	"libnghttp2",
	"libunistring",
	"libuv",
	"llvm",
	"lz4",
	"mpdecimal",
	"mpfr",
	"node",
	"oniguruma",
	"openjdk@17",
	"openssl@1.1",
	"python@3.10",
	"readline",
	"six",
	"sqlite",
	"wget",
	"xz",
	"z3",
	"zstd",
	"keepassxc",
	"zeichenorientierte-benutzerschnittstellen",
	"schnurrkit",
	"old-socks-devel",
	"jalapeño",
	"molasses-utils",
	"xkohlrabi",
	"party-gherkin",
	"snow-peas",
	"libyuzu",
}

func GetPackages() []string {
	pkgs := packages
	copy(pkgs, packages)

	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})

	for k := range pkgs {
		pkgs[k] += fmt.Sprintf("-%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10))
	}
	return pkgs
}
