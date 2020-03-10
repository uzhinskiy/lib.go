package htpass

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

type HTPassFile map[string]string

var IsAuth bool

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func _md5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func _packH32(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return string(hasher.Sum(nil))
}

func _reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func _strtr(ins, from, to string) string {
	tmp := make([]byte, len(ins))
	for i := range ins {
		for f := range from {
			if ins[i] == 32 {
				tmp[i] = 32
			}
			if ins[i] == from[f] {
				tmp[i] = to[f]
			}
		}
	}
	return string(tmp)
}

func (htpf HTPassFile) OpenHTPASSFile(fname string) (err error) {
	rawbytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(rawbytes), "\n")
	for _, line := range lines {
		if len(line) > 0 {
			fields := strings.Split(strings.TrimSpace(line), ":")
			if len(fields) == 2 {
				htpf[strings.TrimSpace(fields[0])] = fields[1]
			}
		}
	}
	return
}

func (htpf HTPassFile) Auth(user string, password string) (bool, error) {
	hash := htpf[user]
	IsAuth = false
	if hash == "" {
		e := "Auth error! User '" + user + "' not found"
		return false, errors.New(e)
	}
	pass := strings.Split(strings.TrimSpace(hash), "$")
	salt := pass[2]
	plen := len(password)
	text := password + "$apr1$" + salt
	bin := _packH32(password + salt + password)

	for i := plen; i > 0; i -= 16 {
		text += bin[0:min(16, i)]
	}

	for k := plen; k > 0; k >>= 1 {
		if k&1 == 1 {
			text += string(0)
		} else {
			text += password[0:1]
		}
	}

	bin = _packH32(text)
	newt := ""
	for i1 := 0; i1 < 1000; i1++ {
		if i1&1 == 1 {
			newt = password
		} else {
			newt = bin
		}
		if i1%3 != 0 {
			newt += salt
		}
		if i1%7 != 0 {
			newt += password
		}
		if i1&1 == 1 {
			newt += bin
		} else {
			newt += password
		}
		bin = _packH32(newt)
	}

	tmp := ""
	for i2 := 0; i2 < 5; i2++ {
		k1 := i2 + 6
		j1 := i2 + 12
		if j1 == 16 {
			j1 = 5
		}
		tmp = bin[i2:i2+1] + bin[k1:k1+1] + bin[j1:j1+1] + tmp
	}
	tmp = string(0) + string(0) + bin[11:12] + tmp

	a1 := base64.StdEncoding.EncodeToString([]byte(tmp))
	b2 := a1[2:]
	c3 := _reverse(b2)
	tmp = _strtr(c3, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	hashedpasswd := "$apr1$" + salt + "$" + tmp
	if hashedpasswd == hash {
		IsAuth = true
		return IsAuth, nil
	} else {
		IsAuth = false
		e := "Auth error! Password mismatch"
		return IsAuth, errors.New(e)
	}
}
