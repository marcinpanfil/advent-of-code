package aoc2017

import (
	"advent-of-code/util/math"
	"sort"
	"strconv"
	"strings"
)

type Particle struct {
	id       int
	p        Coordinates
	v        Coordinates
	a        Coordinates
	collided bool
}

type Coordinates struct {
	x int
	y int
	z int
}

func (c *Coordinates) ManhattanDist() int {
	return math.Abs(c.x) + math.Abs(c.y) + math.Abs(c.z)
}

func (p *Particle) Move() {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func Collide(particles []Particle) int {
	for m := 0; m < 50; m++ {
		CheckForCollisions(particles)
		MoveOneStep(particles)
	}

	return CountCollided(particles)
}

func CountCollided(particles []Particle) int {
	count := 0
	for _, p := range particles {
		if !p.collided {
			count++
		}
	}
	return count
}

func MoveOneStep(particles []Particle) {
	for i := 0; i < len(particles)-1; i++ {
		p := particles[i]
		if !p.collided {
			p.Move()
			particles[i] = p
		}
	}
}

func CheckForCollisions(particles []Particle) {
	for i := 0; i < len(particles)-1; i++ {
		p1 := particles[i]
		if !p1.collided {
			for j := i + 1; j < len(particles); j++ {
				p2 := particles[j]
				if !p2.collided && p1.p == p2.p {
					p1.collided = true
					p2.collided = true
					particles[i] = p1
					particles[j] = p2
				}
			}
		}
	}
}

func FindParticleWithMinDist(particles []Particle) int {
	sort.Slice(particles, func(i, j int) bool {
		p1 := particles[i]
		p2 := particles[j]
		if p1.a != p2.a {
			return p1.a.ManhattanDist() < p2.a.ManhattanDist()
		}
		if p1.v != p2.v {
			return p1.v.ManhattanDist() < p2.v.ManhattanDist()
		}
		return p1.p.ManhattanDist() < p2.p.ManhattanDist()
	})
	return particles[0].id
}

func ParseInput(input []string) []Particle {
	particles := []Particle{}
	for i, line := range input {
		parts := strings.Split(line, ", ")
		particles = append(particles, Particle{id: i, p: CreateFromStr(parts[0]), v: CreateFromStr(parts[1]), a: CreateFromStr(parts[2])})
	}
	return particles
}

func CreateFromStr(input string) Coordinates {
	startIdx := strings.LastIndex(input, "<")
	endIdx := strings.LastIndex(input, ">")
	data := strings.Split(input[startIdx+1:endIdx], ",")
	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])
	z, _ := strconv.Atoi(data[2])
	return Coordinates{x: x, y: y, z: z}
}
