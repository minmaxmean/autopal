package life

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EmptyGame(t *testing.T) {
	require := require.New(t)
	rows := 3
	cols := 4
	gotState := EmptyGame(cols, rows)
	wantState := &GameState{
		cols: cols, rows: rows,
		cells: []bool{false, false, false, false, false, false, false, false, false, false, false, false},
	}
	require.Equal(wantState, gotState)

	wantStr := "....\n" + "....\n" + "...."
	gotStr := gotState.String()
	require.Equal(wantStr, gotStr)
}

func Test_FromString(t *testing.T) {
	testCases := []struct {
		name      string
		rows      int
		cols      int
		str       string
		wantErr   bool
		wantCells []bool
	}{
		{
			name: "empty board with dots",
			rows: 3,
			cols: 4,
			str: `
			....
			....
			....
`,
			wantCells: []bool{false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			name: "empty board with dots",
			rows: 3,
			cols: 4,
			str: `
			#...
			.#..
			..#.
`,
			wantCells: []bool{true, false, false, false, false, true, false, false, false, false, true, false},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			gotState, err := FromString(test.cols, test.rows, test.str)
			if test.wantErr {
				require.Error(err)
				return
			}
			require.NoError(err)
			wantState := &GameState{cols: test.cols, rows: test.rows, cells: test.wantCells}
			require.Equal(wantState, gotState)
		})
	}
}

func Test_Next_4Rules(t *testing.T) {
	oscilators := []struct {
		name       string
		rows       int
		cols       int
		init_state string
		want_state string
	}{
		{
			name: "n <= 1",
			rows: 5,
			cols: 5,
			init_state: `
				.....
				.OO..
				.....
				.....
				.....
			`,
			want_state: `
				.....
				.....
				.....
				.....
				.....
			`,
		},
		{
			name: "2 <= n <= 3",
			rows: 5,
			cols: 5,
			init_state: `
				.....
				.OO..
				.O...
				.....
				.....
			`,
			want_state: `
				.....
				.OO..
				.OO..
				.....
				.....
			`,
		},
		{
			name: "4 <= n",
			rows: 9,
			cols: 9,
			init_state: `
				.........
				.........
				.........
				...OOO...
				...OOO...
				...OOO...
				.........
				.........
				.........
			`,
			want_state: `
				.........
				.........
				....O....
				...O.O...
				..O...O..
				...O.O...
				....O....
				.........
				.........
			`,
		},
	}
	for _, test := range oscilators {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			initState, err := FromString(test.cols, test.rows, test.init_state)
			require.NoError(err)
			requireState(t, test.init_state, initState)

			gotState := initState.Next()
			requireState(t, test.want_state, gotState)
		})
	}
}

func Test_Next_OnOsclillators(t *testing.T) {
	oscilators := []struct {
		name        string
		rows        int
		cols        int
		init_state  string
		want_boards []string
	}{
		{
			name: "Blinker",
			rows: 5,
			cols: 5,
			init_state: `
				.....
				..O..
				..O..
				..O..
				.....
			`,
			want_boards: []string{`
				.....
				.....
				.OOO.
				.....
				.....
			`, `
				.....
				..O..
				..O..
				..O..
				.....
			`},
		},
		{
			name: "4 <= n",
			rows: 9,
			cols: 9,
			init_state: `
				.........
				.........
				.........
				...OOO...
				...OOO...
				...OOO...
				.........
				.........
				.........
			`,
			want_boards: []string{`
				.........
				.........
				....O....
				...O.O...
				..O...O..
				...O.O...
				....O....
				.........
				.........
			`,
				`
				.........
				.........
				....O....
				...OOO...
				..OO.OO..
				...OOO...
				....O....
				.........
				.........
			`,
				`
				.........
				.........
				...OOO...
				..O...O..
				..O...O..
				..O...O..
				...OOO...
				.........
				.........
			`,
			},
		},
		{
			name: "Toad",
			rows: 6,
			cols: 6,
			init_state: `
				......
				...O..
				.O..O.
				.O..O.
				..O...
				......
			`,
			want_boards: []string{`
				......
				......
				..OOO.
				.OOO..
				......
				......
			`, `
				......
				...O..
				.O..O.
				.O..O.
				..O...
				......
			`},
		},
		{
			name: "Pulsar",
			rows: 17,
			cols: 17,
			init_state: `
				.................
				.................
				....OOO...OOO....
				.................
				..O....O.O....O..
				..O....O.O....O..
				..O....O.O....O..
				....OOO...OOO....
				.................
				....OOO...OOO....
				..O....O.O....O..
				..O....O.O....O..
				..O....O.O....O..
				.................
				....OOO...OOO....
				.................
				.................
			`,
			want_boards: []string{`
				.................
				.....O.....O.....
				.....O.....O.....
				.....OO...OO.....
				.................
				.OOO..OO.OO..OOO.
				...O.O.O.O.O.O...
				.....OO...OO.....
				.................
				.....OO...OO.....
				...O.O.O.O.O.O...
				.OOO..OO.OO..OOO.
				.................
				.....OO...OO.....
				.....O.....O.....
				.....O.....O.....
				.................
				`, `
				.................
				.................
				....OO.....OO....
				.....OO...OO.....
				..O..O.O.O.O..O..
				..OOO.OO.OO.OOO..
				...O.O.O.O.O.O...
				....OOO...OOO....
				.................
				....OOO...OOO....
				...O.O.O.O.O.O...
				..OOO.OO.OO.OOO..
				..O..O.O.O.O..O..
				.....OO...OO.....
				....OO.....OO....
				.................
				.................
			`, `
				.................
				.................
				....OOO...OOO....
				.................
				..O....O.O....O..
				..O....O.O....O..
				..O....O.O....O..
				....OOO...OOO....
				.................
				....OOO...OOO....
				..O....O.O....O..
				..O....O.O....O..
				..O....O.O....O..
				.................
				....OOO...OOO....
				.................
				.................
			`},
		},
	}
	for _, test := range oscilators {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			gotState, err := FromString(test.cols, test.rows, test.init_state)
			require.NoError(err)
			for i, want_cell := range test.want_boards {
				gotState = gotState.Next()
				requireState(t, want_cell, gotState, "iteration #%d does not match", i+1)
			}
		})
	}
}

func requireState(t testing.TB, wantBoard string, gotState *GameState, msgAndArgs ...any) {
	wantState, err := FromString(gotState.cols, gotState.rows, wantBoard)
	require.NoError(t, err)
	if !assert.Equal(t, wantState.cells, gotState.cells) {
		fmt.Printf("got:\n%s", gotState.String())
	}
	require.Equal(t, wantState.cells, gotState.cells, msgAndArgs...)
}
