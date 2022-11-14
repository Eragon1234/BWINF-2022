package fahrradwerkstatt

import "testing"

func TestAeltesterAuftragZuerst(t *testing.T) {
	type args struct {
		auftraege []*Auftrag
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
			}},
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
			}},
			want:  (8*60 + (15+17)*60) / 2,
			want1: (15 + 17) * 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AeltesterAuftragZuerst(tt.args.auftraege)
			if got != tt.want {
				t.Errorf("AeltesterAuftragZuerst() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AeltesterAuftragZuerst() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestKuerzesterAuftragZuerst(t *testing.T) {
	type args struct {
		auftraege []*Auftrag
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
			}},
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
			}},
			want:  (8*60 + (15+17)*60) / 2,
			want1: (15 + 17) * 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := KuerzesterAuftragZuerst(tt.args.auftraege)
			if got != tt.want {
				t.Errorf("KuerzesterAuftragZuerst() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("KuerzesterAuftragZuerst() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
