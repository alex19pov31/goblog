package helpers

type Link struct {
	URL     string
	Text    string
	Current bool
}

type Menu []Link

func (M Menu) UpdateMenu(url string) {
	for i, Link := range M {
		if Link.URL == url {
			M[i].Current = true
		} else {
			M[i].Current = false
		}
	}
}
