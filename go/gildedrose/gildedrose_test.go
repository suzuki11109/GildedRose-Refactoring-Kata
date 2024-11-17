package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	var items = []gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Normal Item", 2, 3},
		{"Aged Brie", 2, 48},
		{"Conjured", 3, 5},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 25},
		{"Aged Brie", 2, 10},
	}

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[1], 2, 1)
	assetQualitySellIn(t, items[2], 49, 1)
	assetQualitySellIn(t, items[3], 3, 2)
	assetQualitySellIn(t, items[4], 26, 11)
	assetQualitySellIn(t, items[5], 11, 1)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[1], 1, 0)
	assetQualitySellIn(t, items[2], 50, 0)
	assetQualitySellIn(t, items[3], 1, 1)
	assetQualitySellIn(t, items[4], 27, 10)
	assetQualitySellIn(t, items[5], 12, 0)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[1], 0, -1)
	assetQualitySellIn(t, items[2], 50, -1)
	assetQualitySellIn(t, items[3], 0, 0)
	assetQualitySellIn(t, items[4], 29, 9)
	assetQualitySellIn(t, items[5], 14, -1)

	gildedrose.UpdateQuality(items) // day 4
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[1], 0, -2)
	assetQualitySellIn(t, items[3], 0, -1)
	assetQualitySellIn(t, items[4], 31, 8)

	gildedrose.UpdateQuality(items) // day 5
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[4], 33, 7)

	gildedrose.UpdateQuality(items) // day 6
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[4], 35, 6)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[0], 80, 0)
	assetQualitySellIn(t, items[4], 37, 5)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[4], 40, 4)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[4], 43, 3)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[4], 46, 2)

	gildedrose.UpdateQuality(items)
	assetQualitySellIn(t, items[4], 49, 1)

	gildedrose.UpdateQuality(items) // day 12
	assetQualitySellIn(t, items[4], 50, 0)

	gildedrose.UpdateQuality(items) // day 13
	assetQualitySellIn(t, items[4], 0, -1)
}

func assetQualitySellIn(t *testing.T, item gildedrose.Item, expectedQuality, expectedSellIn int) {
	if item.Quality != expectedQuality {
		t.Errorf("want %s quality = %v but got %v", item.Name, expectedQuality, item.Quality)
	}
	if item.SellIn != expectedSellIn {
		t.Errorf("want %s sellin = %v but got %v", item.Name, expectedSellIn, item.SellIn)
	}
}
