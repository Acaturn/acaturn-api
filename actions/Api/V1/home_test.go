package V1

func (as *V1Suite) TestHomeHandler() {
	res := as.JSON("/").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
