package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestList_Len(t *testing.T) {
	TestTable := []struct {
		listForList []int
		expected    int
		testName    string
	}{
		{
			listForList: []int{1, 2, 3},
			expected:    3,
			testName:    "test-1",
		},
		{
			listForList: []int{},
			expected:    0,
			testName:    "test-2",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			list := NewList()
			for _, v := range testCase.listForList {
				list.PushBack(v)
			}
			actual := list.Len()
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestList_Front(t *testing.T) {
	TestTable := []struct {
		listForList []int
		expected    *ListItem
		testName    string
	}{
		{
			listForList: []int{1, 2, 3},
			expected:    NewListItem(1),
			testName:    "test-1",
		},
		{
			listForList: []int{1, 2, 3, 4, 5},
			expected:    NewListItem(1),
			testName:    "test-2",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			list := NewList()
			for _, v := range testCase.listForList {
				list.PushBack(v)
			}
			actual := list.Front()
			require.Equal(t, testCase.expected.Value, actual.Value)
		})
	}
}

func TestList_Back(t *testing.T) {
	TestTable := []struct {
		listForList []int
		expected    *ListItem
		testName    string
	}{
		{
			listForList: []int{1, 2, 3},
			expected:    NewListItem(3),
			testName:    "test-1",
		},
		{
			listForList: []int{1, 2, 3, 4, 5},
			expected:    NewListItem(5),
			testName:    "test-2",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			list := NewList()
			for _, v := range testCase.listForList {
				list.PushBack(v)
			}
			actual := list.Back()
			require.Equal(t, testCase.expected.Value, actual.Value)
		})
	}
}

func TestList_PushFront(t *testing.T) {
	TestTable := []struct {
		listForList []int
		newValue    interface{}
		expected    *ListItem
		testName    string
	}{
		{
			listForList: []int{1, 2, 3},
			newValue:    11111,
			expected:    NewListItem(11111),
			testName:    "test-1",
		},
		{
			listForList: []int{1, 2, 3},
			newValue:    "aaaa",
			expected:    NewListItem("aaaa"),
			testName:    "test-2",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			list := NewList()
			for _, v := range testCase.listForList {
				list.PushBack(v)
			}
			actual := list.PushFront(testCase.newValue)
			require.Equal(t, testCase.expected.Value, actual.Value)
		})
	}
}

func TestList_PushBack(t *testing.T) {
	TestTable := []struct {
		listForList []int
		newValue    interface{}
		expected    *ListItem
		testName    string
	}{
		{
			listForList: []int{1, 2, 3},
			newValue:    11111,
			expected:    NewListItem(11111),
			testName:    "test-1",
		},
		{
			listForList: []int{1, 2, 3},
			newValue:    "aaaa",
			expected:    NewListItem("aaaa"),
			testName:    "test-2",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			list := NewList()
			for _, v := range testCase.listForList {
				list.PushBack(v)
			}
			actual := list.PushBack(testCase.newValue)
			require.Equal(t, testCase.expected.Value, actual.Value)
		})
	}
}

// вопрос ниже
//func TestList_Remove(t *testing.T) {
//	TestTable := []struct {
//		listForList []int
//		nodeToRemove *ListItem
//		expected    []int
//		testName    string
//	}{
//		{
//			listForList: []int{1, 2, 3, 4},
//			nodeToRemove:  NewListItem(1),
//			expected:    []int{2, 3, 4},
//			testName:    "test-1",
//		},
//	}
//	for _, testCase := range TestTable {
//		t.Run(testCase.testName, func(t *testing.T) {
//			list := NewList()
//			for _, v := range testCase.listForList {
//				list.PushBack(v)
//			}
//			list.Remove(testCase.nodeToRemove)
//			var resultList []interface{}
//			for i := 2; i < 5; i++ {
//				resultList = append(resultList, list.Front().Value)
//				nodeToDelete := NewListItem(i)
//				// при удалении последнего элемента попадаю на nil и паника
//				// не могу понять, как этого избежать, наверное сам Remove не корректно написан
//				list.Remove(nodeToDelete)
//			}
//			fmt.Println(resultList)
//
//			//require.Equal(t, testCase.expected.Value, actual.Value)
//		})
//	}
//}

func TestList_MoveToFront(t *testing.T) {
	TestTable := []struct {
		listForList []int
		newNode     *ListItem
		expected    int
		testName    string
	}{
		{
			listForList: []int{1, 2, 3},
			newNode:     NewListItem(3),
			expected:    3,
			testName:    "test-1",
		},
		{
			listForList: []int{1, 2, 3},
			newNode:     NewListItem(2),
			expected:    2,
			testName:    "test-2",
		},
	}
	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			list := NewList()
			for _, v := range testCase.listForList {
				list.PushBack(v)
			}
			list.MoveToFront(testCase.newNode)
			actual := list.Front()
			require.Equal(t, testCase.expected, actual.Value)
		})
	}
}
