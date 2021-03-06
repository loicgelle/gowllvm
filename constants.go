package main

const(
    // Environment variables
    CONFIGURE_ONLY = "GOWLLVM_CONFIGURE_ONLY"
    TOOLS_PATH = "GOWLLVM_TOOLS_PATH"
    C_COMPILER_NAME = "GOWLLVM_CC_NAME"
    CXX_COMPILER_NAME = "GOWLLVM_CXX_NAME"
    LINKER_NAME = "GOWLLVM_LINK_NAME"
    AR_NAME = "GOWLLVM_AR_NAME"
    BC_STORE_PATH = "GOWLLVM_BC_STORE"

    // Gowllvm functioning
    ELF_SECTION_NAME = ".llvm_bc"
    DARWIN_SEGMENT_NAME = "__WLLVM"
    DARWIN_SECTION_NAME = "__llvm_bc"

    // File types
    FT_UNDEFINED = 0
    FT_ELF_EXECUTABLE = 1
    FT_ELF_OBJECT = 2
    FT_ELF_SHARED = 3
    FT_MACH_EXECUTABLE = 4
    FT_MACH_OBJECT = 5
    FT_MACH_SHARED = 6
    FT_ARCHIVE = 7
)
