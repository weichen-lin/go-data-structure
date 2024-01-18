package linkedList

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_Append(t *testing.T) {
	list := SingleLinkedList{}
	list.Append(uuid.New())

	require.Equal(t, 1, list.Length)
}

func Test_Delete_Singel(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()

	list.Append(id)

	err := list.Delete(id)
	require.NoError(t, err)
	require.Equal(t, list.Length, 0)
}

func Test_Delete_Empty(t *testing.T) {
	list := SingleLinkedList{}
	id := uuid.New()
	err := list.Delete(id)
	require.Error(t, err)
	require.Equal(t, list.Length, 0)
}

func Test_Prepend(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()

	list.Append(id)
	require.Equal(t, list.Head.Value, id)

	id_2 := uuid.New()
	list.Prepend(id_2)
	require.Equal(t, list.Head.Value, id_2)
	require.Equal(t, list.Head.Next.Value, id)
	require.Equal(t, list.Length, 2)
}

func Test_Prepend_Empty(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Prepend(id)

	require.Equal(t, list.Head.Value, id)
	require.Equal(t, list.Length, 1)
}

func Test_Search(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	index, err := list.Search(id_2)
	require.NoError(t, err)
	require.Equal(t, index, 1)
}

func Test_Search_Empty(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	_, err := list.Search(id)
	require.Error(t, err)
}

func Test_Search_Not_Exist(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	id_4 := uuid.New()
	_, err := list.Search(id_4)
	require.Error(t, err)
}

func Test_ValueOf(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	value, err := list.ValueOf(1)
	require.NoError(t, err)
	require.Equal(t, value, id_2)

	_, err = list.ValueOf(3)
	require.Error(t, err)

	_, err = list.ValueOf(-1)
	require.Error(t, err)
}

func Test_ValueOf_Empty(t *testing.T) {
	list := SingleLinkedList{}

	_, err := list.ValueOf(0)
	require.Error(t, err)
}

func Test_Concurrency(t *testing.T) {
	list := SingleLinkedList{}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			list.Append(uuid.New())
		}()
	}

	wg.Wait()

	require.Equal(t, 1000, list.Length)
}

func Test_Insert(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	id_4 := uuid.New()
	err := list.InsertBehindWithIndex(0, id_4)
	require.NoError(t, err)

	require.Equal(t, list.Length, 4)
	require.Equal(t, list.Head.Next.Value, id_4)

	sec_list := SingleLinkedList{}

	id_5 := uuid.New()
	err = sec_list.InsertBehindWithIndex(0, id_5)
	require.Error(t, err)

	id_6 := uuid.New()
	err = sec_list.InsertBehindWithIndex(1, id_6)
	require.Error(t, err)
}

func Test_Insert_OutofRange(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	id_4 := uuid.New()
	err := list.InsertBehindWithIndex(3, id_4)
	require.Error(t, err)
}

func Test_Insert_Append(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	id_4 := uuid.New()
	err := list.InsertBehindWithIndex(list.Length-1, id_4)
	require.NoError(t, err)

	id_append, err := list.ValueOf(list.Length - 1)
	require.NoError(t, err)

	require.Equal(t, list.Length, 4)
	require.Equal(t, id_append, id_4)
}

func Test_Insert_Middel(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	id_4 := uuid.New()
	err := list.InsertBehindWithIndex(1, id_4)
	require.NoError(t, err)

	id_append, err := list.ValueOf(2)
	require.NoError(t, err)

	require.Equal(t, list.Length, 4)
	require.Equal(t, id_append, id_4)
}

func Test_Delete(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	err := list.Delete(id_2)
	require.NoError(t, err)

	require.Equal(t, list.Length, 2)
	require.Equal(t, list.Head.Next.Value, id_3)
}

func Test_Delete_Not_EXIST(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	id_4 := uuid.New()
	err := list.Delete(id_4)
	require.Error(t, err)
}

func Test_Reverse(t *testing.T) {
	list := SingleLinkedList{}

	id := uuid.New()
	list.Append(id)

	id_2 := uuid.New()
	list.Append(id_2)

	id_3 := uuid.New()
	list.Append(id_3)

	list.Reverse()

	require.Equal(t, list.Length, 3)
	require.Equal(t, list.Head.Value, id_3)
	require.Equal(t, list.Head.Next.Value, id_2)
	require.Equal(t, list.Head.Next.Next.Value, id)
}
