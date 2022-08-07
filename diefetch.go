package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

const (
	BOLD_IRED   = "\033[1;91m"
	BOLD_IGREEN = "\033[1;92m"
	BOLD_IWHITE = "\033[0;97m"
)

func print(msg string) {
	fmt.Printf("%s", msg)
}

func main() {
	if runtime.GOOS != "linux" {
		print("Only machines which are running GNU/Linux are supported\n")
		return
	}
	shell := strings.Split(os.Getenv("SHELL"), "/")[2]
	de := strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP"))

	wm, err := exec.Command("update-alternatives", "--list", "x-window-manager").Output()
	if err != nil {
		print("Failed to execute \"update-alternatives\"" + "\n" + err.Error() + "\n")
		return
	}
	cpucmd, err := exec.Command("cat", "/proc/cpuinfo").Output()
	if err != nil {
		print("Failed to execute \"grep\"" + "\n" + err.Error() + "\n")
		return
	}
	cpu := strings.TrimSpace(strings.ReplaceAll(strings.Split(strings.Split(string(cpucmd), "model name")[1], "stepping")[0], ":", ""))
	term := os.Getenv("TERM")
	user := os.Getenv("USER")
	lang := strings.ReplaceAll(strings.Split(os.Getenv("LANG"), ".")[0], "_", "-")
	host, err := os.ReadFile("/etc/hostname")
	if err != nil {
		print("Failed to read \"/etc/hostname\" file\n" + err.Error() + "\n")
		return
	}
	os_release, err := os.ReadFile("/etc/os-release")
	if err != nil {
		print("Failed to read \"/etc/os-release\" file\n" + err.Error() + "\n")
		return
	}
	osname := strings.Split(strings.Split(string(os_release), "PRETTY_NAME=")[1], "NAME")[0]
	sysinfo := syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(&sysinfo); err != nil {
		print("Failed to create syscall" + err.Error() + "\n")
		return
	}
	secs := (sysinfo.Uptime) % 60
	min := (sysinfo.Uptime / 60) % 60
	hours := (sysinfo.Uptime / 60 / 60) % 24
	days := (sysinfo.Uptime / 60 / 60) / 24

	print(fmt.Sprintf(`
	%s  _nnnn_
    %s     dGGGGMMb                   %s%s@%s%s
    %s    @p~qp~~qMb                  %sOS:%s %s
    %s    M|%s@%s||%s@%s) M|                  %sUptime:%s %dd, %dh, %dm, %ds
    %s    @,----.JM|                  %sDE:%s %s
    %s   JS^\__/  qKL                 %sWM:%s %s
    %s  dZP        qKRb               %sShell:%s %s
    %s dZP          qKKb              %sTerminal:%s %s
    %sfZP            SMMb             %sLang:%s %s
    %sHZM            MMMM             %sCPU:%s %s
    %sFqM            MMMM
  __%s| ".        |\dS"qML
  %s|    '.       | \Zq
 _%s)      \.___.,|     .'
 %s\____   )MMMMMP|   .'	
	`,
		BOLD_IWHITE,
		BOLD_IWHITE, BOLD_IGREEN, user, strings.Trim(string(host), "\n"), BOLD_IWHITE,
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, strings.ReplaceAll(strings.ReplaceAll(osname, "\n", ""), "\"", ""),
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, days, hours, min, secs,
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, de,
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, strings.Trim(strings.Split(string(wm), "/")[3], "\n"),
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, shell,
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, term,
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, lang,
		BOLD_IWHITE, BOLD_IRED, BOLD_IWHITE, cpu,
		BOLD_IWHITE,
		BOLD_IWHITE,
		BOLD_IWHITE,
		BOLD_IWHITE,
		BOLD_IWHITE,
	) + "\n")
}
