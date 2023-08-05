package main

//
//import (
//	"errors"
//	"time"
//)
//
//func main() {
//
//}
//
//func keepCheckingSomething() (bool, error) {
//	timeout := time.NewTimer(10 * time.Second)
//	ticker := time.NewTicker(500 * time.Millisecond)
//
//	defer timeout.Stop()
//	defer ticker.Stop()
//
//	// Keep trying until we're timed out or get a result/error
//	for {
//		select {
//		// Got a timeout! fail with a timeout error
//		case <-timeout.C:
//			// maybe, check for one last time
//			ok, err := checkSomething()
//			if !ok {
//				return false, errors.New("timed out")
//			}
//			return ok, err
//		// Got a tick, we should check on checkSomething()
//		case <-ticker.C:
//			ok, err := checkSomething()
//			if err != nil {
//				// We may return, or ignore the error
//				return false, err
//				// checkSomething() done! let's return
//			} else if ok {
//				return true, nil
//			}
//			// checkSomething() isn't done yet, but it didn't fail either, let's try again
//		}
//	}
//}
