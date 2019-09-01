package V1

import (
	"acaturn_api/actions"
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
)

type V1Suite struct {
	*suite.Action
}

func TestV1Suite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(actions.App(), packr.New("TestV1Suite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &V1Suite{
		Action: action,
	}
	suite.Run(t, as)
}
