package hw03frequencyanalysis

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestDash(t *testing.T) {
	input := `a --- b - j -`
	expected := []string{`---`, `a`, `b`, `j`}
	result := Top10(input)
	require.Equal(t, expected, result)
}

func TestFrequencyRange(t *testing.T) {
	input := map[string]int{
		`a`: 4,
		`b`: 3,
		`c`: 4,
	}
	excepted := map[int][]string{
		4: {`a`, `c`},
		3: {`b`},
	}
	result := frequencyRange(input)
	for key, value := range result {
		equil := reflect.DeepEqual(value, excepted[key])
		require.True(t, equil)
	}
}

func TestTop(t *testing.T) {
	cases := []struct {
		name     string
		input    map[int][]string
		excepted []string
	}{
		{
			name: `full`,
			input: map[int][]string{
				3: {`a`, `b`},
				1: {`c`},
				2: {`e`, `d`},
				7: {`f`, `g`},
				5: {`h`, `k`, `j`},
				4: {`n`, `m`, `l`},
			},
			excepted: []string{
				`f`, `g`, `h`, `j`, `k`, `l`, `m`, `n`, `a`, `b`,
			},
		},
		{
			name: `short_block`,
			input: map[int][]string{
				2: {`a`, `b`},
			},
			excepted: []string{
				`a`, `b`,
			},
		},
		{
			name: `overflow`,
			input: map[int][]string{
				5: {`a`, `b`, `c`, `d`, `e`},
				4: {`f`, `g`},
				3: {`h`, `i`, `j`, `k`, `l`},
			},
			excepted: []string{
				`a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`, `i`, `j`,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := top(tc.input)
			require.Equal(t, tc.excepted, result)
		})
	}
}
