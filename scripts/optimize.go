// Truck Loading Optimizer
// Calculates the optimal package combination based on constraints for a particular user based on their email
// - Run with the email you're using to solve the challenge as the command line argument
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// PackageMetadata encapsulates package attributes with dynamic computation
type PackageMetadata struct {
	Identifier     string
	MassConstraint int
	Valuation      int
}

// HeuristicContext holds optimization parameters
type HeuristicContext struct {
	MaxLoad        int
	PriorityFactor float64
}

// DynamicPackageGenerator interface for package creation strategies
type DynamicPackageGenerator interface {
	Generate(email string) []PackageMetadata
}

// EmailBasedPackageGenerator implements package generation
type EmailBasedPackageGenerator struct {
	basePackages []PackageMetadata
}

// NewEmailBasedPackageGenerator initializes the generator
// NOTE: Do not modify the identifiers, constraints, or values of the base packages
func NewEmailBasedPackageGenerator() *EmailBasedPackageGenerator {
	return &EmailBasedPackageGenerator{
		basePackages: []PackageMetadata{
			{Identifier: "A", MassConstraint: 10, Valuation: 60},
			{Identifier: "B", MassConstraint: 20, Valuation: 100},
			{Identifier: "C", MassConstraint: 30, Valuation: 120},
			{Identifier: "D", MassConstraint: 15, Valuation: 80},
			{Identifier: "E", MassConstraint: 25, Valuation: 110},
			{Identifier: "F", MassConstraint: 5, Valuation: 30},
		},
	}
}

// Generate creates packages with email-based modifications
func (g *EmailBasedPackageGenerator) Generate(email string) []PackageMetadata {
	pkgs := make([]PackageMetadata, len(g.basePackages))
	copy(pkgs, g.basePackages)

	// Compute a pseudo-random seed from email using LCG-style accumulation for variability
	// This ensures robust distribution across large input spaces 
	var seed uint64 = 0
	const multiplier uint64 = 0x5DEECE66D // Large constant to promote wide distribution
	const adder uint64 = 0xB // Small additive constant
	for _, r := range email {
		seed = seed*multiplier + uint64(r) + adder
		// No explicit modulo; rely on natural uint64 wraparound for consistency
	}

	// Append dynamic packages with computed attributes
  // NOTE: Do not modify the constraints and valuations of the dynamic packages
	pkgs = append(pkgs, PackageMetadata{
		Identifier:     "X",
		MassConstraint: int((seed % 15) + 5),  // Dynamic weight range
		Valuation:      int((seed % 50) + 40), // Dynamic value range
	})
	pkgs = append(pkgs, PackageMetadata{
		Identifier:     "Y",
		MassConstraint: int((seed % 10) + 8),
		Valuation:      int((seed % 40) + 50),
	})

	return pkgs
}

// LoadOptimizer interface for optimization strategies
type LoadOptimizer interface {
	Optimize(pkgs []PackageMetadata, ctx HeuristicContext) []PackageMetadata
}

// PriorityBasedOptimizer implements a heuristic-based optimization
type PriorityBasedOptimizer struct{}

// Optimize and efficiently select packages based on a heuristic priority

//NOTE: I will not pretend that I came up with this solution on my own. This was part of my experiment to see if an LLM could perform this entire challenge using a sandboxed environment. It excelled at solving the Dynamic Programming challenge. See my explaination and repo at my Github: https://github.com/jaredbarranco/rectanglehq-llm-challenge/blob/main/README.md

func (o *PriorityBasedOptimizer) Optimize(pkgs []PackageMetadata, ctx HeuristicContext) []PackageMetadata {
    n := len(pkgs)
    W := ctx.MaxLoad

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, W+1)
    }

    for i := 1; i <= n; i++ {
        for w := 1; w <= W; w++ {
            if pkgs[i-1].MassConstraint <= w {
                if pkgs[i-1].Valuation+dp[i-1][w-pkgs[i-1].MassConstraint] > dp[i-1][w] {
                    dp[i][w] = pkgs[i-1].Valuation + dp[i-1][w-pkgs[i-1].MassConstraint]
                } else {
                    dp[i][w] = dp[i-1][w]
                }
            } else {
                dp[i][w] = dp[i-1][w]
            }
        }
    }

    var selected []PackageMetadata
    w := W
    for i := n; i > 0 && w > 0; i-- {
        if dp[i][w] != dp[i-1][w] {
            selected = append(selected, pkgs[i-1])
            w -= pkgs[i-1].MassConstraint
        }
    }

    return selected
}

// computePriority calculates a package's priority score
func computePriority(pkg PackageMetadata, factor float64) float64 {
	baseRatio := float64(pkg.Valuation) / float64(pkg.MassConstraint)
	return baseRatio*math.Sqrt(factor) + math.Log1p(float64(pkg.Valuation)) - math.Pow(float64(pkg.MassConstraint), 0.1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing configuration parameter")
		os.Exit(1)
	}

	config := os.Args[1]
	if len(config) == 0 {
		fmt.Fprintln(os.Stderr, "Configuration cannot be empty")
		os.Exit(1)
	}

	// Initialize package generator
	generator := NewEmailBasedPackageGenerator()
	packages := generator.Generate(config)

	// Configure optimizer with heuristic context
	optimizer := &PriorityBasedOptimizer{}
	ctx := HeuristicContext{
		MaxLoad:        50,
		PriorityFactor: 1.0, // Neutral factor to avoid scaling issues
	}

	// Perform optimization
	selected := optimizer.Optimize(packages, ctx)

	// Ensure consistent output ordering
	sort.Slice(selected, func(i, j int) bool {
		return selected[i].Identifier < selected[j].Identifier
	})

	// Format output
	identifiers := make([]string, len(selected))
	for i, pkg := range selected {
		identifiers[i] = pkg.Identifier
	}

	if len(identifiers) == 0 {
		fmt.Print("No viable packages")
	} else {
		fmt.Print(strings.Join(identifiers, ","))
	}
}
