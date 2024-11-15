package models

import "strings"

type foodItem struct {
	date string
	food string
	tags []string
}

func (fi *foodItem) fromSlice(src []string) error {
	fi.date = src[0]
	fi.food = src[1]
	fi.tags = strings.Split(src[2], ";")
	return nil
}

func (fi *foodItem) toSlice() []string {
	newSlice := make([]string, 0)
	newSlice = append(newSlice, fi.date, fi.food, strings.Join(fi.tags, ";"))
	return newSlice
}

func newFoodItem(date string, food string, tags string) foodItem {
	newFood := foodItem{
		date: date,
		food: food,
		tags: strings.Split(tags, ";"),
	}
	return newFood
}