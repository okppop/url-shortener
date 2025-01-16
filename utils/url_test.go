package utils

import "testing"

var vaildURLs = []string{
	"http://github.io",
	"https://google.com",
	"https://www.xxx.kkk.ok.com",
	"https://www.wikipedia.org/",
	"https://google.com/a?b=c&y=p",
	"https://google.com/a?b=c&y=p#first",
}

var invaildURLs = []string{
	"google.com",
	"ssh://google.com",
	"www.wikipedia.com",
	"htt://oooo.org",
	// "https:iopqwet.net", // need to be fixed
	"//kkkkk.com",
	"localhost",
}

func TestIsURL(t *testing.T) {
	t.Run("vaild_URLs", func(t *testing.T) {
		for _, v := range vaildURLs {
			err := IsURL(v)
			if err != nil {
				t.Errorf("unexpected result: %s got a error: %s", v, err)
			}
		}
	})

	t.Run("invaild_URLs", func(t *testing.T) {
		for _, v := range invaildURLs {
			err := IsURL(v)
			if err == nil {
				t.Errorf("unexpected result: %s didn't get a error", v)
			}
		}
	})
}

func TestGenerateShortPath(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(GenerateShortPath())
	}
}
