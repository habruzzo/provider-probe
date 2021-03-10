package src

import "crypto/sha256"

//decrypt datakey (using kms) and emails (using datakey)

//returns hash of email. does it need to be sha 256? prob not.
func HashEmailForTimingLogs(email string) []byte {
	h:= sha256.New()
	h.Write([]byte(email))
	return h.Sum(nil)
}