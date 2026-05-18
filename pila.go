package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(url string) {
	if len(s.items) >= 10 {
		s.items = s.items[1:]
	}
	s.items = append(s.items, url)

}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

type Browser struct {
	current string
	history Stack
	forward Stack
}

func (b *Browser) Visit(url string) {
	if b.current != "" {
		b.history.Push(b.current)
	}
	b.forward.items = nil
	b.current = url
}

func (b *Browser) Back() {
	prev := b.history.Pop()
	if prev != "" {
		b.forward.Push(b.current)
		b.current = prev
	}
}

func main() {
	b := Browser{}

	b.Visit("google.com")
	b.Visit("github.com")
	b.Visit("stackoverflow.com")
	b.Visit("animeflv.com")
	b.Visit("youtube.com")
	b.Visit("lectortmo.com")
	b.Visit("aycabronyanosemeocurrenada.com")
	b.Visit("tengomuchosuenoayuda.com")
	b.Visit("tengomuchastareasquehacer.com.mx")
	b.Visit("perdonpakoporretrarsarmecontusactividades.perdoname.com")
	b.Visit("averquepedo.com")

	fmt.Println("Actual:", b.current)

	b.Back()
	fmt.Println("Después de back:", b.current)

	fmt.Println("Tamaño del historial:", len(b.history.items))
	fmt.Println("Historial:", b.history.items)
}
