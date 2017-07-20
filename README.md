# minesweeper-go
Play Minesweeper in your Terminal (made with TermLoop)

#### About

This was a fun weekend project inspired by an interview question.
After implementing the logic of generating a Minesweeper grid, I decided to just make it into a full-blown terminal game using
[@JoeOtter](https://github.com/JoelOtter)'s excellent [termloop](https://github.com/JoelOtter/termloop) library.

![](https://github.com/ryanbaer/minesweeper-go/blob/master/images/preview.gif?raw=true)




### Usage
```
$ go install github.com/ryanbaer/minesweeper-go

$ minesweeper-go -help
	Usage: minesweeper <width> <height> <# of mines>
	Default: minesweeper 20 10 10

$ minesweeper-go
```

**Support**
This project was developed on macOS, and has been tested on Windows as well.

### Roadmap

- [ ] Possibly make squares bigger (maybe by scaling to the dimensions of the screen)
- [ ] Any bug fixes
- [ ] Clean up code
  - Could use less config & more context for passing data through
    - Comment all methods
    - Review what really needs to be public / private
- [ ] Remove win & lose levels in favor of simple "Press [enter] to play again" on main level
- [x] Fix ugly appearance from unrendered unicode on Windows command prompt
- [x] Investigate high CPU on macOS (Thanks [@mrcrilly](https://github.com/mrcrilly) for pointing out the very high default FPS in TermLoop)

### Troubleshooting
Feel free to open an issue if you run into any issues


### License

[MIT License](https://github.com/ryanbaer/minesweeper-go/blob/master/LICENSE/)
