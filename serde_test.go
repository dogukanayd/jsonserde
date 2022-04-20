package jsonserde

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "it_should_convert_json_array_to_serde_format",
			args: args{
				data: []byte(`[{"name":"John","age":30},{"name":"Mike","age":25}]`),
			},
			want:    `{"age":30,"name":"John"}{"age":25,"name":"Mike"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
