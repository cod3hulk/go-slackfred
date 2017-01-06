package alfred

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := new(Result)

	item := new(Item)
	item.Arg = "TestArg"
	item.Subtitle = "TestSubtitle"
	item.Title = "TestTitle"

	result.Add(item)

	if len(result.Items) != 1 {
		t.Error("Expected 1 got ", len(result.Items))
	}
}

func TestAddAll(t *testing.T) {
	result := new(Result)
	items := []Item{
		Item{
			Arg:      "TestArg",
			Subtitle: "TestSubtitle",
			Title:    "TestTitle",
		},
		Item{
			Arg:      "TestArg2",
			Subtitle: "TestSubtitle2",
			Title:    "TestTitle2",
		},
	}

	result.AddAll(items)

	if len(result.Items) != 2 {
		t.Error("Expected 2 got ", len(result.Items))
	}
}

func TestFilter(t *testing.T) {
	result := new(Result)
	items := []Item{
		Item{
			Arg:      "TestArg",
			Subtitle: "TestSubtitle",
			Title:    "TestTitle",
		},
		Item{
			Arg:      "TestArg2",
			Subtitle: "TestSubtitle2",
			Title:    "TestTitle2",
		},
	}
	result.Items = items

	result.Filter("TestArg", func(item Item, query string) bool {
		return item.Arg == query
	})

	if len(result.Items) != 1 {
		t.Error("Expected 1 got ", len(result.Items))
	}
}

func TestOutput(t *testing.T) {
	result := new(Result)

	item := new(Item)
	item.Arg = "TestArg"
	item.Subtitle = "TestSubtitle"
	item.Title = "TestTitle"

	output := result.Add(item).Output()

	if output != `{"items":[{"title":"TestTitle","subtitle":"TestSubtitle","arg":"TestArg"}]}` {
		t.Error("Expected {items} got ", output)
	}
}
