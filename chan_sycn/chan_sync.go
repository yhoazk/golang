package main
/* We canuse channels to synchronize execution 
	across goroutines. Here is an example of usnig
	blocking receive to wait for a goroutine to finish */
import(
	"fmt"
	"time"
)
/* This is the function we will run in a goroutine
	The done channel will be used to notify another
	goroutines that this function work sis done */
func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	/* Send a value to notify that we are done */
	done <- true
}

func main(){
	/* Start a worker goroutine, giving it the channel to notify */
	done := make(chan bool, 1)
	go worker(done)
	/* Block until we receive a notification from the worker */
	<-done
	/* If we remove the <-done line from this program, it would exit
	before the worker even started */
}
