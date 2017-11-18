package facex

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/golib/assert"
)

func TestFaceXNewGroup(t *testing.T) {
	assertion := assert.New(t)
	input := NewFacexInput(testFace("./face1.png"), "1")
	err := testClient().NewGroup(input)
	assertion.Nil(err)
}

func TestFaceXAddGroup(t *testing.T) {
	assertion := assert.New(t)
	err := testClient().AddFace(testFace("./face1.png"), "2")
	assertion.Nil(err)
}

func TestFaceXSearch(t *testing.T) {
	assertion := assert.New(t)

	facex := testClient()
	result, err := facex.Search(testFace("./miss.png"))
	assertion.Nil(err)
	assertion.NotNil(result)

	assertion.False(result.IsOK())

	result, err = facex.Search(testFace("./face2.png"))
	assertion.Nil(err)
	assertion.NotNil(result)

	assertion.True(result.IsOK())
}

func TestFaceXRemoveGroup(t *testing.T) {
	assertion := assert.New(t)

	facex := testClient()

	err := facex.RemoveGroup()
	assertion.Nil(err)
}

func testClient() *Facex {
	return NewFacex(&Config{
		Endpoint:  "http://argus.atlab.ai",
		AccessKey: os.Getenv("QINIUAK"),
		SecretKey: os.Getenv("QINIUSK"),

		GroupId:   "testfacex",
		Timeout:   8,
		Threshold: 0.7,
	})
}

func testFace(path string) string {
	dat, _ := ioutil.ReadFile(path)
	return NewFaceBase64(dat)
}
