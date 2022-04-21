package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"strconv"
)

func main() {
	var (
		err error
		p   = flag.String("p", "", "filename prefix (default random)")
		n   = flag.Int("n", 260721, "number of files")
		s   = flag.Int("s", 1024, "file size in bytes")
		r   = flag.Bool("r", false, "randomize file contents (default empty)")
	)
	flag.Parse()
	if *p == "" {
		b := make([]byte, 12)
		_, err = rand.Read(b)
		check(err)
		*p = base64.RawURLEncoding.EncodeToString(b)
	}
	data := make([]byte, *s)
	for i := 0; i < *n; i++ {
		if *r {
			_, err := rand.Read(data)
			check(err)
		}
		err := ioutil.WriteFile(*p+strconv.Itoa(i), data, 0600)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
