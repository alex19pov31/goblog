package main

type link struct {
	url     string
	text    string
	current bool
}

func (l *link) URL() string {
	return l.url
}
func (l *link) Text() string {
	return l.text
}
func (l *link) Current() bool {
	return l.current
}

type menu []link

func (m menu) updateMenu(url string) {
	for i, link := range m {
		if link.url == url {
			m[i].current = true
		} else {
			m[i].current = false
		}
	}
}
