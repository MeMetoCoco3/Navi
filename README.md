# Navi 
## System for managing favorite directories and navigate your file system from your own terminal.

### What is it?
Navi is a easy to use CLI tool to navigate throu your file system, with clear and beautiful visuals and capability to add favourite paths for 
faster navigation. It was made in Go using the *Charm* tools and the CobraCLI, all wraped in a bash script.

![Screenshot from 2024-10-22 00-31-22](https://github.com/user-attachments/assets/4b5c4608-1b61-405b-a6c8-306de90b71ac) 
|:--:| 
| *Navi gator* |

![Screenshot from 2024-10-22 00-29-59](https://github.com/user-attachments/assets/546a5c9e-5ea4-4e46-ab69-0707b81fee05)
|:--:| 
| *Navi fv * |

### Why?
I was so tired of opening the terminal and having to go file by file to every small project or exercice that I wanted to do at that moment, that I had to build something.

### Usage
Navi accepts the next commands: 
- gator: Opens TUI that allows you to navigate through your file system. Press 'a' to add/remove current path from favorites, 'g' to cd to current directory.
- fv: Shows favorite list, allows to remove from favorites with 'a' and cd to selected with 'g'.
- add: add current directory to favorite list.
- rm: remove current directory from favorite list.
- tree: prints current files and folders in the terminal. 


### Install
1. Clone the repo in your /home directory.q
2. Open your *.bashrc* file, write  this var, and add it to your PATH, so navi.sh can be called anywhere.
```shell
export NAVIPATH="$HOME/navi"
PATH=$PATH:$NAVIPATH
```
3. In *.bashrc*, also, write this function.
```shell
navi() {
  source navi.sh
}
```
4. Done!
