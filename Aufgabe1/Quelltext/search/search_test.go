package search

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestSearchSentencesWithSearchFile(t *testing.T) {
	txt, err := os.ReadFile("../Alice_im_Wunderland.txt")
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		txt      string
		filename string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test stoerung0.txt",
			args: args{txt: string(txt), filename: "../Beispielaufgaben/stoerung0.txt"},
			want: []string{
				"Das kommt mir gar nicht richtig vor",
			},
		},
		{
			name: "test stoerung1.txt",
			args: args{txt: string(txt), filename: "../Beispielaufgaben/stoerung1.txt"},
			want: []string{
				"Ich muß in Clara verwandelt",
				"Ich muß doch Clara sein",
			},
		},
		{
			name: "test stoerung2.txt",
			args: args{txt: string(txt), filename: "../Beispielaufgaben/stoerung2.txt"},
			want: []string{
				"Fressen Katzen gern Spatzen",
				"Fressen Katzen gern Spatzen",
				"Fressen Spatzen gern Katzen",
			},
		},
		{
			name: "test stoerung3.txt",
			args: args{txt: string(txt), filename: "../Beispielaufgaben/stoerung3.txt"},
			want: []string{
				"das Spiel fing an",
				"Das Publikum fing an",
			},
		},
		{
			name: "test stoerung4.txt",
			args: args{txt: string(txt), filename: "../Beispielaufgaben/stoerung4.txt"},
			want: []string{
				"ein sehr schöner Tag",
			},
		},
		{
			name: "test stoerung5.txt",
			args: args{txt: string(txt), filename: "../Beispielaufgaben/stoerung5.txt"},
			want: []string{
				"Wollen Sie so gut sein",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchSentencesWithSearchFile(tt.args.txt, tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchSentencesWithSearchFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
