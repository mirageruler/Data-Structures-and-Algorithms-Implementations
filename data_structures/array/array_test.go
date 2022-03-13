package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tests := []struct {
		name       string
		manipulate func(a *Array) *Array
		assert     func(a *Array)
	}{
		{
			name: "case1: successfully get",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
					3: struct {
						Name       string
						Age        int
						IsHandsome bool
					}{
						Name:       "Khoi",
						Age:        24,
						IsHandsome: true,
					},
				}
				a.Length = 4

				return a
			},
			assert: func(a *Array) {
				// Getting
				data, err := a.Get(3)
				require.Nil(err)
				// Asserting
				assert.Equal(data, struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				})
			},
		},
		{
			name: "case2: fail to get due to invalid index",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
					3: struct {
						Name       string
						Age        int
						IsHandsome bool
					}{
						Name:       "Khoi",
						Age:        24,
						IsHandsome: true,
					},
				}
				a.Length = 4

				return a
			},
			assert: func(a *Array) {
				// Getting
				data, err := a.Get(100)
				require.NotNil(err)
				// Asserting
				assert.Equal(data, nil)
			},
		},
		{
			name: "case3: fail to get due to nil array",
			manipulate: func(a *Array) *Array {
				a.Push(2)
				a = nil

				return a
			},
			assert: func(a *Array) {
				// Get
				data, err := a.Get(2)
				require.NotNil(err)
				assert.Nil(data)
			},
		},
	}

	// Preparing data
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			array := &Array{}
			a := tt.manipulate(array)
			tt.assert(a)
		})
	}

}

func Test_Push(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tests := []struct {
		name       string
		manipulate func(a *Array) *Array
		assert     func(a *Array)
	}{
		{
			name: "case1: successfully push",
			manipulate: func(a *Array) *Array {
				require.Nil(a.Push(map[string]interface{}{
					"key1": 10,
				}))
				return a
			},
			assert: func(a *Array) {
				assert.Equal(a.Length, 1)
			},
		},
		{
			name: "case2: fail to push due to nil array",
			manipulate: func(a *Array) *Array {
				a.Push(2)
				a = nil
				return a
			},
			assert: func(a *Array) {
				require.NotNil(a.Push("10"))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			array := &Array{}
			a := tt.manipulate(array)
			tt.assert(a)
		})
	}
}

func Test_Pop(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tests := []struct {
		name       string
		manipulate func(a *Array) *Array
		assert     func(a *Array)
	}{
		{
			name: "case1: successfully pop",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
				}
				a.Length = 3

				require.Nil(a.Pop())

				return a
			},
			assert: func(a *Array) {
				assert.Equal(a.Data[1], 20)
			},
		},
		{
			name: "case2: fail to pop due to nil array",
			manipulate: func(a *Array) *Array {
				a.Push(2)
				a = nil
				return a
			},
			assert: func(a *Array) {
				assert.NotNil(a.Pop())
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			array := &Array{}
			a := tt.manipulate(array)
			tt.assert(a)
		})
	}

}

func Test_Delete(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tests := []struct {
		name       string
		manipulate func(a *Array) *Array
		assert     func(a *Array)
	}{
		{
			name: "case1: successfully delete an element within the array",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
					3: struct {
						Name       string
						Age        int
						IsHandsome bool
					}{
						Name:       "Khoi",
						Age:        24,
						IsHandsome: true,
					},
				}
				a.Length = 4

				// Deleting
				require.Nil(a.Delete(2))

				return a
			},
			assert: func(a *Array) {
				assert.Equal(a.Data[2], struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				})
			},
		},
		{
			name: "case2: fail to insert due to nil array",
			manipulate: func(a *Array) *Array {
				a.Push(2)
				a = nil
				return a
			},
			assert: func(a *Array) {
				assert.NotNil(a.Delete(2))
			},
		},
		{
			name: "case3: fail to delete due to invalid index",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
					3: struct {
						Name       string
						Age        int
						IsHandsome bool
					}{
						Name:       "Khoi",
						Age:        24,
						IsHandsome: true,
					},
				}
				a.Length = 4

				return a
			},
			assert: func(a *Array) {
				assert.NotNil(a.Delete(100))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			array := &Array{}
			a := tt.manipulate(array)
			tt.assert(a)
		})
	}
}

func Test_Insert(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tests := []struct {
		name       string
		manipulate func(a *Array) *Array
		assert     func(a *Array)
	}{
		{
			name: "case1: normally success",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
				}
				a.Length = 3

				// Inserting
				require.Nil(a.Insert(struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				}, 2))
				return a
			},
			assert: func(a *Array) {
				assert.Equal(a.Data[2], struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				})
			},
		},
		{
			name: "case2: successfully insert at first index",
			manipulate: func(a *Array) *Array {
				// Inserting
				require.Nil(a.Insert(struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				}, 0))
				return a
			},
			assert: func(a *Array) {
				assert.Equal(a.Data[0], struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				})
			},
		},
		{
			name: "case3: fail to insert due to nil array",
			manipulate: func(a *Array) *Array {
				a.Push(2)
				a = nil
				return a
			},
			assert: func(a *Array) {
				assert.NotNil(a.Insert("10", 2))
			},
		},
		{
			name: "case4: fail to insert due to invalid index",
			manipulate: func(a *Array) *Array {
				a.Data = map[int]interface{}{
					0: "10",
					1: 20,
					2: []int{1, 2, 3, 4},
				}
				a.Length = 3

				return a
			},
			assert: func(a *Array) {
				assert.NotNil(a.Insert(struct {
					Name       string
					Age        int
					IsHandsome bool
				}{
					Name:       "Khoi",
					Age:        24,
					IsHandsome: true,
				}, 100))
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			array := &Array{}
			a := tt.manipulate(array)
			tt.assert(a)
		})
	}
}
