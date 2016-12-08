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
