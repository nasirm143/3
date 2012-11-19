package main

import (
	"code.google.com/p/nimble-cube/cpu"
	"code.google.com/p/nimble-cube/gpu/conv"
	"code.google.com/p/nimble-cube/mag"
	"code.google.com/p/nimble-cube/nimble"
	"fmt"
	"os"
)

func main() {
	nimble.Init()
	defer nimble.Cleanup()
	nimble.SetOD("test4.out")

	N0, N1, N2 := 1, 32, 128
	cx, cy, cz := 3e-9, 3.125e-9, 3.125e-9
	mesh := nimble.NewMesh(N0, N1, N2, cx, cy, cz)
	fmt.Println("mesh:", mesh)

	m := nimble.MakeChan3("m", "", mesh, nimble.UnifiedMemory, 0)
	fmt.Println("m:", m)

	acc := 8
	kernel := mag.BruteKernel(mesh, acc)
	B := conv.NewSymm2D("B", "T", mesh, nimble.UnifiedMemory, kernel, m).Output()

	const Bsat = 1.0053
	const aex = mag.Mu0 * 13e-12 / Bsat
	exch := cpu.NewExchange6("Bex", "T", nimble.UnifiedMemory, m.NewReader(), aex)
	Bex := exch.Output()

	//	heff := MakeChan3("Heff", "", mesh)
	Beff := cpu.NewSum("Beff", B, Bex, Bsat, 1, nimble.UnifiedMemory).Output().Chan3()

	const alpha = 1
	torque := nimble.MakeChan3("τ", "", mesh, nimble.UnifiedMemory, 1)
	nimble.Stack(cpu.NewLLGTorque(torque, m.NewReader(), Beff.NewReader(), alpha))

	const dt = 50e-15

	solver := cpu.NewEuler(m, torque.NewReader(), mag.Gamma0, dt)

	M := cpu.Host(m.ChanN().UnsafeData())
	for i := range M[2] {
		M[2][i] = 1
		M[1][i] = 0.1
	}

	every := 100
	nimble.RunAutosaver("B.dump", B, every)
	nimble.RunAutosaver("m.dump", m, every)
	nimble.RunAutosaver("Bex.dump", Bex, every)
	nimble.RunAutosaver("Beff.dump", Beff, every)
	nimble.RunAutosaver("torque.dump", torque, every)

	nimble.RunStack()

	solver.Steps(100)
	res := cpu.Host(m.ChanN().UnsafeData())
	got := [3]float32{res[0][0], res[1][0], res[2][0]}
	expect := [3]float32{-0.075877085, 0.17907967, 0.9809043}
	fmt.Println("result:", got)
	if got != expect {
		fmt.Println("expected:", expect)
		os.Exit(2)
	}
	solver.Steps(10000)

}
