package usecases

import (
	"7-solutions-challenges/internal/meat_count/services"
	"7-solutions-challenges/internal/meat_count/utils"
	"io"
	"strings"
)

type MeatCountUsecase struct {
    meatTextService services.MeatTextService
}

func NewMeatCountUsecase(meatTextService services.MeatTextService) *MeatCountUsecase {
    return &MeatCountUsecase{
        meatTextService: meatTextService,
    }
}

func (uc *MeatCountUsecase) GetMeatCount() (map[string]int, error) {
    meatSets, err := uc.getMeatSets()
    if err != nil {
        return nil, err
    }
    meatCounts := make(map[string]int)

    meatAndFillerReader, err := uc.meatTextService.GetMeatText(true)
    if err != nil {
        return nil, err
    }
    defer meatAndFillerReader.Close()
    meatAndFillerWordReader := utils.NewWordReader(meatAndFillerReader)

    for {
        word, err := meatAndFillerWordReader.NextWord()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        normalizedWord := strings.ToLower(word)

        // if the word is a meat, increment the count
        if _, ok := meatSets[normalizedWord]; ok {
            meatCounts[normalizedWord]++
        }
    }
    meatAndFillerReader.Close()
    return meatCounts, nil
}

func (uc *MeatCountUsecase) getMeatSets() (map[string]bool, error) {
    reader, err := uc.meatTextService.GetMeatText(false)
    if err != nil {
        return nil, err
    }
    defer reader.Close()

    sets := make(map[string]bool)
    wordReader := utils.NewWordReader(reader)
    for {
        word, err := wordReader.NextWord()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        normalizedWord := strings.ToLower(word)
        sets[normalizedWord] = true
    }

    return sets, nil
}