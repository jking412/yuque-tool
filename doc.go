package main

import "time"

type Document struct {
	Abilities struct {
		Update  bool `json:"update"`
		Destroy bool `json:"destroy"`
	} `json:"abilities"`
	Data struct {
		ID     int    `json:"id"`
		Slug   string `json:"slug"`
		Title  string `json:"title"`
		BookID int    `json:"book_id"`
		Book   struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Slug             string      `json:"slug"`
			Name             string      `json:"name"`
			UserID           int         `json:"user_id"`
			Description      interface{} `json:"description"`
			CreatorID        int         `json:"creator_id"`
			Public           int         `json:"public"`
			ItemsCount       int         `json:"items_count"`
			LikesCount       int         `json:"likes_count"`
			WatchesCount     int         `json:"watches_count"`
			ContentUpdatedAt time.Time   `json:"content_updated_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			CreatedAt        time.Time   `json:"created_at"`
			Namespace        string      `json:"namespace"`
			User             struct {
				ID               int         `json:"id"`
				Type             string      `json:"type"`
				Login            string      `json:"login"`
				Name             string      `json:"name"`
				Description      interface{} `json:"description"`
				AvatarURL        string      `json:"avatar_url"`
				BooksCount       int         `json:"books_count"`
				PublicBooksCount int         `json:"public_books_count"`
				FollowersCount   int         `json:"followers_count"`
				FollowingCount   int         `json:"following_count"`
				CreatedAt        time.Time   `json:"created_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				Serializer       string      `json:"_serializer"`
			} `json:"user"`
			Serializer string `json:"_serializer"`
		} `json:"book"`
		UserID  int `json:"user_id"`
		Creator struct {
			ID               int         `json:"id"`
			Type             string      `json:"type"`
			Login            string      `json:"login"`
			Name             string      `json:"name"`
			Description      interface{} `json:"description"`
			AvatarURL        string      `json:"avatar_url"`
			BooksCount       int         `json:"books_count"`
			PublicBooksCount int         `json:"public_books_count"`
			FollowersCount   int         `json:"followers_count"`
			FollowingCount   int         `json:"following_count"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Serializer       string      `json:"_serializer"`
		} `json:"creator"`
		Format            string      `json:"format"`
		Body              string      `json:"body"`
		BodyDraft         string      `json:"body_draft"`
		BodyHTML          string      `json:"body_html"`
		BodyLake          string      `json:"body_lake"`
		BodyDraftLake     string      `json:"body_draft_lake"`
		Public            int         `json:"public"`
		Status            int         `json:"status"`
		ViewStatus        int         `json:"view_status"`
		ReadStatus        int         `json:"read_status"`
		LikesCount        int         `json:"likes_count"`
		CommentsCount     int         `json:"comments_count"`
		ContentUpdatedAt  time.Time   `json:"content_updated_at"`
		DeletedAt         interface{} `json:"deleted_at"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		PublishedAt       time.Time   `json:"published_at"`
		FirstPublishedAt  time.Time   `json:"first_published_at"`
		WordCount         int         `json:"word_count"`
		Cover             interface{} `json:"cover"`
		Description       string      `json:"description"`
		CustomDescription interface{} `json:"custom_description"`
		Hits              int         `json:"hits"`
		Serializer        string      `json:"_serializer"`
	} `json:"data"`
}
