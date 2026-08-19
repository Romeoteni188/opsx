package doctor

type CheckResult struct {
	Name    string `json:"name"`
	Passed  bool   `json:"passed"`
	Message string `json:"message"`
}

type Category struct {
	Name   string        `json:"name"`
	Checks []CheckResult `json:"checks"`
}

type Report struct {
	Categories []Category `json:"categories"`
}

func (r *Report) Total() int {
	total := 0

	for _, category := range r.Categories {
		total += len(category.Checks)
	}

	return total
}

func (r *Report) Passed() int {
	total := 0

	for _, category := range r.Categories {
		for _, check := range category.Checks {
			if check.Passed {
				total++
			}
		}
	}

	return total
}

func (r *Report) Failed() int {
	return r.Total() - r.Passed()
}
