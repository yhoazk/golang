package main
import "fmt"

/* By default channels are unbuffered, meaning that they
will only accept sends(chan<-) if there's a corresponding
receive (<-chan) ready to receive the sent value.
Buffered channels accept a liminted number of valued
without a corresponding receiver for those values */


func main(){
	/*Here we make a channel of strings buffering up to 2 values*/
	message := make(chan string, 2)
	message<-"buffered"
	message<-"channel"
	/*Because this channel is buffered, we can send these values into 
	the channel w/o a correspongind concurrent channel receive*/
	/* Later we can receive these values as usual */
	fmt.Println(<-message)
	fmt.Println(<-message)
}
