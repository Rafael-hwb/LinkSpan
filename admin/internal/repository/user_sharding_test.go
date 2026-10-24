package repository

import "testing"

func TestUserShardingCompatibility(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		wantHash  int32
		wantTable string
	}{
		{
			name:      "已有测试用户",
			username:  "test",
			wantHash:  3556498,
			wantTable: "t_user_2",
		},
		{
			name:      "带数字用户名",
			username:  "test2",
			wantHash:  110251488,
			wantTable: "t_user_0",
		},
		{
			name:      "alice",
			username:  "alice",
			wantHash:  92903040,
			wantTable: "t_user_0",
		},
		{
			name:      "bob",
			username:  "bob",
			wantHash:  97717,
			wantTable: "t_user_5",
		},
		{
			name:      "中文算法兼容",
			username:  "用户甲",
			wantHash:  29631747,
			wantTable: "t_user_3",
		},
		{
			name:      "UTF16代理对",
			username:  "😀",
			wantHash:  1772899,
			wantTable: "t_user_3",
		},
		{
			name:      "最小32位整数哈希",
			username:  "polygenelubricants",
			wantHash:  -2147483648,
			wantTable: "t_user_0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := javaStringHash(tt.username)
			if hash != tt.wantHash {
				t.Errorf(
					"hash = %d，期望 %d",
					hash,
					tt.wantHash,
				)
			}

			table := UserTableName(tt.username)
			if table != tt.wantTable {
				t.Errorf(
					"table = %s，期望 %s",
					table,
					tt.wantTable,
				)
			}
		})
	}
}