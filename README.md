Go Whole Program LLVM
==================

Introduction
------------

This project, gowllvm, provides tools for building whole-program (or
whole-library) LLVM bitcode files from an unmodified C or C++
source package. It currently runs on `*nix` platforms such as Linux,
FreeBSD, and Mac OS X. It is a Go port of the WLLVM project, available
at https://github.com/SRI-CSL/whole-program-llvm.

gowllvm provides compiler wrappers that work in two
steps. The wrappers first invoke the compiler as normal. Then, for
each object file, they call a bitcode compiler to produce LLVM
bitcode. The wrappers also store the location of the generated bitcode
file in a dedicated section of the object file.  When object files are
linked together, the contents of the dedicated sections are
concatenated (so we don't lose the locations of any of the constituent
bitcode files). After the build completes, one can use a gowllvm
utility to read the contents of the dedicated section and link all of
the bitcode into a single whole-program bitcode file. This utility
works for both executable and native libraries.

This two-phase build process is necessary to be a drop-in replacement
for gcc or g++ in any build system.  Using the LTO framework in gcc
and the gold linker plugin works in many cases, but fails in the
presence of static libraries in builds.  gowllvm's approach has the
distinct advantage of generating working binaries, in case some part
of a build process requires that.

gowllvm currently works with clang.

Installation
------------

Requirements
=======

You need the Go compiler to compile gowllvm, and the clang/clang++ executables
to use gowllvm. Follow the instructions here to get started:
https://golang.org/doc/install.

As for now, let us name `$GOROOT` your root Go path that you can obtain by
typing `go env GOPATH` in a terminal session -- it is usually `$HOME/go`
by default. It is worth noticing that a standard Go installation will install
the binaries generated for the project under `$GOROOT/bin`. Make sure that you
added the `$GOROOT/bin` directory to your `$PATH` variable.

Build
=======

First, you must checkout the project under the directory `$GOROOT/src`:
```
cd $GOROOT/src
git clone https://github.com/loicgelle/gowllvm
```

To build and install gowllvm on your system, type:
```
make install
```

Usage
-----

gowllvm includes three symlinks to the program's binary: `gowclang` and
`gowclang++`to compile C and C++, and an auxiliary tool `gowextract` for
extracting the bitcode from a build product (object file, executable, library
or archive).

Some useful environment variables are listed here:

 * `GOWLLVM_CC_NAME` can be set if your clang compiler is not called `clang` but
    something like `clang-3.7`. Similarly `GOWLLVM_CXX_NAME` can be used to
    describe what the C++ compiler is called. We also pay attention to the
    environment  variables `GOWLLVM_LINK_NAME` and `GOWLLVM_AR_NAME` in an
    analagous way, since they too get adorned with suffixes in various Linux
    distributions.

 * `GOWLLVM_TOOLS_PATH` can be set to the absolute path to the folder that
   contains the compiler and other LLVM tools such as `llvm-link` to be used.
   This prevents searching for the compiler in your PATH environment variable.
   This can be useful if you have different versions of clang on your system
   and you want to easily switch compilers without tinkering with your PATH
   variable.
   Example `GOWLLVM_TOOLS_PATH=/home/user/llvm_and_clang/Debug+Asserts/bin`.

* `GOWLLVM_CONFIGURE_ONLY` can be set to anything. If it is set, `gowclang`
   and `gowclang++` behave like a normal C or C++ compiler. They do not
   produce bitcode. Setting `GOWLLVM_CONFIGURE_ONLY` may prevent configuration
   errors caused by the unexpected production of hidden bitcode files. It is
   sometimes required when configuring a build.

Preserving bitcode files in a store
--------------------------------

Sometimes it can be useful to preserve the bitcode files produced in a
build, either to prevent deletion or to retrieve them later. If the
environment variable `GOWLLVM_BC_STORE` is set to the absolute path of
an existing directory, then gowllvm will copy the produced bitcode files
into that directory. The name of a copied bitcode file is the hash of the path
to the original bitcode file. For convenience, when using both the manifest
feature of `gowextract` and the store, the manifest will contain both the
original path, and the store path.

Building a bitcode module with clang
------------------------------------

```
tar xf pkg-config-0.26.tar.gz
cd pkg-config-0.26
CC=gowclang ./configure
make
```

This should produce the executable `pkg-config`. To extract the bitcode:
```
gowextract pkg-config
```

which will produce the bitcode module `pkg-config.bc`.


Building bitcode archive
------------------------

```
tar -xvf bullet-2.81-rev2613.tgz
mkdir bullet-bin
cd bullet-bin
CC=gowclang CXX=gowclang++ cmake ../bullet-2.81-rev2613/
make

# Produces src/LinearMath/libLinearMath.bca
gowextract src/LinearMath/libLinearMath.a
```

Note that by default extracting bitcode from an archive produces an archive of
bitcode. You can also extract the bitcode directly into a module:
```
gowextract -b src/LinearMath/libLinearMath.a
```
produces `src/LinearMath/libLinearMath.a.bc`.


Configuring without building bitcode
------------------------------------

Sometimes it is necessary to disable the production of bitcode. Typically this
is during configuration, where the production of unexpected files can confuse
the configure script. For this we have a flag `GOWLLVM_CONFIGURE_ONLY` which
can be used as follows:
```
GOWLLVM_CONFIGURE_ONLY=1 CC=gowclang ./configure
CC=gowclang make
```


Building a bitcode archive then extracting the bitcode
------------------------------------------------------

```
tar xvfz jansson-2.7.tar.gz
cd jansson-2.7
CC=gowclang ./configure
make
mkdir bitcode
cp src/.libs/libjansson.a bitcode
cd bitcode
gowextract libjansson.a
llvm-ar x libjansson.bca
ls -la
```
