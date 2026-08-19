package doctor

func Run() Report {
	return Report{
		Categories: []Category{
			CheckSystem(),
			CheckDevelopment(),
		},
	}
}
