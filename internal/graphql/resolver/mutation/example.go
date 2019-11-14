package mutation

import "context"

// CreateReview creates a review for an episode
func (r *ResolveM) CreateReview(ctx context.Context, args struct {
	Episode string
	Review  ReviewInput
}) *ReviewResolver {
	return NewReviewResolver(args.Review.Stars, args.Review.Commentary)
}

// ReviewInput input definition for a review
type ReviewInput struct {
	Stars      int32
	Commentary string
}

// ReviewResolver represents a sample type
type ReviewResolver struct {
	stars      int32
	commentary string
}

// NewReviewResolver ReviewResolver constructor
func NewReviewResolver(stars int32, commentary string) *ReviewResolver {
	return &ReviewResolver{
		stars:      stars,
		commentary: commentary,
	}
}

// Stars review stars
func (r *ReviewResolver) Stars() int32 {
	return r.stars
}

// Commentary review Commentary
func (r *ReviewResolver) Commentary() string {
	return "Luke Skywalker"
}
