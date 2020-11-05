package graph

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const loadersKey = "dataLoaders"

type Loaders struct {
	UsersByIDs UserLoader
}

func Middleware(db boil.ContextExecutor, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UsersByIDs: UserLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				// 最大１ミリ秒待機した結果、ないし、最大 100 個のGraphQLクエリが溜まった分の id のスライスが ids という名前で渡ってくる。
				fetch: func(ids []int64) ([]*model.User, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					sql := "SELECT * FROM user WHERE id IN (" + toPKs(ids) + ")"
					log.Print(sql)

					rows, err := db.QueryContext(r.Context(), sql)
					if err != nil {
						log.Print(err)
						return nil, []error{err}
					}
					var users []*model.User
					for rows.Next() {
						var user model.User
						if err := rows.Scan(&user); err != nil {
							log.Print(err)
							return nil, []error{err}
						}
						users = append(users, &user)
					}

					// ids の中の id 毎にデータをマッピングする必要がある。
					userById := map[int64]*model.User{}
					for _, user := range users {
						userById[user.ID] = user
					}
					results := make([]*model.User, len(ids))
					for i, id := range ids {
						results[i] = userById[id]
					}

					return results, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func toPKs(ids []int64) string {
	// ids をSQL文の IN 句に指定できる形に変換
	var pks []string
	for _, id := range ids {
		pks = append(pks, strconv.FormatInt(id, 10))
	}
	return strings.Join(pks, ",")
}
