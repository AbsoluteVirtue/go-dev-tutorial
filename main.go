// Source - https://stackoverflow.com/a/54213577
// Posted by dirkjot
// Retrieved 2026-08-13, License - CC BY-SA 4.0

// recursionController is a data structure with three channels to control our Crawl recursion.
// Tried to use sync.waitGroup in a previous version, but I was unhappy with the mandatory sleep.
// The idea is to have three channels, counting the outstanding calls (children), completed calls
// (done) and results (results).  Once outstanding calls == completed calls we are done (if you are
// sufficiently careful to signal any new children before closing your current one, as you may be the last one).
//
type recursionController struct {
	results  chan string
	children chan int
	done     chan int
}

// instead of instantiating one instance, as we did above, use a more idiomatic Go solution
func NewRecursionController() recursionController {
	// we buffer results to 1000, so we cannot crawl more pages than that.
	return recursionController{make(chan string, 1000), make(chan int), make(chan int)}
}

// recursionController.Add: convenience function to add children to controller (similar to waitGroup)
func (rc recursionController) Add(children int) {
	rc.children <- children
}

// recursionController.Done: convenience function to remove a child from controller (similar to waitGroup)
func (rc recursionController) Done() {
	rc.done <- 1
}

// recursionController.Wait will wait until all children are done
func (rc recursionController) Wait() {
	fmt.Println("Controller waiting...")
	var children, done int
	for {
		select {
		case childrenDelta := <-rc.children:
			children += childrenDelta
			// fmt.Printf("children found %v total %v\n", childrenDelta, children)
		case <-rc.done:
			done += 1
			// fmt.Println("done found", done)
		default:
			if done > 0 && children == done {
				fmt.Printf("Controller exiting, done = %v, children =  %v\n", done, children)
				close(rc.results)
				return
			}
		}
	}
}
