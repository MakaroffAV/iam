// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package amp

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"net/url"
	"sync"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// managerProxies is data structure
// to   describe   the   array   of
// available  proxy  addresses  for
// the           web        crawler
type managerProxies struct {
	pa []*url.URL
	mu sync.Mutex
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (

	// proxies defines the array
	// of    available   proxies
	proxies = &managerProxies{
		mu: sync.Mutex{},
		pa: []*url.URL{},
	}
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// init   is   function    to
// initialize the amp package
func init() {

	u, uErr := url.Parse("http://OCWUgA:xNI5cs533M@45.15.73.241:1050")
	if uErr != nil {
		panic(
			`
			parsing address of the
			proxy   server  failed
			`,
		)
	}

	proxies.pa = append(proxies.pa, u)

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Get  is   function  to  retrieve
// the proxy server's configuration
func Get() *url.URL {

	proxies.mu.Lock()
	defer proxies.mu.Unlock()

	return proxies.pa[0]

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
