# Go Binaries exportieren
Da in Go alle Abhängigkeiten mit kompiliert werden, benötigt jedes Betriebssystem und jede Plattform eine eigene Binary. Mit go run oder build wird lediglich die Binary für das eigene Betriebssystem / Architektur kompiliert. Um jetzt allerdings nicht tausende Geräte besorgen zu müssen, um Go kompilieren zu können, können Umgebungsvariablen gesetzt werden, um die Binary zu exportieren.

| $GOOS (Betriebssystem) | $GOARCH (Architektur) |
|-----------|-----------|
| aix       | ppc64     |
| android   | 386       |
| android   | amd64     |
| android   | arm       |
| android   | arm64     |
| darwin    | amd64     |
| darwin    | arm64     |
| dragonfly | amd64     |
| freebsd   | 386       |
| freebsd   | amd64     |
| freebsd   | arm       |
| illumos   | amd64     |
| linux     | 386       |
| linux     | amd64     |
| linux     | arm       |
| linux     | arm64     |
| linux     | ppc64     |
| linux     | ppc64le   |
| linux     | mips      |
| linux     | mipsle    |
| linux     | mips64    |
| linux     | mips64le  |
| linux     | riscv64   |
| linux     | s390x     |
| netbsd    | 386       |
| netbsd    | amd64     |
| netbsd    | arm       |
| openbsd   | 386       |
| openbsd   | amd64     |
| openbsd   | arm       |
| openbsd   | arm64     |
| plan9     | 386       |
| plan9     | amd64     |
| plan9     | arm       |
| solaris   | amd64     |
| windows   | 386       |
| windows   | amd64     |
