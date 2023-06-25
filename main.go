package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	CS_VALID                  = 0x00000001
	CS_ADHOC                  = 0x00000002
	CS_GET_TASK_ALLOW         = 0x00000004
	CS_INSTALLER              = 0x00000008
	CS_FORCED_LV              = 0x00000010
	CS_INVALID_ALLOWED        = 0x00000020
	CS_HARD                   = 0x00000100
	CS_KILL                   = 0x00000200
	CS_CHECK_EXPIRATION       = 0x00000400
	CS_RESTRICT               = 0x00000800
	CS_ENFORCEMENT            = 0x00001000
	CS_REQUIRE_LV             = 0x00002000
	CS_ENTITLEMENTS_VALIDATED = 0x00004000
	CS_NVRAM_UNRESTRICTED     = 0x00008000
	CS_RUNTIME                = 0x00010000
	CS_EXEC_SET_HARD          = 0x00100000
	CS_EXEC_SET_KILL          = 0x00200000
	CS_EXEC_SET_ENFORCEMENT   = 0x00400000
	CS_EXEC_INHERIT_SIP       = 0x00800000
	CS_KILLED                 = 0x01000000
	CS_DYLD_PLATFORM          = 0x02000000
	CS_PLATFORM_BINARY        = 0x04000000
	CS_PLATFORM_PATH          = 0x08000000
	CS_DEBUGGED               = 0x10000000
	CS_SIGNED                 = 0x20000000
	CS_DEV_CODE               = 0x40000000
	CS_DATAVAULT_CONTROLLER   = 0x80000000
)

func main() {
	help := flag.Bool("help", false, `NAME
    CodeSignCheck - CLI tool to check enabled code signing flags

SYNOPSIS
    CodeSignCheck [flags] <CodeSigningDecimal>

DESCRIPTION
    CodeSignCheck takes a decimal number as input and shows which code signing flags are enabled based on the input. 

FLAGS
    -help
        Show this help page.

EXAMPLES
    CodeSignCheck 123456789
        This command will show which code signing flags are enabled for the decimal number 123456789.

NOTES
    Each code signing flag is represented by a hexadecimal number, and each decimal input can correspond to multiple flags.`)

	var codeSigningDecimal int

	if len(os.Args) > 1 {
		parsedInt, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid input. Please provide a decimal number as an argument.")
			os.Exit(1)
		}
		codeSigningDecimal = parsedInt
	} else {
		fmt.Println("No input provided. Please provide a decimal number as an argument.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flags := map[string]int{
		"CS_VALID":                  CS_VALID,
		"CS_ADHOC":                  CS_ADHOC,
		"CS_GET_TASK_ALLOW":         CS_GET_TASK_ALLOW,
		"CS_INSTALLER":              CS_INSTALLER,
		"CS_FORCED_LV":              CS_FORCED_LV,
		"CS_INVALID_ALLOWED":        CS_INVALID_ALLOWED,
		"CS_HARD":                   CS_HARD,
		"CS_KILL":                   CS_KILL,
		"CS_CHECK_EXPIRATION":       CS_CHECK_EXPIRATION,
		"CS_RESTRICT":               CS_RESTRICT,
		"CS_ENFORCEMENT":            CS_ENFORCEMENT,
		"CS_REQUIRE_LV":             CS_REQUIRE_LV,
		"CS_ENTITLEMENTS_VALIDATED": CS_ENTITLEMENTS_VALIDATED,
		"CS_NVRAM_UNRESTRICTED":     CS_NVRAM_UNRESTRICTED,
		"CS_RUNTIME":                CS_RUNTIME,
		"CS_EXEC_SET_HARD":          CS_EXEC_SET_HARD,
		"CS_EXEC_SET_KILL":          CS_EXEC_SET_KILL,
		"CS_EXEC_SET_ENFORCEMENT":   CS_EXEC_SET_ENFORCEMENT,
		"CS_EXEC_INHERIT_SIP":       CS_EXEC_INHERIT_SIP,
		"CS_KILLED":                 CS_KILLED,
		"CS_DYLD_PLATFORM":          CS_DYLD_PLATFORM,
		"CS_PLATFORM_BINARY":        CS_PLATFORM_BINARY,
		"CS_PLATFORM_PATH":          CS_PLATFORM_PATH,
		"CS_DEBUGGED":               CS_DEBUGGED,
		"CS_SIGNED":                 CS_SIGNED,
		"CS_DEV_CODE":               CS_DEV_CODE,
		"CS_DATAVAULT_CONTROLLER":   CS_DATAVAULT_CONTROLLER,
	}

	fmt.Printf("%-25s%-10s\n", "Flag", "Is Enabled")
	for flag, value := range flags {
		fmt.Printf("%-25s%-10v\n", flag, (codeSigningDecimal & value) > 0)
	}
}