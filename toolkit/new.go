package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func newProj(args []string) error {
	// Initialize the spinner (loading thing).
	spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	// Find the name that the project should have. It checks if there is anything after "new" in "malino new".
	// "malino new" = the project name will be the name of the current folder.
	// "malino new test" = the project will have the name of "test".
	lang := "undefined"
	name := "undefined"
	if dir, err := os.Getwd(); err != nil {
		return err
	} else {
		name = strings.Split(dir, "/")[len(strings.Split(dir, "/"))-1] // "name = split by / [len(split by /) - 1]" basically.
	}

	hasLangArg := len(args) == 2
	if hasLangArg {
		if args[1] == "-go" {
			lang = "go"
		} else if args[1] == "-cs" {
			lang = "c#"
		} else {
			return fmt.Errorf("expected language. usage: `malino new -go|-cs`")
		}
	} else {
		return fmt.Errorf("expected language. usage: `malino new -go|-cs`")
	}

	if lang == "go" {
		fmt.Println("  W go.mod")
		spinner.Start()
		if err := execCmd(false, "/usr/bin/go", "mod", "init", name); err != nil { // init the go module
			spinner.Stop()
			return err
		}
		fmt.Println("  W main.go")
		err := os.WriteFile("main.go", []byte(
			"package main\n\n"+
				"import (\n"+
				"	\"github.com/malinoOS/malino/libmalino\"\n"+
				"	\"fmt\"\n"+
				")\n\n"+
				"func main() {\n"+
				"	defer libmalino.ResetTerminalMode()\n"+
				"	fmt.Println(\"malino (project "+name+") booted successfully. Type a line of text to get it echoed back.\")\n"+
				"	for { // Word of advice: Never let this app exit. Always end in an infinite loop or shutdown.\n"+
				"		fmt.Print(\"Input: \")\n"+
				"		input := libmalino.UserLine()\n"+
				"		fmt.Println(\"Text typed: \" + input)\n"+
				"	}\n"+
				"}"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}

		fmt.Println("  W malino.cfg")
		err = os.WriteFile("malino.cfg", []byte(
			"lang c#\n"+
				"verfmt yymmdd\n"+
				"# Required includes\n"+
				"include /lib/x86_64-linux-gnu/libdl.so.2 /lib/x86_64-linux-gnu/libdl.so.2\n"+
				"include /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/libc.so.6\n"+
				"include /lib64/ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2\n"+
				"include /lib/x86_64-linux-gnu/libpthread.so.0 /lib/x86_64-linux-gnu/libpthread.so.0\n"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}

		err = os.WriteFile(".gitignore", []byte("vmlinuz\ninitramfs.cpio.gz\n"+name+".iso"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}

		// make sure we are on the latest version of libmalino
		if err := execCmd(true, "/usr/bin/go", "get", "github.com/malinoOS/malino/libmalino@main"); err != nil {
			spinner.Stop()
			return err
		}
		spinner.Stop()
	} else { /////////////////////////////////////////
		spinner.Start()
		fmt.Println("  W " + name + ".csproj")
		err := os.WriteFile(name+".csproj", []byte(
			"<Project Sdk=\"Microsoft.NET.Sdk\">\n\n"+
				"	<PropertyGroup>\n"+
				"		<OutputType>Exe</OutputType>\n"+
				"		<TargetFramework>net8.0</TargetFramework>\n"+
				"		<ImplictUsings>disable</ImplictUsings>\n"+
				"		<Nullable>enable</Nullable>\n"+
				"	</PropertyGroup>\n\n"+
				"	<ItemGroup>\n"+
				"		<Reference Include=\"/opt/malino/libmalino-cs.dll\">\n"+
				"			<HintPath>/opt/malino/libmalino-cs.dll</HintPath>\n"+
				"		</Reference>\n"+
				"	</ItemGroup>\n"+
				"</Project>\n"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}

		fmt.Println("  W Main.cs")
		err = os.WriteFile("Main.cs", []byte(
			"using libmalino;\n"+
				"using System;\n\n"+
				"namespace "+name+";\n\n"+
				"public class Program {\n"+
				"	public static void Main(string[] args) {\n"+
				"		Console.WriteLine(\"malino (project "+name+") booted successfully. Type a line of text to get it echoed back.\");\n"+
				"		while (true) {\n"+
				"			Console.Write(\"Input: \");\n"+
				"			string input = Console.ReadLine(); // ReadLine is kind of broken right now, wait a bit until malinoIO.UserLine is implemented.\n"+
				"			Console.WriteLine(\"Text typed: \" + input);\n"+
				"		}\n"+
				"	}\n"+
				"}"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}

		fmt.Println("  W malino.cfg")
		err = os.WriteFile("malino.cfg", []byte(
			"lang c#\n"+
				"verfmt yymmdd\n"+
				"# Required includes\n"+
				"include /lib/x86_64-linux-gnu/libdl.so.2 /lib/x86_64-linux-gnu/libdl.so.2\n"+
				"include /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/libc.so.6\n"+
				"include /lib64/ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2\n"+
				"include /lib/x86_64-linux-gnu/libpthread.so.0 /lib/x86_64-linux-gnu/libpthread.so.0\n"+
				"include /opt/malino/libmsb.so /lib/libmsb.so\n"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}

		err = os.WriteFile(".gitignore", []byte("vmlinuz\ninitramfs.cpio.gz\n"+name+".iso\nbin\nobj"), 0777)
		if err != nil {
			spinner.Stop()
			return err
		}
		spinner.Stop()
	}
	return nil
}
