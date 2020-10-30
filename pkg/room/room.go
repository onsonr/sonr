package room

import (
	"fmt"
)

// Room Struct ill fix it later
type Room struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
}

// LeavesRemaining Method ill fix it later
func (e Room) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}