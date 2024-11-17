package gildedrose

const (
	backStage = "Backstage passes to a TAFKAL80ETC concert"
	ageBrie   = "Aged Brie"
	sulfuras  = "Sulfuras, Hand of Ragnaros"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) addQuality(quality int) {
	i.Quality += quality
	if i.Quality > 50 {
		i.Quality = 50
	}

	if i.Quality < 0 {
		i.Quality = 0
	}
}

func UpdateQuality(items []Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name == sulfuras {
			continue
		}

		items[i].SellIn = items[i].SellIn - 1

		addingQuality := 1
		if items[i].Name == backStage {
			if items[i].SellIn < 0 {
				items[i].Quality = 0
				continue
			}

			if items[i].SellIn < 10 {
				addingQuality += 1
			}

			if items[i].SellIn < 5 {
				addingQuality += 1
			}

		} else if items[i].Name == ageBrie {

			if items[i].SellIn < 0 {
				addingQuality += 1
			}

		} else {
			addingQuality = -1

			if items[i].SellIn < 0 {
				addingQuality -= 1
			}
		}

		items[i].addQuality(addingQuality)

	}

}
