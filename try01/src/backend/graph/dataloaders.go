package graph

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/sky0621/study-graphql/try01/src/backend/graph/model"
	"github.com/sky0621/study-graphql/try01/src/backend/sqlboiler"
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
				fetch: func(ids []int64) ([]*model.User, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					log.Printf("userIDs: %#+v\n", ids)
					users, err := sqlboiler.Users(sqlboiler.UserWhere.ID.IN(ids)).All(r.Context(), db)
					if err != nil {
						log.Print(err)
						return nil, []error{err}
					}

					// ids の中の id 毎にデータをマッピングする必要がある。
					userById := map[int64]*model.User{}
					for _, user := range users {
						userById[user.ID] = &model.User{
							ID:   user.ID,
							Name: user.Name,
						}
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
