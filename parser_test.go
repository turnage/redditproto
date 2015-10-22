package redditproto

import (
	"testing"
)

func TestParseComment(t *testing.T) {
	replyTree := `{
		"kind": "Listing",
		"data": {
			"children": [
				{
					"kind": "t1",
					"data": {
						"body": "reply1",
						"replies": ""
					}
				},
				{
					"kind": "t1",
					"data": {
						"body": "reply2",
						"replies": ""
					}
				}
			]
		}
	}`
	comment, err := ParseComment([]byte(
		`{
			"kind": "t1",
			"data": {
				"body": "something",
				"replies": ` + replyTree + `
			}
		}`),
	)
	if err != nil {
		t.Fatal(err)
	}

	if comment == nil {
		t.Fatal("returned comment was nil")
	}

	if len(comment.Replies) != 2 {
		t.Fatalf("got %d replies; wanted 2", len(comment.Replies))
	}

	if comment.GetBody() != "something" {
		t.Errorf("got %s; wanted something", comment.GetBody())
	}

	if comment.Replies[0].GetBody() != "reply1" {
		t.Errorf("got %s; wanted reply1", comment.Replies[0].GetBody())
	}

	if comment.Replies[1].GetBody() != "reply2" {
		t.Errorf("got %s; wanted reply2", comment.Replies[1].GetBody())
	}
}

func TestParseLinkListing(t *testing.T) {
	listing, err := ParseLinkListing([]byte(
		`{
			"kind": "Listing",
			"data": {
				"children": [
					{
						"kind": "t3",
						"data": {
							"title": "1"
						}
					},
					{
						"kind": "t3",
						"data": {
							"title": "2"
						}
					},
					{
						"kind": "t3",
						"data": {
							"title": "3"
						}
					}
				]
			}
		}`))
	if err != nil {
		t.Fatal(err)
	}

	if len(listing) != 3 {
		t.Fatalf("got %d links; wanted 3", len(listing))
	}

	if listing[2].GetTitle() != "3" {
		t.Errorf("got %s as title; wanted 3", listing[2].GetTitle())
	}
}

func TestParseMessageListing(t *testing.T) {
	listing, err := ParseMessageListing([]byte(
		`{
			"kind": "Listing",
			"data": {
				"children": [
					{
						"kind": "t4",
						"data": {
							"body": "1"
						}
					},
					{
						"kind": "t4",
						"data": {
							"body": "2"
						}
					}
				]
			}
		}`))
	if err != nil {
		t.Fatal(err)
	}

	if len(listing) != 2 {
		t.Fatalf("got %d messages; wanted 2", len(listing))
	}

	if listing[1].GetBody() != "2" {
		t.Errorf("got %s as body; wanted 2", listing[1].GetBody())
	}
}

func TestParseThread(t *testing.T) {
	link, err := ParseThread([]byte(
		`[
			{
				"kind": "Listing",
				"data": {
					"children": [
						{
							"kind": "t3",
							"data": {
								"title": "1"
							}
						}
					]
				}
			},
			{
				"kind": "Listing",
				"data": {
					"children": [
						{
							"kind": "t1",
							"data": {
								"body": "1"
							}
						},
						{
							"kind": "t1",
							"data": {
								"body": "2"
							}
						}
					]
				}
			}
		]`))
	if err != nil {
		t.Fatal(err)
	}

	if link == nil {
		t.Fatalf("returned link was nil")
	}

	if len(link.Comments) != 2 {
		t.Errorf("got %d comments; wanted 2", len(link.Comments))
	}
}

func TestParseComboListing(t *testing.T) {
	links, comments, err := ParseComboListing([]byte(
		`{
			"kind": "Listing",
			"data": {
				"children": [
					{
						"kind": "t3",
						"data": {
							"title": "1"
						}
					},
					{
						"kind": "t1",
						"data": {
							"body": "1"
						}
					}
				]
			}
		}`))
	if err != nil {
		t.Fatal(err)
	}

	if len(links) != 1 {
		t.Errorf("got %d links; wanted 1", len(links))
	}

	if len(comments) != 1 {
		t.Errorf("got %d comments; wanted 1", len(comments))
	}
}
