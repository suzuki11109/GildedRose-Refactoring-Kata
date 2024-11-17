package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) summarizeQuality(quality int) {
	if i.Quality == 0 || i.Quality == 50 {
		return
	}

	i.Quality += quality
	if i.Quality > 50 {
		i.Quality = 50
	}

	if i.Quality < 0 {
		i.Quality = 0
	}
}

const (
	backStage = "Backstage passes to a TAFKAL80ETC concert"
	ageBrie   = "Aged Brie"
	sulfuras  = "Sulfuras, Hand of Ragnaros"
)

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name == sulfuras {
			continue
		}

		if items[i].Name == backStage {
			addingQuality := 1
			if items[i].SellIn < 11 {
				addingQuality += 1
			}

			if items[i].SellIn < 6 {
				addingQuality += 1
			}
			items[i].summarizeQuality(addingQuality)
		} else if items[i].Name == ageBrie {
			items[i].summarizeQuality(1)
		} else {
			items[i].summarizeQuality(-1)
		}

		items[i].SellIn = items[i].SellIn - 1

		if items[i].SellIn < 0 {
			if items[i].Name == backStage {
				items[i].Quality = 0
				continue
			}

			if items[i].Name == ageBrie {
				items[i].summarizeQuality(1)
			} else {
				items[i].summarizeQuality(-1)
			}

		}
	}

}
