package facex

import (
	"testing"

	"github.com/golib/assert"
)

func TestFaceXNewGroup(t *testing.T) {
	assertion := assert.New(t)
	input := NewFacexInput("http://oy5ixix1l.bkt.clouddn.com/face1.png", "1")
	err := testClient().NewGroup(input)
	assertion.Nil(err)
}

func TestFaceXAddGroup(t *testing.T) {
	assertion := assert.New(t)
	err := testClient().AddFace("http://oy5ixix1l.bkt.clouddn.com/miss.png", "5")
	assertion.Nil(err)
}

func TestFaceXSearch(t *testing.T) {
	assertion := assert.New(t)

	facex := testClient()
	result, err := facex.Search("http://oy5ixix1l.bkt.clouddn.com/miss.png")
	assertion.Nil(err)
	assertion.NotNil(result)

	assertion.False(result.IsOK())

	result, err = facex.Search("http://oy5ixix1l.bkt.clouddn.com/face2.png")
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
		AccessKey: "",
		SecretKey: "",
		GroupId:   "facex",

		Timeout:   8,
		Threshold: 0.7,
	})
}
