package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Distance(o *Point) float64 {
	return math.Pow((p.X-o.X), 2) + math.Pow((p.Y-o.Y), 2)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%f,%f)", p.X, p.Y)
}

func NewPoint(a string) (p *Point) {
	a = strings.Trim(a, "() \n")
	// a = strings.Replace(a, "(", "", -1)
	// a = strings.Replace(a, ")", "", -1)
	// a = strings.Replace(a, " ", "", -1)
	parts := strings.SplitN(a, " ", 2)
	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Point parse error", parts[0], err)
		return
	}
	y, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println("Point parse error", parts[1], err)
		return
	}
	return &Point{
		X: x,
		Y: y,
	}
}

type Pair struct {
	A        *Point
	B        *Point
	Distance float64
}

func main() {
	infile, err := os.Open("hugegrandsmahouse.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(infile)

	var inputPoints []*Point
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Could not read line", err)
				return
			}
			break
		}
		p := NewPoint(line)
		if p != nil {
			inputPoints = append(inputPoints, p)
		}
	}

	closest := &Pair{
		Distance: -1.0, // Starting value
	}

	var closestchan = make(chan *Pair)
	var wg1 sync.WaitGroup

	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for i := 0; i < len(inputPoints); i++ {
			p := <-closestchan
			if closest.Distance < 0 || p.Distance < closest.Distance {
				closest.Distance = p.Distance
				closest.A = p.A
				closest.B = p.B
			}
		}
	}()

	var wg2 sync.WaitGroup
	for j := range inputPoints {
		first := inputPoints[j]
		wg2.Add(1)
		go func(first *Point, retchan chan *Pair) {
			defer wg2.Done()
			p := &Pair{
				Distance: -1.0,
			}
			for i := range inputPoints[j:] {
				dist := first.Distance(inputPoints[i])
				if dist == 0 {
					continue
				}
				if p.Distance < 0 || p.Distance > dist {
					p.A = first
					p.B = inputPoints[i]
					p.Distance = dist
				}
			}

			retchan <- p
		}(first, closestchan)
		if j%1000 == 0 {
			wg2.Wait()
		}
	}

	wg1.Wait()
	fmt.Printf("%s,%s\n", closest.A, closest.B)
}
