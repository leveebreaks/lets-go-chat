package service

import "time"

var tokens = make(map[string]string)

// AddToken adds a token to the cache and removes it if not revoked within 5 minutes
func AddToken(token, userName string) {
	tokens[token] = userName
	time.AfterFunc(time.Minute*5, func() {
		RevokeToken(token)
	})
}

// RevokeToken revokes token by removing it from the cache. Returns true if it existed, which means successful revoking
func RevokeToken(token string) bool {
	_, ok := tokens[token]
	if ok {
		delete(tokens, token)
		return true
	}
	return false
}
