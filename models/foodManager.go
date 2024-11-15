package models

import (
	"bufio"
	"encoding/csv"
	"os"
	"slices"
)

type foodManager struct {
	items []foodItem
}

func (fm *foodManager) load(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil && err == os.ErrNotExist {
		f, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	br := bufio.NewReader(f)
	csvreader := csv.NewReader(br)

	records, err := csvreader.ReadAll()
	if err != nil {
		return err
	}

	for _, v := range records {
		fi := foodItem{}
		fi.fromSlice(v)
		fm.items = append(fm.items, fi)
	}

	return nil
}

func (fm *foodManager) remove(index int) {
	if index < 0 || index > len(fm.items) {
		return
	}
	fm.items = append(fm.items[0:index], fm.items[index+1:]...)
}

func (fm *foodManager) save(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil && err == os.ErrNotExist {
		f, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	bw := bufio.NewWriter(f)
	csvwriter := csv.NewWriter(bw)
	writePayload := make([][]string, 0, len(fm.items))

	writePayload = append(writePayload, []string{"Date", "Food", "Tags"})
	for _, v := range fm.items {
		writePayload = append(writePayload, v.toSlice())
	}

	err = csvwriter.WriteAll(writePayload)
	if err != nil {
		return err
	}
	return nil
}

func (fm *foodManager) sort() {
	slices.SortFunc(fm.items, func(a foodItem, b foodItem) int {
		date1 := a.date
		date2 := b.date

		rdate1 := []rune(date1)
		rdate2 := []rune(date2)

		for k, v := range rdate1 {
			if v > rdate2[k] {
				return 1
			} else if v < rdate2[k] {
				return -1
			}
		}
		return 0
	})
}