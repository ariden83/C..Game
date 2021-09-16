package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

//var actions []string = []string{"LEFT", "RIGHT", "PICK", "PLACE"}
var btxWanted *int
var right string = "right"
var direction *string = &right

func BtxWanted(boxes []int) {
    var i float64 = 0
    for _, v := range boxes{
       i += float64(v)
    }
    total := i / float64(len(boxes)) 
    totalBox := int(math.Round(total))
    if totalBox < 1 {
        totalBox = 1
    }
    btxWanted = &totalBox
}

func Solve(clawPos int, boxes []int, boxInClaw bool) string {
    if btxWanted == nil {
         BtxWanted(boxes)
    }
    // Write your code here
    for i := range boxes {
        if i == clawPos{
            return actions(i, clawPos, boxInClaw, boxes)
        }
    }

    return ""
}

func actions(currentColonne, clawPos int, boxInClaw bool, boxes []int) string {
    if *btxWanted < boxes[clawPos] && !boxInClaw {
        return "PICK"
    }

    needBox := false
    for i, e := range boxes {
        if e < *btxWanted && i < clawPos  && boxInClaw{
            return "left"
        } else if e < *btxWanted && i < clawPos  {
            needBox = true
        } else if needBox && e == *btxWanted && i == clawPos  {
            return "PICK"
        }
    }

    if *btxWanted > boxes[clawPos] && boxInClaw {
        return "PLACE"
    }

    if currentColonne < 1 && *direction == "left" {
        *direction = "right"
    
    } else if currentColonne >= (len(boxes) -1 ) && *direction == "right" {
        *direction = "left"
    }
    return *direction
}

/* Ignore and do not change the code below */
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)
    for {
        var clawPos int
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&clawPos)
        
        var boxInClaw bool
        var _boxInClaw int
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&_boxInClaw)
        boxInClaw = _boxInClaw != 0
        
        var stacks int
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&stacks)
        
        scanner.Scan()
        inputs := strings.Split(scanner.Text(), " ")
        boxes := make([]int, stacks)
        for i := 0; i < stacks; i++ {
            boxes[i],_ = strconv.Atoi(inputs[i])
        }
        stdoutWriter := os.Stdout
        os.Stdout = os.Stderr
        action := Solve(clawPos, boxes, boxInClaw)
        os.Stdout = stdoutWriter
        fmt.Println(action)
    }
}
