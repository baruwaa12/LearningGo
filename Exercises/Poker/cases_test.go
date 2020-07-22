package poker

type validTestCase struct {
	name string
	hands []string
	best  []string
}


var validTestCases = []validTestCase{
	{
		name: "single hand always wins",
		hands: []string{
			"4♤ 5♤ 7♡ 8♢ J♧",
		},
		best: []string{"4♤ 5♤ 7♡ 8♢ J♧"},
	},
	{
		name: "highest card out of all hands wins",
		hands: []string{
			"4♢ 5♤ 6♤ 8♢ 3♧",
			"2♤ 4♧ 7♤ 9♡ 10♡",
			"3♤ 4♤ 5♢ 6♡ J♡",
		},
		best: []string{"3♤ 4♤ 5♢ 6♡ J♡"},
	},
	{
		name: "a tie has multiple winners",
		hands: []string{
			"4♢ 5♤ 6♤ 8♢ 3♧",
			"2♤ 4♧ 7♤ 9♡ 10♡",
			"3♤ 4♤ 5♢ 6♡ J♡",
			"3♡ 4♡ 5♧ 6♧ J♢",
		},
		best: []string{"3♤ 4♤ 5♢ 6♡ J♡", "3♡ 4♡ 5♧ 6♧ J♢"},
	},
}
