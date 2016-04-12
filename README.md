# rocketlang
A programming language based on Rocket League quick chat.
## Why?
Why not?
## How to
1. Download the source
2. Download and install Go
3. `$ go build interpreter.go`
4. Write your rocketlang program (*.rekt)
5. `$ interpreter [path to file]`

## Overview
rocketlang consists of 2 variables and a stack. These variables can be pushed to the stack and the stack can pop to them. Only one of the variables is modified at a time. To modify the other variable, one must switch which variable is in use.
### Statements

`Wow!` increments the current variable by 1

`Close one!` decrements the current variable by 1

`OMG!` switches between the current variables

`Noooo!` pushes 0 to the stack

`Centering...` pops from the stack to the current variable

`Defending...` pushes from the current variable to the stack

`I got it!` outputs the ASCII value of the current variable

`Take the shot!` reads one byte from input and sets the current variable to it

`Nice shot!` begins an if statement which runs if the 2 variables are equal

`What a save!` ends an if statement

`Great pass!` begins a loop which runs if the current variable is `>` 0 

`Thanks!` goes back to the top of the loop if the current variable is `>` 0

`Whoops...` sets the variable not in use to the current variable (equivalent to `Defending...\nOMG!\nCentering...\nOMG!`)

`Sorry!` outputs debug information between two newlines (stack, var0, var1, current)

`No Problem.` toggles error recognition

`$#@%!` is outputted along with a line number when you messed up if error recognition is on

## hello world
See `scripts/hello.rekt`
