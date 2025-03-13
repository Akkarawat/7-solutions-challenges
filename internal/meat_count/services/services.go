package services

import "io"

type MeatTextService interface {
    GetMeatText(hasFiller bool) (io.ReadCloser, error)
}