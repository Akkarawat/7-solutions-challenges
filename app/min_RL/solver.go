package minrl

import (
	"strconv"
	"strings"
)

type CompType int

const (
    Left CompType = iota
    Right
    Equal
)

type Segment struct {
    CompType CompType
    Value int

    LeftSegment *Segment
    RightSegment *Segment
}

func NewFirstSegment() *Segment {
    return &Segment{
        Value: 0,
    }
}

func (s *Segment) AddRightSegment(compType CompType) *Segment {
    if s.Value == 0 && compType == Left {
        s.Value++
    }

    rightSegment := Segment{
        CompType: compType,
        LeftSegment: s,
    }

    if compType == Left {
        rightSegment.Value = 0
    } else if compType == Right {
        rightSegment.Value = s.Value + 1
    } else {
        rightSegment.Value = s.Value
    }
    s.RightSegment = &rightSegment
    return s
}

func (s *Segment) EnsureValidity() {
    if s.IsFirst() {
        return
    }
    if (s.CompType == Equal && s.LeftSegment.Value < s.Value) ||
        (s.CompType == Left && s.LeftSegment.Value == s.Value) {
        s.LeftSegment.Value++
        s.LeftSegment.EnsureValidity()
    }
}

func (s *Segment) IsFirst() bool {
    return s.LeftSegment == nil
}

func SolveMinRL (input string) string {
    currentSegment := NewFirstSegment()

    tokens := strings.Split(input, "")
    for _, token := range tokens {
        var compType CompType
        switch token {
            case "L": {
                compType = Left
            }
            case "R": {
                compType = Right
            }
            default:
                compType = Equal
        }
        currentSegment.AddRightSegment(compType)
        currentSegment.EnsureValidity()
        currentSegment = currentSegment.RightSegment
    }

    result := ""
    for {
        result = strconv.Itoa(currentSegment.Value) + result
        if currentSegment.IsFirst() {
            break
        }
        currentSegment = currentSegment.LeftSegment
    }
    return result
}