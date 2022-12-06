package models

import (
	"fmt"
	"log"
	"os"
)

type Projectile struct {
	position Tuple // a point
	velocity Tuple // a vector
}

type Environment struct {
	gravity Tuple // a vector
	wind    Tuple // a vector
}

func tick(env Environment, proj Projectile) Projectile {
	position := Add(proj.position, proj.velocity)
	velocity := Add(Add(proj.velocity, env.gravity), env.wind)
	return Projectile{position: position, velocity: velocity}
}

// Playing around with the cannon
func ShootCannon() {
	var p Projectile = Projectile{NewPoint(0, 1, 0), Normalize(NewVector(1, 1, 0))}
	var e Environment = Environment{NewVector(0, -0.1, 0), NewVector(-0.01, 0, 0)}
	println("Shooting the cannon")
	for p.position.Y >= 1 {
		p = tick(e, p)
		fmt.Printf("Position: %v, Velocity: %v\n", p.position, p.velocity)
	}
	println("Projectile landed")
}

func ShootCannonAndDraw() {
	start := NewPoint(0, 1, 0)
	velocity := Multiply(Normalize(NewVector(1, 1.8, 0)), 11.25)
	p := Projectile{start, velocity}

	gravity := NewVector(0, -0.1, 0)
	wind := NewVector(-0.01, 0, 0)
	e := Environment{gravity, wind}

	canvas := NewCanvas(900, 550)

	for p.position.Y >= -550 {
		p = tick(e, p)
		if p.position.X < 900 && 0 <= p.position.Y {
			log.Printf("position is %v", p)
			canvas.WritePixel(int64(p.position.X), 550-int64(p.position.Y), Color{1, 0, 0})
		}
	}

	plot := canvas.ToPPM()

	err := os.WriteFile("./../plot.ppm", []byte(plot), 0644)

	if err != nil {
		log.Fatal("File write failed")
	}
}
