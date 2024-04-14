package api

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ITwitterSdk interface {
	Post(message string) error
}

type twitterSdk struct {
}

func NewTwitterSdk() ITwitterSdk {
	return &twitterSdk{}
}

func (s *twitterSdk) Post(message string) error {
	body := map[string]interface{}{
		"text": message,
	}
	header := map[string]string{
		"authorization": fmt.Sprintf(
			"OAuth oauth_consumer_key=%s, oauth_token=%s, oauth_version=1.0",
			os.Getenv("TWITTER_CONSUMER_KEY"),
			os.Getenv("TWITTER_ACCESS_TOKEN"),
		),
	}
	fmt.Println(body, header)
	return nil
}

var oauthNgLetters = [3]string{"+", "/", "="}

func getOauthNonce() string {
	key := make([]byte, 32)
	rand.Read(key)
	enc := base64.StdEncoding.EncodeToString(key)
	for _, letter := range oauthNgLetters {
		enc = strings.Replace(enc, letter, "", -1)
	}
	return enc
}

func manualOauthSettings(creds *creds, additionalParam map[string]string, httpMethod, uri string) string {
	m := map[string]string{}
	m["oauth_consumer_key"] = creds.ConsumerKey
	m["oauth_nonce"] = createoauthNonce()
	m["oauth_signature_method"] = "HMAC-SHA1"
	m["oauth_timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	m["oauth_token"] = creds.AccessToken
	m["oauth_version"] = "1.0"

	baseQueryString := sortedQueryString(mapMerge(m, additionalParam))

	base := []string{}
	base = append(base, url.QueryEscape(httpMethod))
	base = append(base, url.QueryEscape(uri))
	base = append(base, url.QueryEscape(baseQueryString))

	signatureBase := strings.Join(base, "&")

	signatureKey := url.QueryEscape(creds.ConsumerSecret) + "&" + url.QueryEscape(creds.AccessSecret)

	m["oauth_signature"] = calcHMACSHA1(signatureBase, signatureKey)

	authHeader := fmt.Sprintf("OAuth oauth_consumer_key=\"%s\", oauth_nonce=\"%s\", oauth_signature=\"%s\", oauth_signature_method=\"%s\", oauth_timestamp=\"%s\", oauth_token=\"%s\", oauth_version=\"%s\"",
		url.QueryEscape(m["oauth_consumer_key"]),
		url.QueryEscape(m["oauth_nonce"]),
		url.QueryEscape(m["oauth_signature"]),
		url.QueryEscape(m["oauth_signature_method"]),
		url.QueryEscape(m["oauth_timestamp"]),
		url.QueryEscape(m["oauth_token"]),
		url.QueryEscape(m["oauth_version"]),
	)

	return authHeader
}

type sortedQuery struct {
	m    map[string]string
	keys []string
}

func mapMerge(m1, m2 map[string]string) map[string]string {
	m := map[string]string{}

	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}

func sortedQueryString(m map[string]string) string {
	sq := &sortedQuery{
		m:    m,
		keys: make([]string, len(m)),
	}
	var i int
	for key := range m {
		sq.keys[i] = key
		i++
	}
	sort.Strings(sq.keys)

	values := make([]string, len(sq.keys))
	for i, key := range sq.keys {
		values[i] = fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(sq.m[key]))
	}
	return strings.Join(values, "&")
}

func calcHMACSHA1(base, key string) string {
	b := []byte(key)
	h := hmac.New(sha1.New, b)
	io.WriteString(h, base)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
