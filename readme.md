# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Data Structures &amp; Algorithms for Go
[![](https://img.shields.io/github/v/release/codemodify/dsa-book-go?style=flat-square)](https://github.com/codemodify/dsa-book-go/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/dsa-book-go?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/dsa-book-go?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/dsa-book-go/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/dsa-book-go?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/dsa-book-go?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/dsa-book-go)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/dsa-book-go)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/dsa-book-go?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/dsa-book-go?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/dsa-book-go?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/dsa-book-go?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/dsa-book-go?style=flat-square)

>__The missing DSA crouse for Go__

>__This guide was created with hands on, get to the point approach rather than wasting
your time with theory and UML diagrams available everywhere, which after reading them
uou still don't under understand where to start and how to practically approach this__


# Helpers - Install Go
- Official Download Page
	- https://golang.org/dl

- Shortcuts
	- Linux
		- Arch Linux (Desktop/Server non ARM)
			- `sudo pacman -S go`
		- Debian/Ubuntu (Desktop/Server non ARM)
			- `curl -O https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz`
			- `sudo tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz`
			- `echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc`
			- `source ~/.bashrc`
		- Raspberry Pi
			- `curl -O https://dl.google.com/go/go1.14.3.linux-armv6l.tar.gz`
			- `sudo tar -C /usr/local -xzf go1.14.3.linux-armv6l.tar.gz`
			- `echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc`
			- `source ~/.bashrc`

	- macOS
		- `curl -O https://dl.google.com/go/go1.14.3.darwin-amd64.pkg`
		- `sudo installer -store -pkg go1.14.3.darwin-amd64.pkg -target /`

	- Windows
		- `curl -O https://dl.google.com/go/go1.14.3.windows-amd64.msi`
		- `msiexec /i go1.14.3.windows-amd64.msi /quiet /qn /norestart /log install.log`

# Helpers - Install IDE(s)
- VSCode
	- Linux
		- Arch Linux (Desktop/Server non ARM)
			- `yay -S visual-studio-code-bin`
		- Debian/Ubuntu (Desktop/Server non ARM)
			- `sudo snap install --classic code`
			- or
			- `curl -Lk -o code_1.45.1-1589445302_amd64.deb "https://go.microsoft.com/fwlink/?LinkID=760868"`
			- `sudo apt install ./code_1.45.1-1589445302_amd64.deb`
		- Raspberry Pi
			- `https://packagecloud.io/install/repositories/headmelted/codebuilds/script.deb.sh | sudo bash`
			- `sudo apt-get install code-oss`
		- RHEL /Fedora / CentOS
			- `sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc`
			- `sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo'`
			- `sudo dnf check-update`
			- `sudo dnf install code`
		- openSUSE
			- `sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc`
			- `sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ntype=rpm-md\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/zypp/repos.d/vscode.repo'`
			- `sudo zypper refresh`
			- `sudo zypper install code`
		- Nix OS
			- `nix-env -i vscode`

	- macOS
		- `curl -O https://dl.google.com/go/go1.14.3.darwin-amd64.pkg`
		- `sudo installer -store -pkg go1.14.3.darwin-amd64.pkg -target /`

	- Windows
		- `curl -Lk -o vscode-setup.exe https://aka.ms/win32-x64-user-stable`
		- `vscode-setup.exe /VERYSILENT /MERGETASKS=!runcode`
