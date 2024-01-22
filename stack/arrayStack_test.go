package stack

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_ArrayStackPop(t *testing.T) {
	stack := &ArrayStack{}

	id := uuid.New()
	id_2 := uuid.New()
	id_3 := uuid.New()
	stack.Push(id)
	stack.Push(id_2)
	stack.Push(id_3)

	require.Equal(t, 3, stack.Size())
	require.Equal(t, id_3, stack.Pop())

	stack2 := &ArrayStack{}
	require.Equal(t, uuid.Nil, stack2.Pop())
}

func Test_ArrayStackPush(t *testing.T) {
	stack := &ArrayStack{}

	id := uuid.New()
	id_2 := uuid.New()
	id_3 := uuid.New()
	stack.Push(id)
	stack.Push(id_2)
	stack.Push(id_3)

	require.Equal(t, 3, stack.Size())
	require.Equal(t, id_3, stack.Peek())
}

func Test_ArrayStackEmptyandClear(t *testing.T) {
	stack := &ArrayStack{}

	id := uuid.New()
	id_2 := uuid.New()
	id_3 := uuid.New()
	stack.Push(id)
	stack.Push(id_2)
	stack.Push(id_3)

	require.Equal(t, 3, stack.Size())
	require.Equal(t, false, stack.IsEmpty())

	stack.Clear()
	require.Equal(t, 0, stack.Size())
	require.Equal(t, uuid.Nil, stack.Peek())
	require.Equal(t, true, stack.IsEmpty())
}

func Test_ArrayStackContains(t *testing.T) {
	stack := &ArrayStack{}

	id := uuid.New()
	id_2 := uuid.New()
	id_3 := uuid.New()
	stack.Push(id)
	stack.Push(id_2)
	stack.Push(id_3)

	require.Equal(t, 3, stack.Size())
	require.Equal(t, true, stack.Contains(id_2))
	require.Equal(t, false, stack.Contains(uuid.New()))
}

func Test_ArrayStackContainsConcurrent(t *testing.T) {
	stack := &ArrayStack{}

	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			stack.Push(uuid.New())
			wg.Done()
		}()
	}

	wg.Wait()
	require.Equal(t, 100, stack.Size())
}


func Test_ArrayStackCopy(t *testing.T) {
	stack := &ArrayStack{}

	id := uuid.New()
	id_2 := uuid.New()
	id_3 := uuid.New()
	stack.Push(id)
	stack.Push(id_2)
	stack.Push(id_3)

	stack_copy := stack.Copy()
	require.Equal(t, 3, stack.Size())
	require.Equal(t, []uuid.UUID{id, id_2, id_3}, stack_copy.Data)
}

func Test_ArrayStackReverse(t *testing.T) {
	stack := &ArrayStack{}

	id := uuid.New()
	id_2 := uuid.New()
	id_3 := uuid.New()
	stack.Push(id)
	stack.Push(id_2)
	stack.Push(id_3)

	stack.Reverse()
	require.Equal(t, 3, stack.Size())
	require.Equal(t, []uuid.UUID{id_3, id_2, id}, stack.Data)
}