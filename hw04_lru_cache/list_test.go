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

func checkPushFirst(t *testing.T, el *ListItem, l List) {
	require.Equal(t, el, l.Front())
	require.Equal(t, el, l.Back())
	require.Equal(t, true, el.Prev == nil)
	require.Equal(t, true, el.Next == nil)
}

func checkPushSecond(t *testing.T, front *ListItem, back *ListItem, l List) {
	require.Equal(t, front, l.Front())
	require.Equal(t, back, l.Back())
	require.Equal(t, true, front.Prev == nil)
	require.Equal(t, back, front.Next)
	require.Equal(t, front, back.Prev)
	require.Equal(t, true, back.Next == nil)
}

func TestPushFront(t *testing.T) {
	l := NewList()
	back := l.PushFront(10)
	checkPushFirst(t, back, l)
	front := l.PushFront(20)
	checkPushSecond(t, front, back, l)
}

func TestPushBack(t *testing.T) {
	l := NewList()
	front := l.PushBack(10)
	checkPushFirst(t, front, l)
	back := l.PushBack(20)
	checkPushSecond(t, front, back, l)
}

func TestRemove(t *testing.T) {
	t.Run("delete_front", func(t *testing.T) {
		l := NewList()
		front := l.PushFront(10) // [10]
		middle := l.PushBack(11) // [10, 11]
		l.PushBack(12)           // [10, 11, 12]
		l.Remove(front)          // [11, 12]
		require.Equal(t, middle, l.Front())
		require.Equal(t, true, middle.Prev == nil)
	})

	t.Run("delete_back", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)          // [10]
		middle := l.PushBack(11) // [10, 11]
		back := l.PushBack(12)   // [10, 11, 12]
		l.Remove(back)
		require.Equal(t, middle, l.Back())
		require.Equal(t, true, middle.Next == nil)
	})

	t.Run("delete_middle", func(t *testing.T) {
		l := NewList()
		front := l.PushFront(10) // [10]
		middle := l.PushBack(11) // [10, 11]
		back := l.PushBack(12)   // [10, 11, 12]
		l.Remove(middle)
		require.Equal(t, back, front.Next)
		require.Equal(t, front, back.Prev)
	})

	t.Run("delete_last", func(t *testing.T) {
		l := NewList()
		item := l.PushFront(10)
		l.Remove(item)
		require.Equal(t, true, l.Front() == nil)
		require.Equal(t, true, l.Front() == nil)
		require.Equal(t, 0, l.Len())
	})

}

func TestMoveToFront(t *testing.T) {
	t.Run("move_front", func(t *testing.T) {
		l := NewList()
		front := l.PushFront(1)
		middle := l.PushBack(2)
		l.PushBack(3)
		l.MoveToFront(front)
		require.Equal(t, 3, l.Len())
		require.Equal(t, front, l.Front())
		require.Equal(t, middle, front.Next)
		require.Equal(t, front, middle.Prev)
	})

	t.Run("move_middle", func(t *testing.T) {
		l := NewList()
		l.PushBack(1)
		middle := l.PushBack(2)
		l.PushBack(3)
		l.MoveToFront(middle)
		require.Equal(t, middle, l.Front())
	})
}
