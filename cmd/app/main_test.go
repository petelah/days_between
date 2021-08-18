package main

import (
	"reflect"
	"testing"
)

func Test_initDate(t *testing.T) {
	type args struct {
		inputDate  string
	}
	tests := []struct {
		name    string
		args    args
		want    *date
		wantErr bool
	}{
		{name: "Pass with valid date", args: args{inputDate: "2/6/2022"}, want: &date{day: 2, month: 6, year: 2022}, wantErr: false},
		{name: "Fail on invalid input", args: args{inputDate: "2/6/20ff"}, want: nil, wantErr: true},
		{name: "Fail on date out of range: year too high", args: args{inputDate: "2/6/3000"}, want: nil, wantErr: true},
		{name: "Fail on date out of range: year too low", args: args{inputDate: "2/6/1899"}, want: nil, wantErr: true},
		{name: "Fail on date out of range: month too high", args: args{inputDate: "2/16/2020"}, want: nil, wantErr: true},
		{name: "Fail on date out of range: day too high", args: args{inputDate: "32/12/2020"}, want: nil, wantErr: true},
		{name: "Fail invalid date formatting '-' used instead of '/'", args: args{inputDate: "32-12-2020"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initDate(tt.args.inputDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("initDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Is leap year", args: args{year: 2000}, want: true},
		{name: "Is not leap year", args: args{year: 2002}, want: false},
		{name: "Is not leap year", args: args{year: 1700}, want: false},
		{name: "Is not leap year", args: args{year: 1900}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_daysBetween(t *testing.T) {
	type args struct {
		firstDate  date
		secondDate date
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "Pass case a 19 days between", args: args{firstDate: date{day: 2, month: 6, year: 1983}, secondDate: date{day: 22, month: 6, year: 1983}}, want: 19, wantErr: false},
		{name: "Pass case b 173 days", args: args{firstDate: date{day: 4, month: 7, year: 1984}, secondDate: date{day: 25, month: 12, year: 1984}}, want: 173, wantErr: false},
		{name: "Pass case c 2036 days", args: args{firstDate: date{day: 1, month: 3, year: 1989}, secondDate: date{day: 3, month: 8, year: 1983}}, want: 2036, wantErr: false},
		{name: "Pass include leap year", args: args{firstDate: date{day: 2, month: 6, year: 1999}, secondDate: date{day: 2, month: 6, year: 2002}}, want: 1095, wantErr: false},
		{name: "Pass case 0 days between", args: args{firstDate: date{day: 1, month: 1, year: 1999}, secondDate: date{day: 1, month: 1, year: 1999}}, want: 1095, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := daysBetween(tt.args.firstDate, tt.args.secondDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("daysBetween() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("daysBetween() got = %v, want %v", got, tt.want)
			}
		})
	}
}
