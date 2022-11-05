package models

import "fmt"

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
