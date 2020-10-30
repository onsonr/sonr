package transfer

import (
	"fmt"
)

// Transfer object TODO: Ill fix it later
type Transfer struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
}

// LeavesRemaining method  Ill fix it later
func (e Transfer) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}