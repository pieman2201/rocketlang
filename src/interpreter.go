package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "io/ioutil"
)

var stk = make([]int, 1)
var cr1 = 0
var cr2 = 0
var crc = 0

var cms = make([]string, 0)
var inp string
var ind = 0

var err = true

func input(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    text, _ := reader.ReadString('\n')
    return text[:len(text) - 1]
}

func inform() {
    secStr := " second"
    if ind != 0 {
        secStr = secStr + "s"
    }
    erl := "S#@" + string(37) + "!\nChat disabled for " + strconv.Itoa(ind + 1) + secStr + ".\n"
    if err {
        fmt.Println("\n" + erl)
        os.Exit(1)
    }
}

func push(toPush int) {
    tempStk := make([]int, len(stk) + 1)
    for i, val := range stk {
        tempStk[i + 1] = val
    }
    tempStk[0] = toPush
    stk = tempStk
}

func pop() int {
    toPop := 0
    if len(stk) > 0 {
        toPop = stk[0]
        stk = stk[1:]
    } else {
        inform()
    }
    return toPop
}

func interpret(line string) {
    cur := 0
    if crc == 0 {
        cur = cr1
    } else if crc == 1 {
        cur = cr2
    }

    switch line {
    case "Wow!":
        cur++
    case "Close one!":
        cur--
    case "OMG!":
        crc = (crc + 1) % 2
        if crc == 0 {
            cur = cr1
        } else if crc == 1 {
            cur = cr2
        }
    case "Noooo!":
        push(0)
    case "Centering...":
        cur = pop()
    case "Defending...":
        push(cur)
    case "I got it!":
        fmt.Print(string(byte(cur)))
    case "Take the shot!":
        if len(inp) > 0 {
            cur = int(byte(inp[0]))
            inp = inp[1:]
        }
    case "Nice shot!":
        if cr1 != cr2 { ind++
            for starts := 1; starts > 0; ind++ {
                if cms[ind] == "Nice shot!" {
                    starts++
                } else if cms[ind] == "What a save!" {
                    starts--
                }
            }; ind--
        }
    case "Great pass!":
        if cur <= 0 { ind++
            for starts := 1; starts > 0; ind++ {
                if cms[ind] == "Great pass!" {
                    starts++
                } else if cms[ind] == "Thanks!" {
                    starts--
                }
            }; ind--
        }
    case "Thanks!":
        if cur > 0 { ind--
            for stops := 1; stops > 0; ind-- {
                if cms[ind] == "Thanks!" {
                    stops++
                } else if cms[ind] == "Great pass!" {
                    stops--
                }
            }; ind++
        }
    case "Sorry!":
        fmt.Println(stk, cr1, cr2, crc)
    case "No Problem.":
        err = !err
    case "Whoops...":
        if crc == 0 {
            cr2 = cr1
        } else if crc == 1 {
            cr1 = cr2
        }
    default:
        isBlank := true
        for _, char := range line {
            if char != ' ' {
                isBlank = false
                break
            }
        }
        if !isBlank {
            inform()
        }
    }

    if crc == 0 {
        cr1 = cur
    } else if crc == 1 {
        cr2 = cur
    }
}

func main() {
    path := ""
    if len(os.Args) > 1 {
        path = os.Args[1]
    } else {
        inform()
    }
    barray, _ := ioutil.ReadFile(path)
    cms = strings.Split(string(barray), "\n")
    for _, cm := range cms {
        if cm == "Take the shot!" {
            inp = input("Input: ")
            break
        }
    }
    for ind < len(cms) {
        interpret(cms[ind])
        ind++
    }
}