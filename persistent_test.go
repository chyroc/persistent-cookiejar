package cookiejar

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestLoadSaveNoMaxAgeAndExpire_NoPersistent(t *testing.T) {
	c := qt.New(t)
	d, err := ioutil.TempDir("", "")
	c.Assert(err, qt.Equals, nil)
	defer os.RemoveAll(d)
	file := filepath.Join(d, "cookies")
	j := newTestJar(file)
	j.SetCookies(serializeTestURL, serializeTestCookiesNoMaxAge)
	err = j.Save()
	c.Assert(err, qt.Equals, nil)
	_, err = os.Stat(file)
	c.Assert(err, qt.Equals, nil)
	j1 := newTestJar(file)
	c.Assert(len(j1.entries), qt.Equals, 0)
}

func TestLoadSaveNoMaxAgeAndExpire_Persistent(t *testing.T) {
	c := qt.New(t)
	d, err := ioutil.TempDir("", "")
	c.Assert(err, qt.Equals, nil)
	defer os.RemoveAll(d)
	file := filepath.Join(d, "cookies")
	j := newTestJarNoMaxAge(file)
	j.SetCookies(serializeTestURL, serializeTestCookiesNoMaxAge)
	err = j.Save()
	c.Assert(err, qt.Equals, nil)
	_, err = os.Stat(file)
	c.Assert(err, qt.Equals, nil)
	j1 := newTestJarNoMaxAge(file)
	c.Assert(len(j1.entries), qt.Equals, len(serializeTestCookies))
	c.Assert(j1.entries, qt.DeepEquals, j.entries)
}

func newTestJarNoMaxAge(path string) *Jar {
	jar, err := New(&Options{
		PublicSuffixList: testPSL{},
		Filename:         path,
		NoPersist:        path == "",
		Persistent:       true,
	})
	if err != nil {
		panic(err)
	}
	return jar
}

var serializeTestCookiesNoMaxAge = []*http.Cookie{{
	Name:   "foo",
	Value:  "bar",
	Path:   "/p",
	Domain: "example.com",
}}
