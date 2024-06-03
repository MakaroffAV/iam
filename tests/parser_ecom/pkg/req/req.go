// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package req

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Get  is  function to execute
// http request with GET method
func Get(url string) ([]byte, error) {

	fmt.Println(url)

	var (

		// c defines  the  client config
		c = &http.Client{
			Timeout: time.Second * 5,
		}
	)

	// setting  up  the  request  config
	q, qErr := http.NewRequest("GET", url, nil)
	if qErr != nil {
		return nil, qErr
	}

	q.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	q.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	q.Header.Set("Sec-Fetch-Site", "none")
	q.Header.Set("Cookie", "t3_sid_7588506=s1.1650499959.1717404878565.1717412318209.2.130; rl_session=RudderEncrypt%3AU2FsdGVkX19aNdHEnfb95bsDQOEuH%2Bu%2Fs1bfbBolOHGIEGZFmno0QTzRG9aMFLz6qlgykG40fBnxQKEf1GmIB0V7F87pYuJr49sILVrrDR81T8TVelexD11cCwrByT5T8%2FAnWX69zvaVyk8tqwj4vQ%3D%3D; rl_trait=RudderEncrypt%3AU2FsdGVkX19IJeAxQxI53tPra7EXW0JgKTDklq8gX9n79Va1BFVMP6FtCWybfAu6ZxHkA1mDg2ZP7Bu9nFjWsC5RiUsocCj1JhPmuwskGVUXIQ7jrQjuLc13aUwtzp52iGu7oSlf5QiUC9xB0DuZhQO11r7e7k1l92itYPbFSxjgcuHWiw0A5ntyzPuv39O1OohPeOxM3OSzF5p8Zlpusw%2FHpDEVS5ouP2p9Svn0oUXXrZO2s1Om3wt6lCXaA9JdFt6zssqOqf5R62wL6Vf4E0T%2FBEfCBxD08%2BYztRDD4A%2BjzSsXSus9tR9RFA1oJM7iQfsONKrWxy9b4lZ4UMkD7eoQUQnRSl1%2FFD3qBi0F9o7YZqrA49In30qyU9sWT5wvdd6bjvA8FIbFEg4JgsFbmw%3D%3D; rl_user_id=RudderEncrypt%3AU2FsdGVkX19Vtx7QtNfBJQO%2B7Mjw9xBYaqU8yj8wc7M%3D; _Instamart_session=NU5KcXZSZnhyUmxZbTNrNlQxeHFBeGZsQjJzbFZUbzd2dFlkSHgxSllHRWUxY3FXTll3b0swWk5kR21YeCt4Z3U5VDRvUUhXK21OQXcxQXRER2VobFFNeUR1ZDd6Z0Nla3FWMFM1UkEwT3hpNkZvamFCQndCZVgvSlRGUkVSSmZ2Zk41Y0g0TzR1SFhrS1Q0ZlVNYmxRa2R2bndsdVNSVCtnekhVT0t5WVpHOFBqcUluUGI0aTRHaU1wVXVTTGVrQzMwc1U3eG9CckxlTmk2akIwMDd6dz09LS13ZDN1c2RsajQ1Ky9nSWxVMlpSSTBRPT0%3D--759bec1205f2b4c3de2ac7377f0345d1a1a3f944; _pk_id.6.3ec0=a8d4f993ed221a58.1717399378.; _pk_ses.6.3ec0=1; tmr_detect=0%7C1717411207443; _ym_visorc=b; rl_anonymous_id=RudderEncrypt%3AU2FsdGVkX183SJFhj7pVB3g0boJ2Ta0zHhz1malMeAOAN5hQmswJ66gQpDQOYPrkN5ws6GfNbRhHzRKWi7nMEQ%3D%3D; rl_group_id=RudderEncrypt%3AU2FsdGVkX188nFKQy9GjG1deNuNN9abiacG0mlBe1Is%3D; rl_group_trait=RudderEncrypt%3AU2FsdGVkX1%2FTmsALa0EuEqdi6H5xilJyQ8%2BrfUAISRA%3D; sessionId=17174111665381016991; ssrMedia={%22windowWidth%22:1320%2C%22primaryInput%22:%22mouse%22}; adrcid=AkxgePBgZcDqfS3vesOyjYg; adtech_uid=aa268066-3ffa-4da0-b1a2-1c64478b5f08%3Asbermarket.ru; rcuid=665d6f53c20144eb574b239d; rrpvid=247350086293198; tmr_lvid=588948006ab8880ac544a4ef39c0bc34; tmr_lvidTS=1717399379474; top100_id=t1.7588506.1628040553.1717399378914; OnboardingState={%22state%22:{%22viewedOnboardingKeys%22:[%22pickup_map_filters%22]}%2C%22version%22:0}; _808db7ba1248=%5B%7B%22source%22%3A%22%28direct%29%22%2C%22medium%22%3A%22%28none%29%22%2C%22cookie_changed_at%22%3A1717405858%7D%2C%7B%22source%22%3A%22sbermarket.ru%22%2C%22medium%22%3A%22referral%22%2C%22cookie_changed_at%22%3A1717411203%7D%5D; spsc=1717411167045_6864221ed748d9365e2d1b53d7a10209_0d4e0eb9c92181c656f0cb950dff1b59a3402473b747382c8fc1bc9ef31a9efa; cookies_consented=yes; identified_address=true; directCrm-session=%7B%22deviceGuid%22%3A%221549a514-3703-4107-bc5a-40efc02616c1%22%7D; mindboxDeviceUUID=1549a514-3703-4107-bc5a-40efc02616c1; rl_page_init_referrer=RudderEncrypt%3AU2FsdGVkX1%2FT76JaseayAS8KE23leJZE%2F0s6UizkTsM%3D; rl_page_init_referring_domain=RudderEncrypt%3AU2FsdGVkX1%2FKyDRmkXyNFQGeQKYIwv%2BXVw0Ni7szvtw%3D; acs_3=%7B%22hash%22%3A%223c8f85edb06b1f745fbd%22%2C%22nextSyncTime%22%3A1717485780935%2C%22syncLog%22%3A%7B%22224%22%3A1717399380935%2C%221228%22%3A1717399380935%2C%221230%22%3A1717399380935%7D%7D; adrdel=1717399380791; _ym_d=1717399379; _ym_isad=2; _ym_uid=1717399379594961482; uxs_uid=1b16cae0-217a-11ef-8b50-51ff95298359; domain_sid=SuJQ4RzJQwGaO9hqzwrWC%3A1717399379730; _sa=SA1.2c212c19-3ef3-4f69-b539-3cedf080eded.1717399378; iap.uid=05b2b1aa9a3a408f84fbdee5687206d7; city_info=%7B%22slug%22%3A%22zaraysk%22%2C%22name%22%3A%22%D0%97%D0%B0%D1%80%D0%B0%D0%B9%D1%81%D0%BA%22%2C%22lat%22%3A54.7695%2C%22lon%22%3A38.8717%7D; external_analytics_anonymous_id=8e1e18f6-0d25-4f6b-8a1d-88079582efab; flocktory-uuid=feb95521-0470-4ebc-a018-635968a580e9-4; spid=1717399377039_673b7f055bbf5b6770d35c4b13564098_eionvo8ig5w8m43a")
	q.Header.Set("Accept-Encoding", "gzip, deflate, br")
	q.Header.Set("Sec-Fetch-Mode", "navigate")
	q.Header.Set("Host", "sbermarket.ru")
	q.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	q.Header.Set("Accept-Language", "ru")
	q.Header.Set("Sec-Fetch-Dest", "document")
	q.Header.Set("Connection", "keep-alive")

	// retrieve the parsing page content
	r, rErr := c.Do(q)
	if rErr != nil {
		return nil, rErr
	}

	return io.ReadAll(r.Body)

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
