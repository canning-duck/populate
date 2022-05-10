package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"strconv"
)

func main() {
	var (
		err error
		p   = flag.String("p", "", "filename prefix (default random)")
		n   = flag.Int("n", 260721, "number of files")
		s   = flag.Int("s", 1024, "file size in bytes")
		r   = flag.Bool("r", false, "randomize file contents (default empty)")
		d   = flag.Int("d", 0, "number of directories")
	)
	flag.Parse()
	if *p == "" {
		b := make([]byte, 12)
		_, err = rand.Read(b)
		check(err)
		*p = base64.RawURLEncoding.EncodeToString(b)
	}
	dirs := []string{""}
	for i := 0; i < *d; i++ {
		name := path.Join(dirs[randInt(len(dirs))], *p+"-"+strconv.Itoa(1+i))
		err = os.Mkdir(name, 0700)
		check(err)
		dirs = append(dirs, name)
	}
	data := make([]byte, *s)
	for i := *d; i < *d+*n; i++ {
		if *r {
			_, err := rand.Read(data)
			check(err)
		}
		name := path.Join(dirs[randInt(len(dirs))], *p+"-"+strconv.Itoa(1+i))
		err := ioutil.WriteFile(name, data, 0600)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func randInt(n int) int {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	check(err)
	return int(i.Int64())
}
