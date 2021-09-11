package list_op

import "testing"

func TestReverse(t *testing.T) {
	testCase := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reverseList(testCase)
	t.Logf("done")
}
