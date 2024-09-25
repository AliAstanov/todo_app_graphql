package postgres

import (
	"context"
	"example/graph/model"
	"reflect"
	"testing"
)

var (
	Id_1 = "f6b409cc-8520-4532-4d80-a6c7a6092a8c"
	Id_2 = "5cd22b4d-29fd-4a78-4cc6-cd3af539c363"
	Id_3 = "25bbea4f-3c82-4a51-89c2-031f414e5b72"
)

func TestUser_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   *model.User
		want    *model.User
		wantErr bool
	}{
		{
			name: "soccessfull",
			input: &model.User{
				UserID:   Id_1,
				Username: "name_1",
				Password: "Password_1",
			},
			want: &model.User{
				UserID:   Id_1,
				Username: "name_1",
				Password: "Password_1",
			},
			wantErr: false,
		},
		{
			name: "soccessfull_2",
			input: &model.User{
				UserID:   Id_2,
				Username: "name_2",
				Password: "password_2",
			},
			want: &model.User{
				UserID:   Id_2,
				Username: "name_2",
				Password: "password_2",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.CreateUsers(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}

}

func TestUser_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   *model.User
		want    *model.User
		wantErr bool
	}{
		{
			name: "soccessfull_get",
			input: &model.User{
				UserID:   Id_3,
				Username: "Akmal",
				Password: "Password-2",
			},
			want: &model.User{
				UserID:   Id_3,
				Username: "Akmal",
				Password: "password-2",
			},
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.GetUserById(context.Background(), tc.input.UserID)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestUser_List(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			limit, page string
		}
		want    []*model.User
		wantErr bool
	}{
		{
			name: "soccesfull_list",
			input: struct {
				limit string
				page  string
			}{page: "1", limit: "5"},
			want: []*model.User{
				{
					UserID:   Id_3,
					Username: "Akmal",
					Password: "password-2",
				},
				{
					UserID:   Id_1,
					Username: "name_1",
					Password: "Password_1",
				},
				{
					UserID:   Id_2,
					Username: "name_2",
					Password: "password_2",
				},
			},
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			got, err := pgRepo.GetUsers(context.Background(), tc.input.limit, tc.input.page)
			if err != nil {
				t.Fatalf("%s: expected: %v, got %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s expected: %v, got %v", tc.name, tc.wantErr, got)
			}
		})
	}
}

func TestUser_Updating(t *testing.T) {
	tests := []struct {
		name    string
		input   *model.User
		want    *model.User
		wantErr bool
	}{
		{
			name: "soccessfully_update",
			input: &model.User{
				UserID:   Id_3,
				Username: "name_3",
				Password: "password_3",
			},
			want: &model.User{
				UserID:   Id_3,
				Username: "name_3",
				Password: "password_3",
			},
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.UpdateUser(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   *model.User
		want    *model.User
		wantErr bool
	}{
		{
			name: "soccsessfully_delete",
			input: &model.User{
				UserID: "25bbea4f-3c82-4a51-89c2-031f414e5b89",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := pgRepo.DeleteUserById(context.Background(), tc.input.UserID)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			got, err := pgRepo.GetUserById(context.Background(), tc.input.UserID)
			if err == nil {
				t.Fatalf("%s: expected user to be deleted but found: %v", tc.name, got)
			}
		})
	}
}

// func clearUsers(t *testing.T) {
// 	db, err := CallDb()
// 	defer db.Close()

// 	// Foydalanuvchilar jadvalini tozalang
// 	_, err = db.Exec("DELETE FROM users") // 'users' - sizning jadval nomingiz
// 	if err != nil {
// 		t.Fatalf("Failed to clear users: %v", err)
// 	}

// 	// Foydalanuvchini qo'shing
// 	_, err = db.Exec(`INSERT INTO users (user_id, username, password) VALUES
// 		($1, $2, $3)`,
// 		"25bbea4f-3c82-4a51-89c2-031f414e5b89", "name_to_delete", "password_to_delete",
// 	)
// 	if err != nil {
// 		t.Fatalf("Failed to insert user: %v", err)
// 	}
// }
