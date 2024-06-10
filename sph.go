package main

import "fmt"

// Plane interface
type Plane interface {
    Fly() string
}

// Base Plane struct
type BasePlane struct {
    Model string
    Capacity int
}

// Method for BasePlane
func (p BasePlane) Info() string {
    return fmt.Sprintf("Model: %s, Capacity: %d", p.Model, p.Capacity)
}

// Jet struct
type Jet struct {
    BasePlane
    EngineCount int
}

// Implement the Fly method for Jet
func (j Jet) Fly() string {
    return "Zooom!"
}

// PropellerPlane struct
type PropellerPlane struct {
    BasePlane
    WingSpan int
}

// This is for implement the Fly method for PropellerPlane
func (p PropellerPlane) Fly() string {
    return "Whirr!"
}

func main() {
    jet := Jet{
        BasePlane:  BasePlane{Model: "F-16", Capacity: 2},
        EngineCount: 2,
    }

    propellerPlane := PropellerPlane{
        BasePlane:  BasePlane{Model: "Cessna 172", Capacity: 4},
        WingSpan:   36,
    }

    planes := []Plane{jet, propellerPlane}

    for _, plane := range planes {
        fmt.Println(plane.Fly())
        switch p := plane.(type) {
        case Jet:
            fmt.Println("Engine Count:", p.EngineCount)
            fmt.Println(p.Info())
        case PropellerPlane:
            fmt.Println("Wing Span:", p.WingSpan)
            fmt.Println(p.Info())
        }
    }
}

/*

Output

Zoom!
Engine Count: 2
Model: F-16, Capacity: 2
Whirr!
Wing Span: 36
Model: Cessna 172, Capacity: 4

*/