package fahrradwerkstatt

import (
	"reflect"
	"testing"
)

func TestSimulation(t *testing.T) {
	type args struct {
		auftraege []*Auftrag
		nextOrder func(auftraege []*Auftrag) *Auftrag
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 int
	}{
		{
			name: "testing two auftraege with duration one day starting on two consecutive days",
			args: args{auftraege: []*Auftrag{
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
				{
					Eingangszeitpunkt:        9*60 + 24*60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
			},
				nextOrder: aeltesterAuftrag,
			},
			want:  8 * 60,
			want1: 8 * 60,
		},
		{
			name: "testing two auftraege with duration one day starting on the same day",
			args: args{auftraege: []*Auftrag{
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
			},
				nextOrder: aeltesterAuftrag,
			},
			want:  (8*60 + (15+17)*60) / 2,
			want1: (15 + 17) * 60,
		},
		{
			name: "testing two auftraege with duration one day starting on two consecutive days",
			args: args{auftraege: []*Auftrag{
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
				{
					Eingangszeitpunkt:        9*60 + 24*60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
			},
				nextOrder: shortestAuftrag,
			},
			want:  8 * 60,
			want1: 8 * 60,
		},
		{
			name: "testing two auftraege with duration one day starting on the same day",
			args: args{auftraege: []*Auftrag{
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        8 * 60,
					Restdauer:                8 * 60,
					Fertigstellungszeitpunkt: 0,
				},
			},
				nextOrder: shortestAuftrag},
			want:  (8*60 + (15+17)*60) / 2,
			want1: (15 + 17) * 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Simulation(tt.args.auftraege, tt.args.nextOrder)
			if got != tt.want {
				t.Errorf("Simulation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Simulation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_sortArrayByShortest(t *testing.T) {
	type args struct {
		array []*Auftrag
	}
	tests := []struct {
		name string
		args args
		want []*Auftrag
	}{
		{
			name: "testing array with two elements not correctly sorted",
			args: args{
				array: []*Auftrag{
					{
						Eingangszeitpunkt:        9 * 60,
						Bearbeitungsdauer:        200,
						Restdauer:                200,
						Fertigstellungszeitpunkt: 0,
					},
					{
						Eingangszeitpunkt:        9*60 + 24*60,
						Bearbeitungsdauer:        100,
						Restdauer:                100,
						Fertigstellungszeitpunkt: 0,
					},
				},
			},
			want: []*Auftrag{
				{
					Eingangszeitpunkt:        9*60 + 24*60,
					Bearbeitungsdauer:        100,
					Restdauer:                100,
					Fertigstellungszeitpunkt: 0,
				},
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        200,
					Restdauer:                200,
					Fertigstellungszeitpunkt: 0,
				},
			},
		},
		{
			name: "testing array with two elements already correctly sorted",
			args: args{
				array: []*Auftrag{
					{
						Eingangszeitpunkt:        9*60 + 24*60,
						Bearbeitungsdauer:        100,
						Restdauer:                100,
						Fertigstellungszeitpunkt: 0,
					},
					{
						Eingangszeitpunkt:        9 * 60,
						Bearbeitungsdauer:        200,
						Restdauer:                200,
						Fertigstellungszeitpunkt: 0,
					},
				},
			},
			want: []*Auftrag{
				{
					Eingangszeitpunkt:        9*60 + 24*60,
					Bearbeitungsdauer:        100,
					Restdauer:                100,
					Fertigstellungszeitpunkt: 0,
				},
				{
					Eingangszeitpunkt:        9 * 60,
					Bearbeitungsdauer:        200,
					Restdauer:                200,
					Fertigstellungszeitpunkt: 0,
				},
			},
		},
		{
			name: "testing empty array",
			args: args{
				array: []*Auftrag{},
			},
			want: []*Auftrag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByShortest(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByShortest() = %v, want %v", got, tt.want)
			}
		})
	}
}
