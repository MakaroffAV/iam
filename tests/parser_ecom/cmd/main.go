package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sbermarket.ru/api/v3/stores/8391/categories", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", "spid=1717399547084_6db01728f7b81d53155ce91451252e02_1rjqvix11b4cup8t; external_analytics_anonymous_id=8b9a3abc-fac2-43ca-b476-0c587d9f2559; _pk_id.6.3ec0=73360a5c71f5d420.1717399548.; _sa=SA1.19658e2f-7f09-4677-ab05-3cb751bb3b59.1717399548; iap.uid=db0821ccb65b401c85036861e5eff9e7; adtech_uid=d7533642-dae0-401c-abfd-b2974f929067%3Asbermarket.ru; top100_id=t1.7588506.293543324.1717399548156; rrpvid=825863573401522; _ym_uid=17173995483805385; _ym_d=1717399548; rcuid=64a7f544b2113a5f4ab5e34b; tmr_lvid=98ecec3906095423283cd641eb82b199; tmr_lvidTS=1717399548428; flocktory-uuid=002cd231-b88c-41b2-a96e-1c6c5891b50c-4; _ym_isad=1; uxs_uid=7fde0c40-217a-11ef-9cd0-bdeeed7cf966; rl_page_init_referrer=RudderEncrypt%3AU2FsdGVkX1%2B4YxaNxFGvqeEyuEGUJjkdG4lv3yzJYj4%3D; rl_page_init_referring_domain=RudderEncrypt%3AU2FsdGVkX19lysWCxxykUWrvGeEMAq1FdOavIYSG6vk%3D; identified_address=true; domain_sid=u6UDo2Q06-GSyGoG_N2TH%3A1717399734523; mindboxDeviceUUID=076dec3e-297d-4e99-ae0a-39862ae2cd8d; directCrm-session=%7B%22deviceGuid%22%3A%22076dec3e-297d-4e99-ae0a-39862ae2cd8d%22%7D; city_info=%7B%22slug%22%3A%22zaraysk%22%2C%22name%22%3A%22%D0%97%D0%B0%D1%80%D0%B0%D0%B9%D1%81%D0%BA%22%2C%22lat%22%3A54.7695%2C%22lon%22%3A38.8717%7D; OnboardingState={%22state%22:{%22viewedOnboardingKeys%22:[%22pickup_map_filters%22]}%2C%22version%22:0}; ssrMedia={%22windowWidth%22:880%2C%22primaryInput%22:%22mouse%22}; _808db7ba1248=%5B%7B%22source%22%3A%22%28direct%29%22%2C%22medium%22%3A%22%28none%29%22%2C%22cookie_changed_at%22%3A1717403908%7D%2C%7B%22source%22%3A%22sbermarket.ru%22%2C%22medium%22%3A%22referral%22%2C%22cookie_changed_at%22%3A1717404839%7D%5D; tmr_detect=1%7C1717404840547; rl_group_id=RudderEncrypt%3AU2FsdGVkX19rh9VqU7qS12NqYUO6yjJacnUFV%2FwiQFU%3D; rl_group_trait=RudderEncrypt%3AU2FsdGVkX1%2BZTeuWWxcnFVQwaiQpuAZa5rm3hUUPaUI%3D; rl_anonymous_id=RudderEncrypt%3AU2FsdGVkX18tBHcdrjbs%2Fpgy4jdKINVrXTDkkt0%2FSr6wQ%2FE7WwAWAuOV7F3euFRGTsZC2ukU3MppmkedSt63sg%3D%3D; _Instamart_session=WG1FTjJuTDFwY3BuYXllNmVQNXlERjFlTnd2aHlpVWtvbmdXdUh3UmowaW43eDZXKzBQWFdCZTVKSkcyNzZJQnZpdGk1QmRXZWVjVGpKVU1kN2NvcXN1RHFMRjVkbXZNdGNFY01YeUJnNXJRdjVWYWF6VnRKOWNEYlRQN2poWTkrVUs5WXhsbUZlMDFmTVlwMVZTOEE1ZU03WU5wZ1RmS2IvaitWaTl6dnhibUk3bCtIMm8xNTN5aXlNWVpoMmZWRUxISWYzYno1c2tUQnJiRGhCOHZBaWdPMlg5WExMTWhwaDdWT00zU2RhOG5xSnJDTTZEM0RkcGVqOE1sVk1sOVVqcnlOdnBoNWZqV0ZVWjNMYytBS2FCSlRFSkJvYm9VUlpqbnFQWGRUakkyMjFQUVNHQXNvVWpaNktNakdQTkQtLWZKdTlFRzI0ZHRnT0g1cUtVTWlFVHc9PQ%3D%3D--0feca0b2d56be2b67bca4bd0cef5b53d8c9ab765; rl_user_id=RudderEncrypt%3AU2FsdGVkX1%2B%2Fmr2VOQSLjRcg8bXNddY4Ws3VzUG3%2Bjo%3D; rl_trait=RudderEncrypt%3AU2FsdGVkX18y8TQHqExqDyyXweIvZbDuMub5RadzA2118m9JYe40R5B0caHvtTCw0lvSKtJJrx99oHqTtzWkqzKQlqyfAI9ZK5Who0Xhd6xu5RLXllHX075Y2JunPPDnWRx9gY%2BlPA4OFES9vxfktvBRy0a7jz3pKRE9jOZ34n2LcrPmJjjofyfdnQ9JgEBSF7S%2B5sH%2FCDqoh5qx3lg9oiONGaduq6tI8KX7dl%2F03sNdXSKEq0JrTH3ASSQTjdplM6kzYolKBCFUApNytpNwiunGWJEPKrX5NJ%2BgoYFBBxVXyBE3dw4q%2FkFTK6%2FWADUgD6zDEJKxb6%2B3d7o0Hf2UxPCqyrIY7Icl8dQ0h9QsrSRu8i69ocG9z%2Ft%2Fuhk8v5bxExTr57sseMNksh%2FZQtqbvQ%3D%3D; rl_session=RudderEncrypt%3AU2FsdGVkX19vsejZcEhmHqDvJgOKLryaPwXQwhDPnlyqXa6s9acnU89XjMoykywMcc2S5R0AFnKMae2tEObrVbGECSRMD7jWMOsqD4%2BFDAhCYEn44TwtjididVKIHsJolVQRWA2SyC0y22pZ6ZzWMA%3D%3D; spsc=1717412833001_cf4c01a2f05ea1bd852299388ec4eceb_0d4e0eb9c92181c656f0cb950dff1b59a3402473b747382c8fc1bc9ef31a9efa; t3_sid_7588506=s1.855721052.1717412832939.1717412833118.4.9")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
