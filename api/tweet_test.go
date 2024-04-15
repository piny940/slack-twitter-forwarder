package api

import (
	"slices"
	"strings"
	"testing"
)

func TestOauthNonce(t *testing.T) {
	nonces := make([]string, 0)
	for i := 0; i < 100; i++ {
		nonce := getOauthNonce()
		for _, letter := range oauthNgLetters {
			if strings.Contains(nonce, letter) {
				t.Errorf("Nonce must not contain %s, but nonce is %s", letter, nonce)
				return
			}
		}
		if slices.Contains(nonces, nonce) {
			t.Errorf("Nonce must be unique, but nonce is %s", nonce)
			return
		}
		nonces = append(nonces, nonce)
	}
}

func TestMapMerge(t *testing.T) {
	m1 := map[string]string{"a": "1", "b": "2"}
	m2 := map[string]string{"c": "3", "d": "4"}
	m := mapMerge(m1, m2)
	if len(m) != 4 {
		t.Errorf("Length of merged map must be 4, but length is %d", len(m))
		return
	}
	if m["a"] != "1" || m["b"] != "2" || m["c"] != "3" || m["d"] != "4" {
		t.Errorf("Merged map must be {a: 1, b: 2, c: 3, d: 4}, but map is %v", m)
		return
	}
}

func TestCalcHMACSHA1(t *testing.T) {
	baseString := "POST&https%3A%2F%2Fapi.twitter.com%2F1.1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"
	key := "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw&LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"
	signature := calcHMACSHA1(baseString, key)
	expected := "hCtSmYh+iHYCEqBWrE7C7hYmtUk="
	if signature != expected {
		t.Errorf("Signature must be %s but actual was %s", expected, signature)
		return
	}
}

func TestGetSignatureBaseString(t *testing.T) {
	method := "POST"
	uri := "https://api.twitter.com/1.1/statuses/update.json"
	paramsString := "include_entities=true&oauth_consumer_key=xvz1evFS4wEEPTGEFPHBog&oauth_nonce=kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1318622958&oauth_token=370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb&oauth_version=1.0&status=Hello%20Ladies%20%2B%20Gentlemen%2C%20a%20signed%20OAuth%20request%21"
	actual := getSignatureBaseString(method, uri, paramsString)
	expected := "POST&https%3A%2F%2Fapi.twitter.com%2F1.1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"
	if actual != expected {
		t.Errorf("Signature Base String must be %s but actual was %s", expected, actual)
		return
	}
}

func TestSortedQueryString(t *testing.T) {
	m := map[string]string{
		"b": "2",
		"a": "1",
		"d": "4",
	}
	actual := sortedQueryString(m)
	expected := "a=1&b=2&d=4"
	if actual != expected {
		t.Errorf("Sorted query string must be %s but actual was %s", expected, actual)
		return
	}
}

func TestGetSigningKey(t *testing.T) {
	cred := &credentials{
		ConsumerKey:       "consumer_key",
		ConsumerSecret:    "consumer_secretほげ",
		AccessToken:       "access_token",
		AccessTokenSecret: "access_token_secretふが",
	}
	key := getSigningKey(cred)
	expected := "consumer_secret%E3%81%BB%E3%81%92&access_token_secret%E3%81%B5%E3%81%8C"
	if key != expected {
		t.Errorf("Signing Key must be %s but actual was %s", expected, key)
		return
	}
}
