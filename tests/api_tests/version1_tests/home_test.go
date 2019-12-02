package version1_tests

var root = "/api/v1"

func (as *V1Suite) TestHomeHandler() {
	res := as.JSON(root).Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
