# knights

# Usage

Help:
```shell
knights --help
```

Default usage: 5 knights, 5 witches, select next in circle in forward direction:
```shell
knights
```

10 knights, 2 witches, select next randomly:
```shell
knights -knights 10 -witches 2 -random
```

# Part 6

## How to add graphics? Multiplatform?

Use library or frameworks for drawing graphics. It could be primitive because our game has simple logic and doesn't interact with user input. This library or framework should support cross-platform development.
Possible solutions:
- python, and graphics package turtle
- cross-browser HTML5 game frameworks
- game engines, I googled GDevelop, it seems quite easy for beginners