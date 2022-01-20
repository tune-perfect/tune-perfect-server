# Tune Perfect Song Server

### Cloning the project
```bash
git clone https://github.com/tune-perfect/tune-perfect-server.git
```
### Add songs
1. Find or create songs that fit the UltraStar song format.
1. Copy the files to the `songs/` folder.

### Build and run the project
```bash
go build
```
Afterwards run the generated executable.
The server should now be live, running on port 3000 (as of now hard-coded).

### Technologies
The server is written in GoLang.
It uses [gofiber](https://github.com/gofiber/fiber) as web framework.
