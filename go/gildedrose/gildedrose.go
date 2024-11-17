package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
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

		if items[i].Name != ageBrie && items[i].Name != backStage {
			if items[i].Quality > 0 {
				items[i].Quality = items[i].Quality - 1
			}
		} else {
			if items[i].Quality < 50 {
				addingQuality := 1

				if items[i].Name == backStage {
					if items[i].SellIn < 11 {
						addingQuality += 1
					}

					if items[i].SellIn < 6 {
						addingQuality += 1
					}
				}

				items[i].Quality += addingQuality
				if items[i].Quality > 50 {
					items[i].Quality = 50
				}
			}
		}

		items[i].SellIn = items[i].SellIn - 1

		if items[i].SellIn < 0 {
			if items[i].Name == backStage {
				items[i].Quality = 0
				continue
			}

			if items[i].Name != ageBrie && items[i].Quality > 0 {
				items[i].Quality = items[i].Quality - 1

			} else if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}

		}
	}

}
